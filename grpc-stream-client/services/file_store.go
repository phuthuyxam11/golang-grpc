package services

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

type StoreImageInterface interface {
	Save(path string, fileName string, fileType string, byteChunk bytes.Buffer) error
}

type FileInfo struct {
	path     string
	fileName string
	fileSize int32
}

type DiskStoreFile struct {
	mutex    sync.Mutex
	folder   string
	fileInfo map[string]*FileInfo
}

func NewDiskStoreFile(folder string) *DiskStoreFile {
	return &DiskStoreFile{
		mutex:    sync.Mutex{},
		folder:   folder,
		fileInfo: nil,
	}
}

func (store *DiskStoreFile) Save(folder string, fileName string, fileType string, byteChunk bytes.Buffer) error {
	// create file
	imagePath, _ := filepath.Abs(fmt.Sprintf("./%s/%s_%s.%s", folder, fileName, strconv.FormatInt(time.Now().UTC().UnixNano(), 10), fileType))
	file, err := os.Create(imagePath)
	if err != nil {
		log.Fatalln("can't create file: ", err)
		return err
	}
	store.mutex.Lock()
	_, err = byteChunk.WriteTo(file)
	if err != nil {
		log.Fatalln("Can't write file: ", err)
		return err
	}
	defer store.mutex.Unlock()
	return nil
}
