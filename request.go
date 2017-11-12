package deepl

import "encoding/json"
import "strings"

func createRequestBody(text string, fromLang string, toLang string) ([]byte, error) {
	fromLang = strings.ToUpper(fromLang)
	toLang = strings.ToUpper(toLang)

	type Job struct {
		Kind          string `json:"kind"`
		RawEnSentence string `json:"raw_en_sentence"`
	}

	type Lang struct {
		UserPreferredLangs   []string `json:"user_preferred_langs"`
		SourceLangUserSelect string   `json:"source_lang_user_selected"`
		TargetLang           string   `json:"target_lang"`
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
		UserPreferredLangs:   []string{fromLang, toLang},
		SourceLangUserSelect: fromLang,
		TargetLang:           toLang,
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
