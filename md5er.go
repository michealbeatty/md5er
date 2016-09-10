package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// func walkpath(path string, file os.FileInfo, err error) error {

// 	if !file.IsDir() {
// 		fileData, err := ioutil.ReadFile(path)
// 		if err != nil {
// 			log.Print(err)
// 		}
// 		fmt.Printf("%s\t%x\n", path, md5.Sum(fileData))
// 		//fmt.Println(path, file.Name())
// 	}
// 	return nil
// }

// func readDir(directory string) {

// 	files, err := ioutil.ReadDir(directory)
// 	fmt.Println(os.Getwd())
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	for _, file := range files {
// 		if file.IsDir() {
// 			readDir(file.Name())
// 		} else {
// 			fmt.Printf("%s\t", file.Name())
// 			fileData, err := ioutil.ReadFile(file.Name())
// 			if err != nil {
// 				log.Print(err)
// 			}
// 			fmt.Printf("%x\n", md5.Sum(fileData))
// 		}
// 	}
// }

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
