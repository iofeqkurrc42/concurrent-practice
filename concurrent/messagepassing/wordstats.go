package messagepassing

import (
	"fmt"
	"sort"
	"strings"
)

func frequentWords(quit <-chan int, words <-chan string) <-chan string {
	mostFrequentWords := make(chan string)
	go func() {
		defer close(mostFrequentWords)
		freqMap := make(map[string]int)
		freqList := make([]string, 0)
		moreData, word := true, ""
		for moreData {
			select {
			case word, moreData = <-words:
				if moreData {
					if freqMap[word] == 0 {
						freqList = append(freqList, word)
					}
					freqMap[word] += 1
				}
			case <-quit:
				return
			}
		}
		sort.Slice(freqList, func(a, b int) bool {
			return freqMap[freqList[a]] > freqMap[freqList[b]]
		})
		mostFrequentWords <- strings.Join(freqList[:10], ", ")
	}()
	return mostFrequentWords
}

const dowloaders = 20

func Run22() {
	quit := make(chan int)
	defer close(quit)
	urls := generateUrl(quit)
	pages := make([]<-chan string, dowloaders)
	for i := 0; i < dowloaders; i++ {
		pages[i] = downloadPage(quit, urls)
	}
	words := extractWords(quit, FanIn(quit, pages...))
	wordsMulti := Broadcast(quit, words, 2)
	longestResult := longestWords(quit, wordsMulti[0])
	frequentResult := frequentWords(quit, wordsMulti[1])
	fmt.Println("longest words:", <-longestResult)
	fmt.Println("most frequent words:", <-frequentResult)
}
