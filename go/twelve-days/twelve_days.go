package twelve

import "fmt"

var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

const baseString = "On the %s day of Christmas my true love gave to me: %s."

// Verse returns the lyrics for a single verse (day) of the song The Twelve Days of Christmas
func Verse(verse int) string {
	daysGifts := ""
	for i := verse - 1; i >= 0; i-- {
		if verse > 0 && daysGifts != "" {
			daysGifts += ", "
			if i == 0 {
				daysGifts += "and "
			}
		}
		daysGifts += gifts[i]
	}
	return fmt.Sprintf(baseString, days[verse-1], daysGifts)
}

// Song returns the lyrics for the entire song The Twelve Days of Christmas
func Song() string {
	allVerses := ""
	for i := 1; i < 13; i++ {
		allVerses += Verse(i) + "\n"
	}
	return allVerses
}
