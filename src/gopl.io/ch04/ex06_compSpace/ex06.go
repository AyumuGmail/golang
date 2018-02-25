package main

import (
	"unicode"
	"unicode/utf8"
)

func main() {
	data := "世界で　　前に全角スペース 　美しい"
	b := []byte(data)
	compSpace(b)
	//fmt.Printf("%q\n", data)
}
func compSpace(bs []byte) []byte {
	for i := 0; i < len(bs); {
		//fmt.Printf("bslength:%d\n", len(bs))
		r, size := utf8.DecodeRune(bs[i:])
		if unicode.IsSpace(r) && size != 1 {
			bs[i] = ' '
			copy(bs[i+1:], bs[i+size:])
			bs = bs[:len(bs)-size+1] //1は追加したASCII　SPACE分
			//size = 1
			i = i - size + 1 //1は追加したASCII SPACE分
		}
		//fmt.Printf("%d\t%c\n", i, r)
		i = i + size
	}
	//fmt.Printf("%c\t%s\n", bs, bs)
	return bs
}
