package main

import "testing"

func Test_main(t *testing.T) {
	var chars []rune
	for char := 'a'; char <= 'z'; char++ {
		chars = append(chars, char)
	}
	for char := 'A'; char <= 'Z'; char++ {
		chars = append(chars, char)
	}
	for char := '0'; char <= '9'; char++ {
		chars = append(chars, char)
	}
	println(string(chars))
}
