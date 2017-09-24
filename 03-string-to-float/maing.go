package main 


import (
	"os"
	"fmt"
	"encoding/csv"
	"strconv"
	"sort"
)


func main() {

	// let red the file
	f, err := os.Open("../Environmental_Data_Deep_Moor_2015.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	rdr := csv.NewReader(f)
	rdr.Comma = '\t'
	fmt.Println(rdr.TrimLeadingSpace)
	rdr.TrimLeadingSpace = true
	fmt.Println(rdr.TrimLeadingSpace)

	rows, err := rdr.ReadAll()

	if err != nil {
		panic(err)
	}

	// print the values to the screen
	fmt.Println("Total Records : ", len(rows) - 1)
	fmt.Println("Mean Air Temp\t", "Mean", mean(rows, 1), "Median", median(rows, 1))
	fmt.Println("Mean Barometric\t", "Mean", mean(rows, 2), "Median", median(rows, 2))
	fmt.Println("Mean Win Speed\t", "Mean", mean(rows, 7), "Median", median(rows, 7))


}


// creatint a method for the mean
func mean(rows [][]string, idx int) float64 {
	// local properties
	var total float64

	// loop throw each row
	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			total += val
		}
	}

	// return the results
	return total / float64(len(rows) - 1)
}

// median calculation method
func median(rows [][]string, idx int) float64 {
	// hold data taht will be sorted
	var sorted []float64

	// populate the sorte data looping throw the rows
	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			sorted = append(sorted, val)
		}
	}

	// sort the values
	sort.Float64s(sorted)

	// finding the median
	if len(sorted) % 2 == 0 {
		// even number of items, for example: 3,5,8,9. median are: (5+8) / 2 = 6.5
		middle := len(sorted) / 2
		higher := sorted[middle]
		lower := sorted[middle]
		return (higher + lower) / 2
	}

	// return the median 
	middle := len(sorted) / 2
	return sorted[middle]

}




