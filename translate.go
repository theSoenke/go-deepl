package deepl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

// Translate text with deepl.com
func Translate(text string, fromLang string, toLang string) ([]Translation, error) {
	url := "https://www.deepl.com/jsonrpc"
	reqBody, err := createRequestBody(text, fromLang, toLang)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	translations, err := parseResponse(respBody)
	return translations, err
}

// ToLangAvailable returns true when language code is a supported target language
func ToLangAvailable(lang string) bool {
	supported := []string{"DE", "EN", "FR", "ES", "IT", "NL", "PL"}
	lang = strings.ToUpper(lang)
	for _, l := range supported {
		if l == lang {
			return true
		}
	}

	return false
}

// FromLangAvailable returns true when language code is a supported source language
func FromLangAvailable(lang string) bool {
	supported := []string{"AUTO", "DE", "EN", "FR", "ES", "IT", "NL", "PL"}
	lang = strings.ToUpper(lang)
	for _, l := range supported {
		if l == lang {
			return true
		}
	}

	return false
}
