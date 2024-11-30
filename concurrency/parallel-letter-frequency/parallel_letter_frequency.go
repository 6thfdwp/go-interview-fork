package letterfrq

import (
	"fmt"
	"sync"
)

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

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.

func ConcurrentFrequency(l []string) FreqMap {
	// m := FreqMap{}
	m := make(FreqMap)
	var wg sync.WaitGroup
	wg.Add(len(l))

	ch := make(chan FreqMap, len(l))
	for idx, str := range l {
		go func(idx int, s string) {
			res := Frequency(s)
			fmt.Printf(" %d unique keys found for idx %d letter block \n", len(res), idx)
			fmt.Printf("# writing frq to number %d letter block \n", idx)
			ch <- res
			wg.Done()
		}(idx, str)
	}

	wg.Wait()
	close(ch)

	for fm := range ch {
		fmt.Printf("# reading frq with %d unique keys from chan \n", len(fm))
		// reduce to single map
		for k, v := range fm {
			m[k] += v
		}
	}
	return m
}

func ConcurrentFrq(txtBlocks []string) FreqMap {
	res := make(FreqMap)

	// frqChan := make(chan FreqMap, len(txtBlocks))
	frqChan := make(chan FreqMap) // unbuffered without size
	for idx, text := range txtBlocks {
		go func(i int, t string) {
			fmt.Printf("# writing frq count result to NO. %d chan block \n", i)
			frqChan <- Frequency(t)
		}(idx, text)
	}

	n := 0
	for range txtBlocks {
		for frqEach := range frqChan {
			n += 1
			fmt.Printf("# reading from NO. %d chan block \n", n)
			for k, count := range frqEach {
				res[k] += count
			}
		}
	}
	// for range txtBlocks {
	// 	for k, count := range <-frqChan {
	// 		res[k] += count
	// 	}
	// }

	return res
}
