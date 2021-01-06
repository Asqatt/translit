# translit
This transliteration program converts Kazakh scrips from Cyrillic Alphabet to Arabic and vice versa.

There are two files called atoc and ctoa,contianing conversion functions from 

Arabic script to Cyrillic script and Cyrillic script to Arabic script , respectively.

They can be used as following:
```
func main(){
    .... code above

	a :="arabic string needed to conversion"
	converted:= translit.ArabicToCyrillic(a)

	//or if you need this in reverse direction:

	c :="cyrillic string needed to conversion"
	converted= translit.CyrillicToArabic(c)

    ..... code below

}

```


