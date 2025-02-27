package main

import (
	"fmt"
	"sync"
)

func main() {

	//http.HandleFunc("/login", handlers.Login)
	//http.HandleFunc("/home", handlers.Home)
	//http.HandleFunc("/refresh", Refresh)

	//log.Fatal(http.ListenAndServe(":8080", nil))

	//userInfo := structs.Product{Name: "ajay", Review: 5, IsUS: true}
	//fmt.Println(userInfo)
	//fmt.Printf("%+v\n", userInfo)
	//fmt.Println(userInfo.SingleLine())

	ch := make(chan bool)
	// go SayHello(ch)
	// <-ch

	counter := 0
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	wg.Add(1)
	go SayHello(&wg, ch, &counter, &mx)

	wg.Add(1)
	go SayBye(&wg, ch, &counter, &mx)

	wg.Wait()
	fmt.Println("welcome to the jwt authentication demo server")
	fmt.Println("counter value: ", counter)
}

func SayHello(wg *sync.WaitGroup, ch chan bool, counter *int, mx *sync.Mutex) {
	for i := 1; i <= 5; i++ {
		fmt.Println("saying Hello...")
		mx.Lock()
		*counter++
		mx.Unlock()
	}
	ch <- true
	wg.Done()
}

func SayBye(wg *sync.WaitGroup, ch chan bool, counter *int, mx *sync.Mutex) {
	mx.Lock()
	*counter++
	mx.Unlock()
	<-ch
	fmt.Println("saying Bye...")

	wg.Done()
}

// func (*structs.Product) SingleLine() string {
// 	//userInfo := structs.Product{}
// 	return fmt.Sprintf("Name: %v Review %v IsUS %v", userInfo.Name, userInfo.Review, userInfo.IsUS)
// }
