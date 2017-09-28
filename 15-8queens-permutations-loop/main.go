package main 


import(
	"fmt"
	"strings"
)

func main() {

	// loading variables
	ans := [][]string{}
	options := []string{ "A", "B", "C" }

	// loop all the issues
	for _, x := range options {
		for _, y := range options {
			for _, z := range options {
				// doing the board
				if x != y && x != z && y != z {
					ans = append(ans, strings.Fields(x + " " + y + " " + z))
				}
			}
		}
	}

	// loop throw the ans variabls
	for _, v := range ans {
		fmt.Println(v)
	}
}