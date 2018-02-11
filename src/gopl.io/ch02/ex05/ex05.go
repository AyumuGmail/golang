package popcount

//pc[i]はiのポピュレーションカウント
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		//DEBUG fmt.Printf("init pc[%d]=%d\n", i, pc[i])
	}
}

//PopCountはxのポピュレーションカウント
func PopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) int {
	var mask uint64 = 1
	var num int
	//var i int

	for ; mask != 0; mask = mask << 1 {
		if (x & mask) != 0 {
			num++
		}
		//i++
	}
	//fmt.Printf("loop cnt:%d\n", i)
	return num
}
func PopCountXminusOneLoop(x uint64) int {
	var num int
	for ; x != 0; x &= x - 1 {
		num++
	}
	return num
}
