package common

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var space = regexp.MustCompile(`\s+`)

func ToSnakeCase(str string) string {
	snake := space.ReplaceAllString(str, " ")
	snake = strings.ReplaceAll(snake, " ", "_")
	return strings.ToLower(snake)
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
