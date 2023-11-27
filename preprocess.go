package main

import (
	"strings"
)

type PreprocessFunc func(text string) string

var preprocessorMap = map[string]PreprocessFunc{
	"remove_newlines": removeNewlines,
}

func removeNewlines(text string) string {
	text = strings.ReplaceAll(text, "\r\n", " ")
	text = strings.ReplaceAll(text, "\n", " ")

	return text
}
