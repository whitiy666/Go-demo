package main

import (
	"fmt"
	"sync"
)

//éð ä¿è¯åªæä¸ä¸ªgor.t.è®¿é®æ°æ®ï¼é¿åæ°æ®ç«äºçé®é¢

var x int64

var wgc sync.WaitGroup
var mu sync.Mutex

func main() {

	wgc.Add(2)
	go add()
	go add()
	wgc.Wait() //ç­å¾å­åç¨å¨é¨å®æ¯
	fmt.Printf("ç»æï¼%dï¼éåºä¸»åç¨", x)

}

func add() {
	for i := 0; i < 5000; i++ {
		mu.Lock() //æ°æ®ä¸é
		x++
		mu.Unlock() //è§£é
	}
	wgc.Done() //åå°åç¨
}
