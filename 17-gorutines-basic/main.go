/* Basic go rutines */
package main 

import "fmt"

func main() {
	go foo()
	go boo()
}

func foo() {

	for i := 0; i < 45; i++ {
		fmt.Println("Foo: ",i)
	}
}

func boo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Boo: ", i)
	}
}