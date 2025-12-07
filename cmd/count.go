package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := map[string]int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		counts[text] = counts[text] + 1
	}
	if scanner.Err() != nil {
		_, err := fmt.Fprintln(os.Stderr, scanner.Err())
		if err != nil {
			fmt.Fprintln(os.Stderr, scanner.Err())
		}
		os.Exit(1)
	}
	for k, v := range counts {
		fmt.Printf("%v \t %v \n", k, v)
	}
}
