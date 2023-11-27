package main

import "github.com/plzfgme/clitrans/translators"

type TranslateFunc func(text string, from string, to string) (string, error)

var translatorMap = map[string]TranslateFunc{
	"google": translators.GoogleTranslate,
}
