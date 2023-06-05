package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	fr := FreqMap{}
	ch := make(chan FreqMap)
	for _, v := range texts {
		go ChFreq(ch, v)
	}

	for i := 0; i < len(texts); i++ {
		select {
			case res := <- ch:
				for idx, val := range res {
					fr[idx] += val
				}
		}
	}

	return fr
}

func ChFreq(ch chan FreqMap, text string) {
	fm := Frequency(text)
	ch <- fm
}
