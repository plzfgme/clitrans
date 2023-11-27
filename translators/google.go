package translators

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

func GoogleTranslate(text string, from string, to string) (string, error) {
	url := fmt.Sprintf("https://translate.google.com/m?sl=%s&tl=%s&q=%s", from, to, url.QueryEscape(text))
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	page, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`(?s)<div class="result-container">(.+?)</div>`)
	caps := re.FindStringSubmatch(string(page))
	if len(caps) < 2 {
		return "", errors.New("Could not find translation in response")
	}

	return caps[1], nil
}
