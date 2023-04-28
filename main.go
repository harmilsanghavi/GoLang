package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var signals = []string{"test"}
var mut sync.Mutex

func main() {
	web := []string{
		"http://google.com",
		"http://github.com",
		"http://w3school.com",
	}
	for _, we := range web {
		go call(we)
		// go sum()
		// go sub()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Printf("%s\n", signals)
}
func call(web string) {
	defer wg.Done()
	res, err := http.Get(web)
	if err != nil {
		fmt.Println("oops")
	} else {
		mut.Lock()
		signals = append(signals, web)
		mut.Unlock()
		fmt.Printf("%d connected %s\n", res.StatusCode, web)
	}
}

// func sum() {
// 	defer wg.Done()
// 	fmt.Printf("%d sum\n", 5+3)
// }
// func sub() {
// 	defer wg.Done()
// 	fmt.Printf("%d sub\n", 5-3)
// }
