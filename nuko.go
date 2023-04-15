package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var showLineNum = flag.Bool("n", false, "行番号を表示する")

func main() {
	flag.Parse()

	cnt := 0
	for _, arg := range flag.Args() {
		content, err := readFile(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, "読み込みに失敗しました:", err)
			os.Exit(1)
		}
		for _, text := range content {
			cnt++
			if *showLineNum {
				text = fmt.Sprintf("%d: %s", cnt, text)
			}
			fmt.Println(text)
		}
	}
}

func readFile(fn string) ([]string, error) {
	content := []string{}
	f, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		content = append(content, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return content, nil
}
