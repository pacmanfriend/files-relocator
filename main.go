package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"strings"
)

func main() {
	dstDir := flag.String("ddir", "", "destination directory")
	srcDir := flag.String("sdir", "", "source directory")
	fileExt := flag.String("ext", "", "file extension")

	flag.Parse()

	if *dstDir == "" {
		log.Fatal("destination directory is empty")
	}

	if *srcDir == "" {
		log.Fatal("source directory is empty")
	}

	if *fileExt == "" {
		log.Fatal("file extension is empty")
	}

	err := filepath.Walk(*srcDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if strings.Index(path, "") != -1 {
				fmt.Println(path)
			}

			//files, err := filepath.Glob(filepath.Join(path, "*"+*fileExt))
			//if err != nil {
			//	log.Fatal(err)
			//}
			//
			//if len(files) > 0 {
			//	fmt.Println(path)
			//}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
