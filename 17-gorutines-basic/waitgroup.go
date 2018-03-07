package main


import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go boo()
	wg.Wait()

}

func foo() {

	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ",i)
	}

	wg.Done()
}

func boo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Boo: ", i)
	}

	wg.Done()
}