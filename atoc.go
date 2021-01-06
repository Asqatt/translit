package translit

import (
	"fmt"
	"strings"
	"unicode"
)

var changetoj = map[rune]rune{
	'Ә': 'А',
	'О': 'Ө',
	'Ұ': 'Ү',
	'Ы': 'І',
}

var jVowels = map[rune]bool{
	'گ': true,
	'ە': true,
	'ك': true,
}

var vowels = map[rune]bool{
	'ا': true,
	'و': true,
	'ۇ': true,
	'ە': true,
}
var atoc = map[rune]rune{
	'ا': 'А',
	'ب': 'Б',
	'ۆ': 'В',
	'گ': 'Г',
	'ع': 'Ғ',
	'د': 'Д',
	'ە': 'Е',
	'ج': 'Ж',
	'ز': 'З',
	'ي': 'И',
	//'ي':'Й',
	'ك': 'К',
	'ق': 'Қ',
	'ل': 'Л',
	'م': 'М',
	'ن': 'Н',
	'ڭ': 'Ң',
	'و': 'О',
	'پ': 'П',
	'ر': 'Р',
	'س': 'С',
	'ت': 'Т',
	'ۋ': 'У',
	'ۇ': 'Ұ',
	'ف': 'Ф',
	'ح': 'Х',
	'ھ': 'Һ',
	//'س':'Ц',
	'چ': 'Ч',
	'ش': 'Ш',
	//'ش':'Щ',
	'ﯼ': 'Ы',
	'ى': 'Ы',
	// 'ا':'Э',
	'،': ',',
	'؟': '?',
	'؛': ';',
	'！': '!',
}

//ArabicToCyrillic convertes Kazakh Arabic script to Cyrillc one.
func ArabicToCyrillic(s string) string {
	s = strings.ReplaceAll(s, "يا", "Я")
	s = strings.ReplaceAll(s, "يۋ", "Ю")
	var sentences []string
	withdots := strings.SplitAfter(s, ".")
	temp := make([]string, 0, len(withdots))
	for _, dotted := range withdots {
		for _, withquestion := range strings.SplitAfter(dotted, "؟") {
			temp = append(temp, withquestion)
		}
	}
	for _, unexclamed := range temp {
		for _, exclamed := range strings.SplitAfter(unexclamed, "！") {
			sentences = append(sentences, exclamed)
		}
	}

	var cyrillicSentences strings.Builder
	for _, sentence := range sentences {

		var buf []string
		for i, word := range strings.Split(strings.Trim(sentence, " "), " ") {
			var builder strings.Builder
			needHarmonify := checkNeedHaromify(&word)
			var previous rune
			for j, char := range word {
				var cyrillicChar rune

				if char == 'ي' {
					if vowels[previous] {
						cyrillicChar = 'Й'
					} else {
						cyrillicChar = 'И'
					}
				} else {
					cyrillicChar = getCyillicChar(char)
				}
				previous = char
				if i == 0 && j == 0 {
					builder.WriteRune(cyrillicChar)
				} else {
					builder.WriteRune(unicode.ToLower(cyrillicChar))
				}
			}
			var cyrillicWord string
			if needHarmonify {
				cyrillicWord = strings.Map(func(char rune) rune {
					if jV, ok := changetoj[char]; ok {
						return jV
					}
					return char
				}, builder.String())
			} else {
				cyrillicWord = builder.String()
			}

			buf = append(buf, cyrillicWord)
		}
		ci := strings.Join(buf, " ")
		fmt.Println(ci)
		cyrillicSentences.WriteString(strings.Join(buf, " "))
	}
	return cyrillicSentences.String()
}

func checkNeedHaromify(word *string) bool {
	if strings.HasPrefix(*word, "ء") {
		*word = strings.TrimPrefix(*word, "ء")
		return true
	}
	for _, char := range *word {
		if jVowels[char] {
			return true
		}
	}
	return false
}

func getCyillicChar(char rune) rune {
	if cyrill, ok := atoc[char]; ok {
		return cyrill
	}
	return char
}
