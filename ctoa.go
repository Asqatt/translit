package translit

import (
	"strings"
)

const (
	//const flags
	fCHECKMARK = 0
	fCHECKDROP = 1
)

var mapping = map[rune]string{
	'А': "ا",
	'Ә': "ا",
	'Б': "ب",
	'В': "ۆ",
	'Г': "گ",
	'Ғ': "ع",
	'Д': "د",
	'Е': "ە",
	'Ё': "يە",
	'Ж': "ج",
	'З': "ز",
	'И': "ي",
	'Й': "ي",
	'К': "ك",
	'Қ': "ق",
	'Л': "ل",
	'М': "م",
	'Н': "ن",
	'Ң': "ڭ",
	'О': "و",
	'Ө': "و",
	'П': "پ",
	'Р': "ر",
	'С': "س",
	'Т': "ت",
	'У': "ۋ",
	'Ү': "ۇ",
	'Ұ': "ۇ",
	'Ф': "ف",
	'Х': "ح",
	'Һ': "ھ",
	'Ц': "س",
	'Ч': "چ",
	'Ъ': "",
	'Ш': "ش",
	'Щ': "ش",
	'Ы': "ﯼ",
	'І': "ﯼ",
	'Ь': "",
	'Э': "ا",
	'Ю': "يؤ",
	'Я': "يا",
	',': "،",
	'?': "؟",
	';': "؛",
	'!': "！",
}
var jSet = map[rune]bool{'Ә': true, 'І': true, 'Ө': true, 'Ү': true, 'Э': true}
var dropMark = map[rune]bool{'Ё': true, 'К': true, 'Г': true, 'Е': true}

//CyrillicToArabic convertes Kazakh Cyrillic scripts to Arabic one
func CyrillicToArabic(s string) string {
	var buf []string
	for _, word := range strings.Split(strings.ToUpper(s), " ") {
		var b strings.Builder
		var mark bool
		var drop bool
		for _, uc := range word {

			if !mark && checkMark(uc, fCHECKMARK) {
				mark = true
			}
			if !drop && checkMark(uc, fCHECKDROP) {
				drop = true
			}
			b.WriteString(getArabicChar(uc))
		}
		//fmt.Println(b.String())
		word = b.String()
		if mark && !drop {
			word = "ء" + word
		}
		buf = append(buf, word)
	}
	return strings.Join(buf, " ")
}

func checkMark(char rune, flag int) bool {
	if flag == fCHECKMARK {
		//fmt.Printf("mark get %c %q!",char,jSet[char])
		return jSet[char]
	}

	return dropMark[char]
}

//returns arabic from the mapping or return the original when encounters
//normal non-kazakh literals like numbers, special chars
func getArabicChar(char rune) string {
	if ar, ok := mapping[char]; ok {
		return ar
	}
	return string(char)
}
