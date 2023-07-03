package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"unicode/utf8"
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
	delimiterFlag := flag.String("d", ",", "delimiter")
	outputFlag := flag.String("o", os.Stdout.Name(), "output file")
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "usage: ./fiter [-flags] [dir]")
		flag.CommandLine.PrintDefaults()
		fmt.Println("  [dir]")
		fmt.Println("\tpath to directory")
	}
	flag.Parse()
	delimiter, _ := utf8.DecodeRuneInString(*delimiterFlag)
	output, err := os.OpenFile(*outputFlag, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	path := flag.Arg(0)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal("invalid path")
	}
	log.Println("start")
	Write(path, output, delimiter)
	output.Close()
	log.Println("done")

}

// Write calls Walk and writes info to output
func Write(root string, output *os.File, delimiter rune) {
	entries := Walk(root)
	w := csv.NewWriter(output)
	w.Comma = delimiter
	err := w.Write([]string{"path", "datetime", "length"})
	if err != nil {
		log.Println(err)
	}
	for entry := range entries {
		err := w.Write(entry)
		if err != nil {
			log.Println(err)
		}
	}
	w.Flush()
}
