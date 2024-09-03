package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // use pointer instead

var mut sync.Mutex // use pointer instead

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.golang.org",
		"https://www.stackoverflow.com",
	}

	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string){
// 	for i := 0; i < 6; i++ {
// 		time.Sleep((3 * time.Millisecond))
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d Status Code for %s \n", res.StatusCode, endpoint)
	}
}
