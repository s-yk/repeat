package main

import (
	"fmt"
	"regexp"
)

func repeat(chars string, repeatCount int) string {
	base := fmt.Sprintf("%*s", repeatCount, "")
	r := regexp.MustCompile(` `)
	return r.ReplaceAllString(base, chars)
}
