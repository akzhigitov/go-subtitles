package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{"*"},
	}))

	e.Static("/", "public")
	e.POST("/upload", upload)

	lemmas = loadLemmas("lemmas_60k_words.txt", 61000)
	freqMap = parseEnglishWordsFreq("unigram_freq.csv")

	knownWordsMap = parseKnownWords("knownWords.txt")

	e.Logger.Fatal(e.Start(":1323"))
}

var lemmas map[string]string
var freqMap map[string]int
var knownWordsMap map[string]interface{}

func upload(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)

	words := map[string]Word{}
	phrase := ""
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && unicode.IsLetter(rune(line[0])) {
			phrase += line + " "
		} else if line == "" {

			for _, word := range parsePhrase(phrase) {
				if existWord, ok := words[word.Value]; ok {
					if len(existWord.Phrase) < len(word.Phrase) {
						words[word.Value] = word
					}
				} else {
					words[word.Value] = word
				}
			}
			phrase = ""
		}
	}

	result := Words{}
	for _, w := range words {
		if isKnownWord(w.Value) {
			result.KnownWords = append(result.KnownWords, w)
		} else if w.Freq > 0 {
			result.UnknownWords = append(result.UnknownWords, w)
		} else {
			result.BrokenWords = append(result.BrokenWords, w)
		}
	}

	return c.JSON(http.StatusOK, result)
}

func isKnownWord(w string) bool {
	_, ok := knownWordsMap[w]
	return ok
}

type Words struct {
	UnknownWords []Word `json:"unknownWords"`
	KnownWords   []Word `json:"knownWords"`
	BrokenWords  []Word `json:"brokenWords"`
}

type Word struct {
	Value  string `json:"value"`
	Freq   int    `json:"freq"`
	Phrase string `json:"phrase"`
}

func parsePhrase(phrase string) []Word {
	var result []Word

	for _, field := range strings.Fields(phrase) {

		word := trim(strings.ToLower(field))

		if strings.Contains(word, "'") {
			splitted := strings.Split(word, "'")
			word = splitted[0]
		}

		lemma, ok := lemmas[word]
		if !ok {
			lemma = word
		}
		result = append(result, Word{
			Value:  lemma,
			Freq:   freqMap[lemma],
			Phrase: phrase,
		})
	}

	return result
}

func trim(word string) string {
	return strings.Trim(strings.ReplaceAll(word, "</i>", ""), ".,!?-")
}

func loadLemmas(filepath string, size int) (result map[string]string) {

	result = make(map[string]string, size)
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lemma, word := parseLemma(scanner.Text())
		result[word] = lemma
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func parseLemma(line string) (lemma, word string) {
	// File headers:
	// lemRank	lemma	PoS	lemFreq	wordFreq	word

	fields := strings.Fields(line)
	return fields[1], fields[5]
}

func parseEnglishWordsFreq(filepath string) map[string]int {
	result := make(map[string]int)

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan() // skip header

	for scanner.Scan() {
		word, freq := parseFreq(scanner.Text())
		result[word] = freq
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func parseFreq(line string) (word string, freq int) {
	fields := strings.Split(line, ",")

	freq, err := strconv.Atoi(fields[1])
	if err != nil {
		log.Fatal(err)
	}

	return fields[0], freq
}

func parseKnownWords(filepath string) map[string]interface{} {
	result := make(map[string]interface{})

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		result[word] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}
