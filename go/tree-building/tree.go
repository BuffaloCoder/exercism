package tree

import (
	"errors"
	"sort"
)

// Record represents a single record containg an ID and it's Parent record
type Record struct {
	ID     int
	Parent int
}

// Node represents the a list of records in tree form, with each ID containing
// children records that refer to it as their parent
type Node struct {
	ID       int
	Children []*Node
}

// Build takes a slice of Records and returns a Node representation of them,
// or an error if a tree cannot be generated
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].Parent < records[j].Parent || records[i].ID < records[j].ID
	})
	if (records[0] != Record{}) {
		return nil, errors.New("No root")
	}

	maxID := len(records)
	var result = Node{}
	nodeCache := map[int]*Node{
		0: &result,
	}
	for _, record := range records[1:] {
		if _, exists := nodeCache[record.ID]; exists {
			return nil, errors.New("Node already exists")
		}
		if record.ID <= record.Parent {
			return nil, errors.New("ID less than or equal to Parent")
		}
		if record.ID >= maxID {
			return nil, errors.New("Non-continuous ID")
		}

		if parent, exists := nodeCache[record.Parent]; exists {
			newNode := Node{ID: record.ID}
			parent.Children = append(parent.Children, &newNode)
			nodeCache[record.ID] = &newNode
		} else {
			return nil, errors.New("Parent does not exist")
		}
	}

	return &result, nil
}
