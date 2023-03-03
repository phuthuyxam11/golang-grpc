package example

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func FindCountText() (int, error) {
	countChan := make(chan int)
	count := 0
	folder := "./tmp"
	entries, err := os.ReadDir(folder)

	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	go func() {
		for _, e := range entries {
			path, _ := filepath.Abs(fmt.Sprintf("%s/%s", folder, e.Name()))
			fileContent, err := ioutil.ReadFile(path)
			if err != nil {
				log.Println("can not read file: ", err)
				break
			}
			countChan <- 1
			count += strings.Count(string(fileContent), "lorem")
		}
		close(countChan)
	}()
	for _, _ = range entries {
		<-countChan
	}
	return count, nil
}
