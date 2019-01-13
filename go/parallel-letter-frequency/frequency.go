package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(sList []string) FreqMap {
	c := make(chan FreqMap)

	for _, s := range sList {
		go func(s string) {
			c <- Frequency(s)
		}(s)
	}

	var ret FreqMap
	for i := 0; i < len(sList); i++ {
		result := <-c
		if ret == nil {
			ret = result
		} else {
			for k, v := range result {
				ret[k] += v
			}
		}
	}

	return ret
}
