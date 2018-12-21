package ux

import "sync"

type AsyncSliceStringMap struct {
	data sync.Map
}

func (as AsyncSliceStringMap) Get(key string) []string {
	v, found := as.data.Load(key)
	if found {
		return v.([]string)
	} else {
		return []string{}
	}
}
func (as AsyncSliceStringMap) Set(key string, v []string) {
	as.data.Store(key, v)
}
