package exercise

import (
	"fmt"
	"sync"
)

var mapShared = make(map[string]int)
var wg sync.WaitGroup
var mutex sync.Mutex

func readMap() {
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		_ = mapShared["SHARED"]
		mutex.Unlock()
	}
	wg.Done()
}

func writeMap() {
	for i := 0; i < 100000; i++ {
		mutex.Lock()
		mapShared["SHARED"] += 1
		mutex.Unlock()
	}
	wg.Done()
	fmt.Println(mapShared["SHARED"])
}

func Exercise_6() {
	wg.Add(2)
	go readMap()
	go writeMap()
	wg.Wait()
	fmt.Println("-----")
}
