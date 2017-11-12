package deepl

import (
	"encoding/json"
	"fmt"
	"strings"
)

func createRequestBody(text string, sourceLang string, targetLang string) ([]byte, error) {
	sourceLang = strings.ToUpper(sourceLang)
	targetLang = strings.ToUpper(targetLang)

	if !SourceLangAvailable(sourceLang) {
		return nil, fmt.Errorf("source language %s is not supported by deepl ", sourceLang)
	}

	if !TargetLangAvailable(targetLang) {
		return nil, fmt.Errorf("target language %s is not supported by deepl ", targetLang)
	}

	type Job struct {
		Kind          string `json:"kind"`
		RawEnSentence string `json:"raw_en_sentence"`
	}

	type Lang struct {
		UserPreferredLangs     []string `json:"user_preferred_langs"`
		SourceLangUserSelected string   `json:"source_lang_user_selected"`
		TargetLang             string   `json:"target_lang"`
	}

	type Params struct {
		Jobs     []Job `json:"jobs"`
		Lang     Lang  `json:"lang"`
		Priority int   `json:"priority"`
	}

	type Request struct {
		JSONRPC string `json:"jsonrpc"`
		Method  string `json:"method"`
		Params  Params `json:"params"`
	}

	job := Job{
		Kind:          "default",
		RawEnSentence: text,
	}
	lang := Lang{
		UserPreferredLangs:     []string{sourceLang, targetLang},
		SourceLangUserSelected: sourceLang,
		TargetLang:             targetLang,
	}
	params := Params{
		Jobs:     []Job{job},
		Lang:     lang,
		Priority: -1,
	}
	content := Request{
		JSONRPC: "2.0",
		Method:  "LMT_handle_jobs",
		Params:  params,
	}

	return json.Marshal(content)
}
