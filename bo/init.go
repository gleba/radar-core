package bo

import (
	"github.com/gleba/radar-core/ns/bridge"
)

func ReIndexerNS(namespaces ...bridge.NS) {
	bridge.ToReIndexer(namespaces...)
}
