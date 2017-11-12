package deepl

import "encoding/json"

// Translation contains a resulting translation
type Translation struct {
	Text        string
	Probability float32
}

func parseResponse(message []byte) ([]Translation, error) {
	type Beam struct {
		NumSymbols            int     `json:"num_symbols"`
		PostprocessedSentence string  `json:"postprocessed_sentence"`
		Score                 float32 `json:"score"`
		TotalLogProb          float32 `json:"totalLogProb"`
	}

	type ResultTranslation struct {
		Beams                    []Beam `json:"beams"`
		TimeAfterPreprocessing   int    `json:"timeAfterPreprocessing"`
		TimeReceivedFromEndpoint int    `json:"timeReceivedFromEndpoint"`
		TimeSentToEndpoint       int    `json:"timeSentToEndpoint"`
		TotalTimeEndpoint        int    `json:"total_time_endpoint"`
	}

	type Result struct {
		SourceLang            string              `json:"source_lang"`
		SourceLangIsConfident float32             `json:"source_lang_is_confident"`
		TargetLang            string              `json:"target_lang"`
		Translations          []ResultTranslation `json:"translations"`
	}

	type Response struct {
		ID      int    `json:"id"`
		JSONRPC string `json:"jsonrpc"`
		Result  Result `json:"result"`
	}

	var response Response
	err := json.Unmarshal(message, &response)
	if err != nil {
		return nil, err
	}

	beams := response.Result.Translations[0].Beams
	translations := make([]Translation, 0)
	for _, beam := range beams {
		translation := Translation{
			Text:        beam.PostprocessedSentence,
			Probability: beam.TotalLogProb,
		}
		translations = append(translations, translation)
	}

	return translations, nil
}
