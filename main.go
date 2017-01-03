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
		replace_str := replace_url(v)
		fmt.Println(replace_str)
		fw.Write([]byte(replace_str + "\n"))
	}
	fw.Close()
}

func replace_url(str string) string {
	if http := regexp.MustCompile("^http"); http.MatchString(str) {
		return http.ReplaceAllString(str, "- http")
	} else if title := regexp.MustCompile(`^[^\-\#]`); title.MatchString(str) {
		new_str := str
		rep1 := regexp.MustCompile(`[\[\]～－\/\.　\?\!\:\(\)]`)
		new_str = rep1.ReplaceAllString(new_str, "")

		rep2 := regexp.MustCompile(`[\s\,]+`)
		new_str = rep2.ReplaceAllString(new_str, "_")

		new_str = "## " + new_str + ".mp4" + "\n"
		return new_str
	}
	return str
}
