package main  


import(
	"fmt"
	"time"
)


func main() {
	c := make(chan int)

	go func() {
		for i:= ; i < 45; i++ {
			c <- i
		}
	}()

	go func() {
		for {
			fmt.Println(<- c)
		}
	}()

	// time the calling 
	time.Sleep(time.Second)
}