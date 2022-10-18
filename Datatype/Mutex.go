package Datatype

import "sync"

var w sync.Mutex

func Mutex() {
	var count int
	w.Lock()
	count++
	w.Unlock()
}
