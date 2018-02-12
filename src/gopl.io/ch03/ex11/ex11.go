package comma

import (
	"bytes"
	"math"
	"strconv"
)

func commaFloat(f float64) string {
	s := ""
	absI := int(math.Abs(f))
	if f < 0 {
		s = "-"
	}
	s = s + comma(strconv.Itoa(absI)) + deminalPoint(strconv.FormatFloat(f, 'f', 10, 64))
	return s
}

func commaBuf(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for i := 0; i < n; i++ {
		if i%3 == 1 && n > 3 {
			buf.WriteString(",")
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func deminalPoint(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			s = s[i:]
			break
		}
	}
	//".12300020"
	//".0000"
	//fmt.Printf("start:%s\n", s)
	for i := len(s) - 1; i >= 0; i-- {
		//後ろのゼロは切り捨て
		if s[i] != '0' {
			//fmt.Printf("math:%d\n", i)
			//iも含む文字を保持する
			s = s[:i+1]
			break
		}
	}
	//fmt.Printf("切り捨て:%s\n", s)

	if len(s) == 1 && s[0] == '.' {
		s = ""
	}
	return s
}
