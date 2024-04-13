package tail

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	FilePtr := flag.String("file", "", "File to be tailed.")
	flag.Parse()
	if *FilePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	file, err := os.Open(*FilePtr)
	if err != nil {
		log.Fatal("Could not Open file %s", *FilePtr)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// XXX: Without this sleep you would hogg the CPU
				// NOTE: How to make it better ?
				time.Sleep(500 * time.Millisecond)
				truncated, errTruncated := isTruncated(file)
				if errTruncated != nil {
					break
				}
				if truncated {
					_, errSeek := file.Seek(0, io.SeekStart)
					if errSeek != nil {
						break
					}
				}
				continue
			}
			log.Fatal("Could not read line. Jumping to next line.")
		}
		fmt.Println("%s\n", string(line))
	}
}

func isTruncated(file *os.File) (bool, error) {
	currentPosition, err := file.Seek(0, io.SeekCurrent)
	if err != nil {
		return false, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return false, err
	}
	return currentPosition > fileInfo.Size(), nil
}
