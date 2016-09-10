package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("This is the md5 program")
	cwd, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}
	filepath.Walk(cwd, func (path string, file os.FileInfo, err error) error {
		if !file.IsDir() {
			fileData, err := ioutil.ReadFile(path)
			if err != nil {
				log.Print(err)
			}
			rel, err := filepath.Rel(cwd, path)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%x\t*%s\n", md5.Sum(fileData), rel)
			//fmt.Println(path, file.Name())
	}
	return nil
})
}
