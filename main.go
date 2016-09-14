package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var fp *os.File
	var err error
	fp, err = os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fp)
	var str []string
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fp.Close()
	// 書き出し

    fw, _ := os.Create(os.Args[1])
	for _, v := range str {
		if http := regexp.MustCompile("^http"); http.MatchString(v) {
			v = http.ReplaceAllString(v, "- http")
		} else if title := regexp.MustCompile(`^[^\-\#]`); title.MatchString(v) {
			rep1 := regexp.MustCompile(`[\[\]～－]`)
			v = rep1.ReplaceAllString(v, "")

			rep2 := regexp.MustCompile(`[\s\,]`)
			v = rep2.ReplaceAllString(v, "_")

			v = "## " + v
		}
		fmt.Println(v)
        fw.Write([]byte( v + "\n"))
	}
    fw.Close()
}
