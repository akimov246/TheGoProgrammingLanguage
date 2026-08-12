// Измените программу dup2 так, чтобы она выводила имена всех
// файлов, в которых найдены повторяющиеся строки.
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	counts := make(map[string]int)
	lineLocation := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, lineLocation)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "excercise1.4: %v\n", err)
				continue
			}
			countLines(f, counts, lineLocation)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%v\t%s\n", n, strings.Join(lineLocation[line], " "), line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, lineLocation map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !slices.Contains(lineLocation[line], f.Name()) {
			lineLocation[line] = append(lineLocation[line], f.Name())
		}
	}
	// Примечание: игнорируем потенциальные ошибки из input.Err()
}
