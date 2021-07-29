package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	file := flag.String("f", "/home/ming/learn/go/pingcap-projects/tidb/util/stmtsummary/evicted_test.go", "file path to read from")
	flag.Parse()

	f, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Text()
		//line = process(line)
		fmt.Println(line)
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
