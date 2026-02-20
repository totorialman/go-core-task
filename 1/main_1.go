package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strings"
)

func GetType(v any) string {
	return reflect.TypeOf(v).String()
}

func ToString(v ...any) string {
	var res strings.Builder
	for _, val := range v {
		fmt.Fprint(&res, val)
	}
	return res.String()

}

func ToRunes(s string) []rune {
	return []rune(s)
}

func Hash(r []rune, salt string) string {
	res := make([]rune, 0, len(r)+len(salt))
	mid := len(r) / 2

	res = append(res, r[:mid]...)
	res = append(res, []rune(salt)...)
	res = append(res, r[mid:]...)

	hash := sha256.Sum256([]byte(string(res)))
	return hex.EncodeToString(hash[:])
}

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	fmt.Println("numDecimal: ", GetType(numDecimal))
	fmt.Println("numOctal: ", GetType(numOctal))
	fmt.Println("numHexadecimal: ", GetType(numHexadecimal))
	fmt.Println("pi: ", GetType(pi))
	fmt.Println("name: ", GetType(name))
	fmt.Println("isActive: ", GetType(isActive))
	fmt.Println("complexNum: ", GetType(complexNum))

	joined := ToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	runes := ToRunes(joined)

	hash := Hash(runes, "go-2024")

	fmt.Println(hash)
}
