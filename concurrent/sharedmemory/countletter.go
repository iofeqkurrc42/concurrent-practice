package sharedmemory

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const allLetters = "abcdefghijklmnopqrstuvwxyz"

func countLetters(url string, frequency []int) {
	resp, _ := http.Get(url)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()
	if resp.StatusCode != 200 {
		panic("Server returning error status code:" + resp.Status)
	}

	body, _ := io.ReadAll(resp.Body)
	for _, b := range body {
		c := strings.ToLower(string(b))
		cIndex := strings.Index(allLetters, c)
		if cIndex >= 0 {
			frequency[cIndex] += 1
		}
	}

	fmt.Println("Completed: ", url)
}

func Run1() {
	frequecey := make([]int, 26)
	for i := 1000; i <= 1030; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		// it is exist race condition
		go countLetters(url, frequecey)
	}
	time.Sleep(10 * time.Second)
	for i, c := range allLetters {
		fmt.Printf("%c-%d", c, frequecey[i])
	}
}
