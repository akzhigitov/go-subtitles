package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
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

	e.Logger.Fatal(e.Start(":1323"))
}

var lemmas map[string]string
var freqMap map[string]int

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

	b, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}
	words := parseText(string(b))
	result := Words{}
	for _, w := range words {

		freq := freqMap[w]

		if isKnownWord(w) {
			result.KnownWords = append(result.KnownWords, Word{
				Value: w,
				Freq:  freq,
			})
		} else if freq > 0 {
			result.UnknownWords = append(result.UnknownWords, Word{
				Value: w,
				Freq:  freq})
		} else {
			result.BrokenWords = append(result.BrokenWords, Word{
				Value: w,
				Freq:  freq})
		}
	}

	return c.JSON(http.StatusOK, result)
}

func isKnownWord(w string) bool {
	switch w {
	case "I", "you":
		return true
	}

	return false
}

type Words struct {
	UnknownWords []Word `json:"unknownWords"`
	KnownWords   []Word `json:"knownWords"`
	BrokenWords  []Word `json:"brokenWords"`
}

type Word struct {
	Value string `json:"value"`
	Freq  int    `json:"freq"`
}

func parseText(text string) []string {
	var result []string
	words := map[string]struct{}{}
	for _, field := range strings.Fields(text) {
		if unicode.IsLetter(rune(field[0])) {
			word := trim(strings.ToLower(field))

			if strings.Contains(word, "'") {
				splitted := strings.Split(word, "'")
				word = splitted[0]
			}

			lemma, ok := lemmas[word]
			if !ok {
				lemma = word
			}

			if _, ok := words[lemma]; !ok {
				words[lemma] = struct{}{}
				result = append(result, lemma)
			}
		}
	}

	sort.Strings(result)
	return result
}

func trim(word string) string {
	return strings.ReplaceAll(strings.Trim(word, ".,!?-"), "</i>", "")
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
