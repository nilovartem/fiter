package main

import (
	"encoding/csv"
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// Walk searches for file in every directory and sends data to channel
func Walk(path string) chan []string {
	entries := make(chan []string)
	go func() {
		filepath.WalkDir(path, func(path string, entry fs.DirEntry, _ error) error {
			if entry.Type().IsDir() {
				return nil
			}
			info, err := entry.Info()
			if err != nil {
				return filepath.SkipDir
			}
			res := []string{path, (info.ModTime().Format("2006-01-02 15:04:05")), strconv.Itoa(int(info.Size()))}
			entries <- res
			return nil
		})
		defer close(entries)
	}()
	return entries
}
func main() {
	out := flag.String("o", "", "output file")
	flag.Parse()
	var f *os.File = os.Stdout
	if *out != "" {
		var err error
		f, err = os.Create(*out)
		if err != nil {
			log.Fatal(err)
		}
	}
	if flag.Arg(0) == "" {
		log.Fatal("provide path to directory")
	}
	log.Println("starting...")
	Write(flag.Arg(0), f)
	f.Close()
	log.Println("done")
}

// Write calls Walk and writes info to output
func Write(root string, output *os.File) {
	entries := Walk(root)
	w := csv.NewWriter(output)
	for entry := range entries {
		err := w.Write(entry)
		if err != nil {
			log.Println(err)
		}
	}
	w.Flush()
}
