package loki

import (
	"fmt"
	"github.com/afiskon/promtail-client/promtail"
	"log"
	"math/rand"
	"os"
	"time"
)

var loki promtail.Client

func init() {
	rand.Seed(time.Now().UnixNano())
	conf := promtail.ClientConfig{
		PushURL: "http://88.214.236.159:13100/api/prom/push",
		//PushURL:            "https://loki.unq.onl/api/prom/push",
		Labels:             `{app="test1", mode="dev"}`,
		BatchWait:          5 * time.Second,
		BatchEntriesNumber: 10000,
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.ERROR,
	}
	var err error
	loki, err = promtail.NewClientProto(conf)
	if err != nil {
		log.Printf("promtail.NewClient: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("loki logging fine")
}
func Infof(s string, args ...interface{}) {
	loki.Infof(s, args...)
}
