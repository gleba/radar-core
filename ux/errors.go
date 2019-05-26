package ux

import "log"

func Safe(err error, messages ...interface{}) bool {
	if err != nil {
		log.Println(messages...)
		log.Fatal(err)
		return false
	}
	return true
}
