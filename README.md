# go-deepl

Go package for the [DeepL](https://www.deepl.com) translation API

## Example
```go
translations, err := deepl.Translate("Hello", "en", "de")
if err != nil {
	log.Fatal(err)
}

for _, translation := range translations {
	fmt.Printf(translation.Sentence)
}
```