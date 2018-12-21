package ux

import "log"

func Err(err error, messages ...interface{}) {
	if err != nil {
		log.Println(messages...)
		log.Fatal(err)
	}
}
