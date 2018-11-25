package ns

import (
	"github.com/nats-io/go-nats"
	"github.com/vmihailenco/msgpack"
	"log"
	"os"
)

var nc *nats.Conn

func ifError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func Connect() {
	var err error
	nc, err = nats.Connect("nats://" + os.Getenv("NATS_SERVER"))
	ifError(err)
	log.Println("start core NATS")
}

func Subscribe(subj string, cb nats.MsgHandler) {
	_, err := nc.Subscribe(subj, cb)
	ifError(err)
}
func Publish(channel string, data interface{}) {
	bytes, _ := msgpack.Marshal(&data)
	err := nc.Publish(channel, bytes)
	ifError(err)
}
