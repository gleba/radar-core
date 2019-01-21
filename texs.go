package main

import (
	"fmt"
	"github.com/afiskon/promtail-client/promtail"
	"log"
	"math/rand"
	"os"
	"time"
)

func displayUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s proto|json\n", os.Args[0])
	os.Exit(1)
}
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

func main() {
	if len(os.Args) < 2 {
		displayUsage()
	}

	format := os.Args[1]
	if format != "proto" && format != "json" {
		displayUsage()
	}

	conf := promtail.ClientConfig{
		PushURL: "http://88.214.236.159:13100/api/prom/push",
		//PushURL:            "https://loki.unq.onl/api/prom/push",
		Labels:             `{app="test1", mode="dev"}`,
		BatchWait:          5 * time.Second,
		BatchEntriesNumber: 10000,
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.ERROR,
	}

	var (
		loki promtail.Client
		err  error
	)

	if format == "proto" {
		loki, err = promtail.NewClientProto(conf)
	} else {
		loki, err = promtail.NewClientJson(conf)
	}

	if err != nil {
		log.Printf("promtail.NewClient: %s\n", err)
		os.Exit(1)
	}

	for i := 1; i < 50; i++ {
		loki.Warnf("some random rune %s, i = %d\n", MakeRune(10), i)
		loki.Infof("%d %s some random rune", i, MakeRune(i))
		//loki.Warnf("The time is %s, i = %d\n", time.Now().String(), i)
		//loki.Errorf("The time is %s, i = %d\n", time.Now().String(), i)
		time.Sleep(1 * time.Millisecond)
		fmt.Println("x", i)
	}

	loki.Shutdown()
	fmt.Println("done")
}
