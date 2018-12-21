package bo

import (
	"github.com/gleba/radar-core/ns/bridge"
)

func Init(namespaces ...bridge.NS) {
	bridge.ToReIndexer(namespaces...)
}
