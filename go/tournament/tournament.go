package tournament

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
)

type teamStats struct {
	name   string
	wins   int
	draws  int
	losses int
	points int
}

// matchesPlayed tallies up the number of matches played by the team
func (team *teamStats) matchesPlayed() int {
	return team.wins + team.losses + team.draws
}

type tournament map[string]*teamStats

// getTeamStats will the stats for the if it exists or it will create a new one
func (t tournament) getTeamStats(teamName string) *teamStats {
	if team, exists := t[teamName]; exists {
		return team
	}
	newTeam := teamStats{name: teamName}
	t[teamName] = &newTeam
	return &newTeam
}

// Tally generates a tally from a given file describing tournament match results
// and outputs it the other given file
func Tally(input io.Reader, output io.Writer) error {
	r := csv.NewReader(input)
	r.Comma = ';'
	r.Comment = '#'

	tourn := tournament{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if len(record) != 3 {
			return errors.New("incorrect count of parameters")
		}

		team1 := tourn.getTeamStats(record[0])
		team2 := tourn.getTeamStats(record[1])
		switch record[2] {
		case "win":
			team1.wins++
			team1.points += 3
			team2.losses++
		case "draw":
			team1.draws++
			team1.points++
			team2.draws++
			team2.points++
		case "loss":
			team1.losses++
			team2.wins++
			team2.points += 3
		default:
			return errors.New("Match outcome not supported")
		}
	}

	teamResults := []*teamStats{}
	for _, v := range tourn {
		teamResults = append(teamResults, v)
	}
	sort.Slice(teamResults, func(i, j int) bool {
		team1 := teamResults[i]
		team2 := teamResults[j]
		if team1.points == team2.points {
			return team1.name < team2.name
		}
		return team1.points > team2.points
	})

	fmt.Fprintf(output, "Team                           | MP |  W |  D |  L |  P\n")
	for _, team := range teamResults {
		if _, err := fmt.Fprintf(output, "%-31v|%3d |%3d |%3d |%3d |%3d\n", team.name, team.matchesPlayed(), team.wins, team.draws, team.losses, team.points); err != nil {
			return err
		}
	}

	return nil
}
