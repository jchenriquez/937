package main

import ("fmt"
		"regexp")

func reorderLogFiles(logs []string) (result []string) {
	letterRegex := regexp.MustCompile(`let[1-9] ([a-z]+( )?)+`)
	digitRegex := regexp.MustCompile(`dig[1-9] ([0-9]( )?)+`)

	result = make([]string, 0, len(logs))

	for _, log := range logs {
		if letterRegex.Match([]byte(log)) {
			result = append(result, log)
		}
	}

	for _, log := range logs {
		if digitRegex.Match([]byte(log)) {
			result = append(result, log)
		}
	}

	return
}

func main() {
	fmt.Printf("reodered %v\n", reorderLogFiles([]string{"dig1 8 1 5 1","let1 art can","dig2 3 6","let2 own kit dig","let3 art zero"}))
}
