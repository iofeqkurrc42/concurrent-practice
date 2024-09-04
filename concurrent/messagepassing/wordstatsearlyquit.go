package messagepassing

import "fmt"

func Run23() {
	quitWords := make(chan int)
	quit := make(chan int)
	defer close(quit)
	urls := generateUrl(quit)
	pages := make([]<-chan string, downloaders)
	for i := 0; i < downloaders; i++ {
		pages[i] = downloadPage(quitWords, urls)
	}
	words := Take(quitWords, 10000, extractWords(quitWords, FanIn(quitWords, pages...)))
	wordsMulti := Broadcast(quit, words, 2)
	longestResult := longestWords(quit, wordsMulti[0])
	frequentResult := frequentWords(quit, wordsMulti[1])

	fmt.Println("longest words:", <-longestResult)
	fmt.Println("most frequent words:", <-frequentResult)
}
