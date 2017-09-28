/* creating the board with recursive fucntions */

package main 


import (
	"fmt"
)


func main() {
	options := []string{ "A", "B", "C" }

	perm := permutations(options)

	for _, v := range perm {
		fmt.Println(v)
	}

	fmt.Println("Numbers of answers:", len(perm))
}

// this method will recibe a string object and return a 2 dimensional object
func permutations(options []string) [][]string {
	// sending the permute method the string object and the index = 0
	return permute(options, 0)
}

func permute(opt []string, start int) [][]string {
	// calculating the permutations doing a loop
	answer := [][]string{}
	end := len(opt) - 1

	if start == end {
		ans := make([]string, len(opt))
		copy(ans, opt)
		answer = append(answer, ans)

	} else {

		for i := start; i <= end; i++ {
			// we're going to swap the board
			swap(opt, start, i)

			// we're going to use recursive to make it faster 
			subAnswer := permute(opt, start + 1)

			// looping throw the result after permute again
			for _, ans := range subAnswer {
				answer = append(answer, ans)
			}

			swap(opt, start, i)
		}
	}

	return answer
} 

// swaping the numbers
func swap(a []string, x, y int) {
	a[x], a[y] = a[y], a[x]
}

