package tools

import "strings"

func WCount(content string) int  {
	return len(strings.Split(content, " "))
}
