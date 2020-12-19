
package main

import (
	"strings"
	"os"
	"unicode"
	"fmt"
)

func main() {
	f, err := os.Create("out.md")
	if err!=nil {
		panic(err)
	}
	defer f.Close()
	a:= ` پياز بارشامىزعا تانىس، باعاسى ارزان، كۇندەلىكتى تۇرمىستا كوپ ىستەتىلەتىن، قورەكتىك قۇنى ەرەكشە جوعارى كوكونىستىڭ ءبىرى. ول قان تامىردى جۇمساتىپ قانا قالماستان، قاننىڭ جابىسقاقتىعىن دا تومەندەتە الادى. الايدا، ونى تۋراعاندا كوزدى اشىتىپ جاس شىعاراتىندىقتان، كوپتەگەن ادام ونى ىستەتۋدى قالامايدى. دەسەدە، ەۆروپا مەن امەريكادا ونىڭ « كوكونىس پاتشاسى » دەيتىن اتاعى بار بولىپ، ورتا جاستاعىلار مەن قارتتاردىڭ كۇندەلىكتى تۇتىنۋىنا ۇيلەسەدى. ەندەشە، بۇگىن پيازعا قاتىستى بىلىمدەرمەن تانىسالىق.
 
 1. كالتسيدى تولىقتاپ، سۇيەكتى بەكەمدەيدى
  پيازدىڭ قۇرامىندا كالتسي مول بولىپ، مايكوك، اسجاپىراق سياقتى كالتسي قۇرامى مول كوكونىستەرگە ۇقساپ كەتەدى. ونىڭ ۇستىنە، قۇرامىندا ماگني مەن كالي دا بار، كالتسي مەن فوسفوردىڭ سالىستىرماسى دا لايىقتى بولىپ،كالتسيدى تولىقتاۋمەن بىرگە،وڭاي سىمىرىلەدى.

2. ىشەك جولدارىن جاقسارتادى
    پياز تۇيەجاپىراق، قويان ءشوپ سياقتى كوكونىستەرمەن ۇقساس قۇرامىندا وليگوزا ساحاريد مول بولىپ، اسقازان مەن ىشەكتە وڭاي قورىتىلمايدى، الايدا، توق ىشەككە بارعاندا قوس اشالى تاياقشا باكتەريا سياقتى پايدالى باكتەريالاردىڭ تاماشا ازىعىنا اينالادى.

3. ءىش قاتۋدىڭ الدىن الىپ، ارىقتاتۋ رولى بار
پيازدىڭ قۇرامىندا ازىقتىق تالشىق مول بولىپ، تابەتتى تەجەپ، دارەتتى ايدايدى، مۇنان تىس، پيازدىڭ قۇرامىنداعى قوس اشالى نەگىز ىشەك جولدارىن مايدالايدى.
4. مايدى ىدىراتادى
مايلى تاعام جەگەندە، ماسەلەن، مايعا قارىلعان ەت، قاقتالعان بالىق سياقتىلاردى جەگەندە پياز قوسىپ جەسە، كۇلىمسى ءيىستى جويىپ، مايدى ىدىراتىپ، مايدىڭ ءسىمىرىلۋ مولشەرىن تومەندەتەدى.

5. باكتەريالاردى جويادى
پيازدىڭ قۇرامىنداعى ۇشپا كۇكىرتتى پروپيلەن ميكروبتى ءولتىرۋ، تەجەۋ رولىنا يە بولىپ، زياندى جاندىكتەردەن ساقتايدى، سوندىقتان پياز اۋىل شارۋاشىلىق ءدارىسىن كوپ قاجەت ەتپەيتىن ناعىز زالالسىز كوكونىس ەسەپتەلەدى.

6. راك اۋرۋىنان ساقتايدى
پيازدىڭ قۇرامىندا كۆەرتسيتين مەن سەلەن مول بولادى. كۆەرتسيتين راك كلەتكاسىنىڭ اكتيۆتىگىن تەجەپ، راك كلەتكاسىنىڭ ءوسىپ-ءوربۋىن شەكتەيدى، سەلەن كۇشتى انتيوكسيگەن ەسەپتەلەتىندىكتەن، دەنەنىڭ توتىعىپ زاقىمدالۋىن ازايتادى.

7. تۇماۋدىڭ الدىن الادى
پياز ءدامى قىشقىل، قاسيەتى جىلى كوكونىس بولىپ، سۋىقتى تاراتۋ رولى بار. تۇماۋدىڭ الدىن الادى ءارى تۇماۋدان تۋىنداعان مۇرىننىڭ ءبىتۋى، مۇرىننان سۋ اعۋ سياقتى جاعدايلاردى جازادى.

8. دەنەنىڭ قابىنۋعا بولعان سەزىمتالدىعىن باسەڭدەتەدى
پيازدىڭ قۇرامىنداعى پروستاگلاندين a نىڭ تامىردى كەڭەيتۋ رولى بار بولىپ، دەنەنىڭ قابىنۋعا بولعان سەزىمتالدىعىن باسەڭدەتۋ، قان قىسىمىن مەڭگەرۋ جانە قاننىڭ جابىسقاقتىعىن تومەندەتۋدە كورنەكتى ونىمدىلىگى بار.
               اۋدارعان ـ احات سانياز ۇلى
                                   كەلۋ قاينارى: توراپتان  الىندى`
f.WriteString(ArabicToCyrillic(a))
	// s :="Бүкіл өмірін балаларына арнап, өмірлік серігіне адал жар, ардақты ана бола білген, бүгінде қызын -қияға, ұлын – ұяға қондырып, немерелерінің сүйікті Әжесі болып отырған асыл Анамыз, сізді бүгінгі мерейтойыңызбен шын жүректен құттықтаймыз!"
	// f.WriteString(CyrillicToArabic(s))
}


var changetoj =map[rune]rune{
	'Ә' :'А',
	'О':'Ө',
'Ұ':'Ү',
'Ы':'І',
}

var jVowels =map[rune]bool{
	'گ':true,
	'ە':true,
	'ك':true,
}

var vowels = map[rune]bool{
	'ا':true,
	'و':true,
	'ۇ':true,
	'ە':true,
}
var atoc = map[rune]rune {
	             'ا':'А',
                   'ب':'Б',
                   'ۆ':'В',
                   'گ':'Г',
                   'ع':'Ғ',
                   'د':'Д',
                  'ە':'Е' ,
                  'ج':'Ж' ,
                   'ز':'З',
                   'ي':'И',
                   //'ي':'Й',
                   'ك':'К',
                   'ق':'Қ',
                   'ل':'Л',
                  'م':'М' ,
                   'ن':'Н',
                   'ڭ':'Ң',
                   'و':'О',
                   'پ':'П',
                   'ر':'Р',
                   'س':'С',
                   'ت':'Т',
                   'ۋ':'У',
                   'ۇ':'Ұ',
                   'ف':'Ф',
                   'ح':'Х',
                   'ھ':'Һ',
                   //'س':'Ц',
                   'چ':'Ч',
				   'ش':'Ш',
                   //'ش':'Щ',
				    'ﯼ':'Ы',
				   'ى':'Ы',
                   // 'ا':'Э',
				  '،':',',
				  '؟':'?',
				  '؛':';',
				 '！':'!',
}

func ArabicToCyrillic(s string) string {
	s =strings.ReplaceAll(s,"يا","Я")
	s =strings.ReplaceAll(s,"يۋ","Ю")
	var sentences []string
	withdots:= strings.SplitAfter(s,".")
	temp:=make([]string, 0,len(withdots))
	for _, dotted := range withdots {
		for _, withquestion := range strings.SplitAfter(dotted,"؟") {
			temp=append(temp,withquestion)
		}
	}
	for _, unexclamed := range temp {
		for _, exclamed := range strings.SplitAfter(unexclamed,"！") {
			sentences=append(sentences,exclamed)
		}
	}
	
	var cyrillicSentences strings.Builder
	for _, sentence := range sentences {
		
		var buf []string
		for i, word := range strings.Split(strings.Trim(sentence," ")," ") {
			var builder strings.Builder
			needHarmonify := checkNeedHaromify(&word)
			var previous rune
			for j, char := range word { 
				var cyrillicChar  rune
				
				if char =='ي'{
					if vowels[previous]{
						cyrillicChar='Й'
					}else {
					    cyrillicChar='И'	
					}
				}else{
				cyrillicChar=getCyillicChar(char)
				}
				previous=char
				if i==0&&j==0 {
					builder.WriteRune(cyrillicChar)
				}else{
					builder.WriteRune(unicode.ToLower(cyrillicChar))
				}
			}
			var cyrillicWord string
			if needHarmonify {
				cyrillicWord = strings.Map(func (char rune)rune  {
					if jV,ok:= changetoj[char];ok {
						return jV
					}
					return char
				},builder.String())
			}else{
				cyrillicWord = builder.String()
			}
			
		  buf = append( buf,cyrillicWord)
		}
		ci := strings.Join(buf," ")
		fmt.Println(ci)
		cyrillicSentences.WriteString(strings.Join(buf," "))
	}
	return cyrillicSentences.String()
}



func checkNeedHaromify(word *string) bool {
	if strings.HasPrefix(*word,"ء") {
		*word =strings.TrimPrefix(*word,"ء")
		return true
	}
	for _, char  := range *word {
		if jVowels[char] {
			return true
		}
	}
	return false
}

func getCyillicChar(char rune) rune {
	if cyrill,ok := atoc[char];ok {
		return cyrill
	}
	return char
}



