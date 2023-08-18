//Задача 1
//Получение максимального значения для типа

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Если значение не нулевое, вернуть максимальное значение для типа
func getIntMaxValue(in8 int8, in16 int16, in32 int32, in64 int64) (int8, int16, int32, int64) {

	if in8 == 0 {
		in8 = 1<<(8-1) - 1
	}
	if in16 == 0 {
		in16 = 1<<(16-1) - 1
	}
	if in32 == 0 {
		in32 = 1<<(32-1) - 1
	}
	if in64 == 0 {
		in64 = 1<<(64-1) - 1
	}
	return in8, in16, in32, in64
}

// Если значение не нулевое, вернуть максимальное значение для типа
func getUintMaxValue(uin8 uint8, uin16 uint16, uin32 uint32, uin64 uint64) (uint8, uint16, uint32, uint64) {
	if uin8 == 0 {
		uin8 = ^uint8(0)
	}
	if uin16 == 0 {
		uin16 = ^uint16(0)
	}
	if uin32 == 0 {
		uin32 = ^uint32(0)
	}
	if uin64 == 0 {
		uin64 = ^uint64(0)
	}
	return uin8, uin16, uin32, uin64
}

func getBits(v interface{}) int {
	rawType := fmt.Sprintf("%T", v)
	typeBits := strings.Split(rawType, "t")[1]
	bits, _ := strconv.Atoi(typeBits)

	return bits
}

func comparison(x, y int) bool {
	if x^y == 0 {
		return true
	}
	return false
}

func increment(n *int) {
	*n++
}

func main() {
	x := 4
	y := 4
	fmt.Println(comparison(x, y))
	increment(&x)
	fmt.Println(x)


}
