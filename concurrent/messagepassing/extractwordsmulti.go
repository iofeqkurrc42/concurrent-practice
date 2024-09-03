package messagepassing

import (
	"fmt"
	"sort"
	"strings"
)

const downloaders = 20

func longestWords(quit <-chan int, words <-chan string) <-chan string {
	longWords := make(chan string)
	go func() {
		defer close(longWords)
		uniqueWordsMap := make(map[string]bool)
		uniqueWords := make([]string, 0)
		moreData, word := true, ""
		for moreData {
			select {
			case word, moreData = <-words:
				if moreData && !uniqueWordsMap[word] {
					uniqueWordsMap[word] = true
					uniqueWords = append(uniqueWords, word)
				}
			case <-quit:
				return
			}
		}
		sort.Slice(uniqueWords, func(a, b int) bool {
			return len(uniqueWords[a]) > len(uniqueWords[b])
		})
		longWords <- strings.Join(uniqueWords[:10], ", ")
	}()
	return longWords
}

func Run20() {
	quit := make(chan int)
	defer close(quit)
	urls := generateUrl(quit)
	pages := make([]<-chan string, downloaders)
	for i := 0; i < downloaders; i++ {
		pages[i] = downloadPage(quit, urls)
	}
	results := extractWords(quit, FanIn(quit, pages...))
	for result := range results {
		fmt.Println(result)
	}
}

func Run21() {
	quit := make(chan int)
	defer close(quit)
	urls := generateUrl(quit)
	pages := make([]<-chan string, downloaders)

	for i := 0; i < downloaders; i++ {
		pages[i] = downloadPage(quit, urls)
	}
	results := longestWords(quit, extractWords(quit, FanIn(quit, pages...)))
	fmt.Println("Longest words: ", <-results)
}
