package ux

import (
	"math/rand"
	"time"
)

//
//func NewSessionToken() string {
//	hasher := sha1.New()
//	rand.ExpFloat64()
//	hasher.Write([]byte(strconv.FormatInt(rand.Int63(), 13)))
//	hasher.Write([]byte(strconv.FormatInt(rand.Int63(), 13)))
//	hasher.Write([]byte(strconv.FormatInt(rand.Int63(), 13)))
//	return "n." + hex.EncodeToString(hasher.Sum(nil))
//}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1092384756")

func MakeRune(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func RemoveString(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
