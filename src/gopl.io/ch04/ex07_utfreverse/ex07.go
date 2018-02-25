package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	data := "山本山の最後"
	b := []byte(data)
	reverseOnlyUTF8(b)
	reverse(b)
	fmt.Printf("%s\n", b)
	reverse(b)
	fmt.Printf("%s\n", b)
	b2 := []byte("山本山の最後の英語編abc")
	b2 = reverse(b2)
	fmt.Printf("%s\n", b2)
}

func reverseOnlyUTF8(bs []byte) {
	//for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
	//rCnt := utf8.RuneCount(bs)
	for i, j := 0, len(bs); i < j; {
		r, size := utf8.DecodeRune(bs[i:])
		r2, size2 := utf8.DecodeLastRune(bs[i:j])

		//このアルゴリズムで何とかする感じなのかしら..

		if size == size2 {
			for k := 0; k < size; k++ {
				bs[i+k], bs[j-size2+k] = bs[j-size2+k], bs[i+k]
			}
		}
		//byteSliceのサイズが同じであれば単純な交換が成立するが、サイズが
		//異なるのでその部分はバッファリングとスライスの一次拡張が必要な気がする。
		//異なるRune長さの際にどうするか？
		fmt.Printf("%d\t%c\n", i, r)
		fmt.Printf("%d\t%c size2:%d\n", i, r2, size2)
		i = i + size
		j = j - size2
	}
	fmt.Printf("%s\n", bs)
}

func reverse(bs []byte) []byte {
	//for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
	runes := []rune(string(bs))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	revBs := []byte(string(runes))
	for i := 0; i < len(bs); i++ {
		bs[i] = revBs[i]
	}
	return revBs
}
