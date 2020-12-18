package main
import(
	_ "net/http"
	"unicode/utf8"
	"strings"
	"fmt"
) 

const (
	CHECKMARK = 0
	CHECKDROP =1
)


var mapping = map[rune]string {
	           'А':  "ا",
                 'Ә':  "ا",
                 'Б':  "ب",
                 'В':  "ۆ",
                 'Г':  "گ",
                 'Ғ':  "ع",
                 'Д':  "د",
                 'Е':  "ە",
                 'Ё':  "يە",
                 'Ж':  "ج",
                 'З':  "ز",
                 'И':  "ي",
                 'Й':  "ي",
                 'К':  "ك",
                 'Қ':  "ق",
                 'Л':  "ل",
                 'М':  "م",
                 'Н':  "ن",
                 'Ң':  "ڭ",
                 'О':  "و",
                 'Ө':  "و",
                 'П':  "پ",
                 'Р':  "ر",
                 'С':  "س",
                 'Т':  "ت",
                 'У':  "ۋ",
                 'Ү':  "ۇ",
                 'Ұ':  "ۇ",
                 'Ф':  "ف",
                 'Х':  "ح",
                 'Һ':  "ھ",
                 'Ц':  "س",
                 'Ч':  "چ",
                 'Ъ':  "",
                 'Ш':  "ش",
                 'Щ':  "ش",
                 'Ы':  "ﯼ",
                 'І':  "ﯼ",
                 'Ь':  "",
                 'Э':  "ا",
                 'Ю':  "يؤ",
				 'Я':  "يا",
				 ',':  "،",
				 '?': "؟",
				 ';': "؛",
				 '!':"！",
}
var jSet = map[rune]bool {'Ә':true, 'І':true, 'Ө':true, 'Ү':true, 'Э':true}
var dropMark = map[rune]bool{'Ё':true, 'К':true, 'Г':true, 'Е':true}



func CyrillicToArabic(s string) string  {
 var buf []string
 for _, word := range strings.Split(strings.ToUpper(s)," ") {
	 var b strings.Builder
	 var mark  bool  
	 var drop bool
	 for _, uc := range word {
		 
		 if !mark &&  checkMark(uc,CHECKMARK){
			 mark = true
		 }
		 if !drop &&checkMark(uc,CHECKDROP){
			 drop =true
		 }
		 b.WriteString(getArabicChar(uc))
	 }
	 //fmt.Println(b.String())
	 word = b.String()
	 if mark&&!drop {
		 word="ء"+word
	 }
	 buf=append(buf,word)
	 fmt.Println(word,utf8.RuneCountInString(word))
	 
 }
 return strings.Join(buf," ")
}

func checkMark(char rune , flag int) bool {
	if flag==CHECKMARK{
		 //fmt.Printf("mark get %c %q!",char,jSet[char])
	  return jSet[char]
	}else{
		return dropMark[char]

	}
}


//returns arabic from the mapping or return the original when encounters 
//normal non-kazakh literals like numbers, special chars
func getArabicChar(char rune) string {
	if ar, ok := mapping[char]; ok{
		return ar
	}
	return string(char)
}

func main() {
	s :="Бүкіл өмірін балаларына арнап, өмірлік серігіне адал жар, ардақты ана бола білген, бүгінде қызын -қияға, ұлын – ұяға қондырып, немерелерінің сүйікті Әжесі болып отырған асыл Анамыз, сізді бүгінгі мерейтойыңызбен шын жүректен құттықтаймыз!"
	fmt.Println(CyrillicToArabic(s))
}

