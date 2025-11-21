package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"slices"
)

func main() {
	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64
	args := []any{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}

	for _, val := range args {
		fmt.Println(TypeOf(val))
	}
	strArgs := ContainToString(args...)
	fmt.Println(strArgs)
	runeArgs := ToSliceRune(strArgs)
	fmt.Println(runeArgs)
	fmt.Println(Sha256EndcodingWithSalt([]rune("aaa"), []rune("go-2024")))
}

func TypeOf(a any) reflect.Type {
	return reflect.TypeOf(a)
}

func ContainToString(a ...any) string {
	return fmt.Sprint(a...)
}

func ToSliceRune(s string) []rune {
	return []rune(s)
}

func Sha256EndcodingWithSalt(s, salt []rune) [32]byte {
	return sha256.Sum256([]byte(string(addSalt(s, salt))))
}

func addSalt(sl []rune, salt []rune) []rune {
	n := len(sl)
	return slices.Concat(sl[:n], salt, sl[n:])
}
