package patterns

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type CodeDepth struct {
	file  string
	level int
}

func deepestNestBlock(filename string) CodeDepth {
	code, _ := os.ReadFile(filename)
	max := 0
	level := 0
	for _, c := range code {
		if c == '{' {
			level += 1
			max = int(math.Max(float64(max), float64(level)))
		} else if c == '}' {
			level -= 1
		}
	}
	return CodeDepth{filename, max}
}

func forkIfNeeded(path string, info os.FileInfo, wg *sync.WaitGroup, result chan CodeDepth) {
	if !info.IsDir() && strings.HasSuffix(path, ".go") {
		wg.Add(1)
		go func() {
			result <- deepestNestBlock(path)
			wg.Done()
		}()
	}
}

func joinResults(partialResult chan CodeDepth) chan CodeDepth {
	finalResult := make(chan CodeDepth)
	max := CodeDepth{"", 0}
	go func() {
		for pr := range partialResult {
			if pr.level > max.level {
				max = pr
			}
		}
		finalResult <- max
	}()
	return finalResult
}

func Run2() {
	dir := os.Args[1]
	_, err := os.Stat(dir)
	if err != nil {
		panic(err)
	}
	partialResults := make(chan CodeDepth)
	wg := sync.WaitGroup{}
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		forkIfNeeded(path, info, &wg, partialResults)
		return nil
	})
	finalResult := joinResults(partialResults)
	wg.Wait()

	close(partialResults)
	result := <-finalResult
	fmt.Printf("%s has the deepest nested code block of %d\n", result.file, result.level)
}
