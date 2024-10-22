package deadlock

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func handleDirectories(dirs <-chan string, files chan<- string) {
	for fullpath := range dirs {
		fmt.Println("Reading all files from", fullpath)
		filesInDir, _ := os.ReadDir(fullpath)
		fmt.Println("Pushing %d files from %s\n", len(filesInDir), fullpath)
		for _, file := range filesInDir {
			files <- filepath.Join(fullpath, file.Name())
		}
	}
}

func handleFiles(files chan string, dirs chan string) {
	for path := range files {
		file, _ := os.Open(path)
		fileInfo, _ := file.Stat()
		if fileInfo.IsDir() {
			fmt.Printf("Pushingj %s directory\n", fileInfo.Name())
			dirs <- path
		} else {
			fmt.Printf("File %s, size: %dMB, last modifed: %s\n",
				fileInfo.Name(),
				fileInfo.Size()/(1024*1024),
				fileInfo.ModTime().Format("15:04:05"))
		}
	}
}

func Run6() {
	filesChannel := make(chan string)
	dirsChannel := make(chan string)
	go handleFiles(filesChannel, dirsChannel)
	go handleDirectories(dirsChannel, filesChannel)
	dirsChannel <- os.Args[1]
	time.Sleep(60 * time.Second)
}

// using a separate goroutine to write on a channel
func handleFiles2(files chan string, dirs chan string) {
	for fullpath := range dirs {
		fmt.Println("Reading all files from", fullpath)
		filesInDir, _ := os.ReadDir(fullpath)
		fmt.Printf("Pushing %d files from %s\n", len(filesInDir), fullpath)
		for _, file := range filesInDir {
			// starts new goroutine that sends each file to the files channel
			go func(fp string) {
				files <- fp
			}(filepath.Join(fullpath, file.Name()))
		}
	}
}

// using select to break the circular wait
func handleDirectories3(dirs <-chan string, files chan<- string) {
	// create a slice to store files that need to be pushed to the file handler's channel
	toPush := make([]string, 0)
	appendAllFiles := func(path string) {
		fmt.Println("Reading all files from", path)
		fileInDir, _ := os.ReadDir(path)
		fmt.Printf("Pushing %d files from %s\n", len(fileInDir), path)
		// appends all files in a directory to the slice
		for _, f := range fileInDir {
			toPush = append(toPush, filepath.Join(path, f.Name()))
		}
	}
	for {
		// if there are no files to push, reads directory from the input channel and adds all files in the directory
		if len(toPush) == 0 {
			appendAllFiles(<-dirs)
		} else {
			select {
			// reads the next directory from the input channel and adds all files in the directory
			case fullpath := <-dirs:
				appendAllFiles(fullpath)
				// pushes the first file on the slice to the channel
			case files <- toPush[0]:
				// removes the first file from the slice
				toPush = toPush[1:]
			}
		}
	}
}
