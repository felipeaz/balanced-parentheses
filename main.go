package main

import (
	"log"
)

var (
	openSignalMapper = map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
)

func main() {
	inputs := []string{
		// balanced
		"()",
		// balanced
		"([]){}",
		// not balanced
		"([",
		// not balanced
		"([]{}]",
	}
	for _, i := range inputs {
		b := isBalanced(i)
		log.Printf("%s balanced? %v\n", i, b)
	}
}

func isBalanced(str string) bool {
	var toClose []string
	for _, ch := range str {
		chStr := string(ch)
		_, ok := openSignalMapper[chStr]
		if !ok {
			if len(toClose) == 0 {
				return false
			}
			// F.I.L.A - First In, Last Out
			numOfOpenedSignals := len(toClose) - 1
			expectedSignal, ok := openSignalMapper[toClose[numOfOpenedSignals]]
			if !ok {
				return false
			}
			if expectedSignal != chStr {
				return false
			}
			// remove the last element
			toClose = toClose[:numOfOpenedSignals]
			continue
		}
		toClose = append(toClose, chStr)
	}
	return len(toClose) == 0
}
