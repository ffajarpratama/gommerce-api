package util

import (
	"strings"

	"math/rand"
)

var (
	BASE    = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	LETTERS = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	NUMBERS = []rune("0123456789")
)

func RandomNumber(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = NUMBERS[rand.Intn(len(NUMBERS))]
	}

	return string(b)
}

func RandomString(length int, upper bool, alphaNumeric bool) string {
	b := make([]rune, length)

	if alphaNumeric {
		for i := range b {
			b[i] = BASE[rand.Intn(len(BASE))]
		}
	} else {
		for i := range b {
			b[i] = LETTERS[rand.Intn(len(LETTERS))]
		}
	}

	if upper {
		return strings.ToUpper(string(b))
	}

	return string(b)
}
