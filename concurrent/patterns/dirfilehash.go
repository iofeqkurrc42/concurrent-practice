package patterns

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

func FHash(filepath string) []byte {
	file, _ := os.Open(filepath)
	defer file.Close()

	sha := sha256.New()
	io.Copy(sha, file)

	return sha.Sum(nil)
}

func Run() {
	dir := os.Args[1]
	files, _ := os.ReadDir(dir)
	wg := sync.WaitGroup{}
	for _, file := range files {
		if !file.IsDir() {
			wg.Add(1)
			go func(filename string) {
				fPath := filepath.Join(dir, filename)
				hash := FHash(fPath)
				fmt.Printf("%s - %x\n", filename, hash)
				wg.Done()
			}(file.Name())
		}
	}
	wg.Wait()
}

func Run1() {
	dir := os.Args[1]
	files, _ := os.ReadDir(dir)
	sha := sha256.New()
	var prev, next chan int
	for _, file := range files {
		if !file.IsDir() {
			next := make(chan int)
			go func(filename string, prev, next chan int) {
				fpath := filepath.Join(dir, filename)
				hashOnFile := FHash(fpath)
				if prev != nil {
					<-prev
				}
				sha.Write(hashOnFile)
				next <- 0
			}(file.Name(), prev, next)
			prev = next
		}
	}
	<-next
	fmt.Printf("%x\n", sha.Sum(nil))
}

// using waitgroup
func Run7() {
	dir := os.Args[1]
	files, _ := os.ReadDir(dir)
	hMd5 := md5.New()
	var prev, next *sync.WaitGroup
	for _, file := range files {
		if !file.IsDir() {
			next = &sync.WaitGroup{}
			next.Add(1)
			go func(filename string, prev, next *sync.WaitGroup) {
				fpath := filepath.Join(dir, filename)
				fmt.Println("Processing", fpath)
				hashOnFile := FHash(fpath)
				if prev != nil {
					prev.Wait()
				}
				hMd5.Write(hashOnFile)
				next.Done()
			}(file.Name(), prev, next)
			prev = next
		}
	}
	next.Wait()
	fmt.Printf("%s - %x\n", dir, hMd5.Sum(nil))
}
