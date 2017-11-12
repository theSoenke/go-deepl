package deepl

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

// Translate text with deepl.com
func Translate(text string, sourceLang string, targetLang string) ([]Translation, error) {
	url := "https://www.deepl.com/jsonrpc"
	reqBody, err := createRequestBody(text, sourceLang, targetLang)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	translations, err := parseResponse(respBody)
	return translations, err
}

// LangAvailable returns true when language code is supported
func LangAvailable(lang string) bool {
	supported := []string{"DE", "EN", "FR", "ES", "IT", "NL", "PL"}
	lang = strings.ToUpper(lang)
	for _, l := range supported {
		if l == lang {
			return true
		}
	}

	return false
}
