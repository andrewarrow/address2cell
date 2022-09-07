package address

import (
	"fmt"
	"strings"
)

func Process(body string) {
	lines := strings.Split(body, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == "" {
			continue
		}
		fmt.Println(trimmedLine)
	}
}
