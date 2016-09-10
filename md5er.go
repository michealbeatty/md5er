package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//RecursiveHash will perform the MD5 operation recursively beginning with the
// directory in which the application is executed.
func recursiveHash(dir string, fout string) {
	f, err := os.Create(fout)
	if err != nil {
		log.Print(err)
	}
	defer f.Close()
	filepath.Walk(dir, func(path string, file os.FileInfo, _ error) error {
		if !file.IsDir() {
			fileData, err := ioutil.ReadFile(path)
			if err != nil {
				log.Print(err)
			}
			rel, err := filepath.Rel(dir, path)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%x *%s\n", md5.Sum(fileData), rel)
			value := md5.Sum(fileData)

			s := hex.EncodeToString(value[:]) + " *" + rel + "\n"
			f.WriteString(s)
		}
		return nil
	})
	f.Sync()
}

func cwdHash(dir string, fout string) {
	f, err := os.Create(fout)
	if err != nil {
		log.Print(err)
	}
	defer f.Close()

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Print(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			fileData, err := ioutil.ReadFile(file.Name())
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%x *%s\n", md5.Sum(fileData), file.Name())
			value := md5.Sum(fileData)

			s := hex.EncodeToString(value[:]) + " *" + file.Name() + "\n"
			f.WriteString(s)
		}
	}

	f.Sync()
}

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}
	base := filepath.Base(cwd)

	recursive := flag.Bool("r", false, "crawl the directory recursively")
	fname := flag.String("o", base+".md5", "name of output file (should have md5 extension")
	flag.Parse()
	if !strings.HasSuffix(*fname, ".md5") {
		*fname += ".md5"
		log.Print("Appending .md5 extension to given filename")
	}
	if *recursive {
		recursiveHash(cwd, *fname)
	} else {
		fmt.Println("not recursive")
		cwdHash(cwd, *fname)
	}
}
