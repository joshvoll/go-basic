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
	fmt.Println("Mean Air Temp\t", mean(rows, 1), median(rows, 1))
	fmt.Println("Mean Barometric\t", mean(rows, 2), median(rows, 2))
	fmt.Println("Mean Win Speed\t", mean(rows, 7), median(rows, 7))


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

	// before sorted the float64; the first 10 records
	fmt.Println("BEFORE SORTING")

	// loop to check the sorting using sorted variable
	for i, v := range sorted {
		fmt.Println(v)

		if i == 9 {
			break
		}
	}

	// using teh go method to sort 
	sort.Float64s(sorted)

	// after sorting
	fmt.Println("AFTER SORTNG")

	for i, v := range sorted {
		fmt.Println(v)

		if i == 9 {
			break
		}
	}

	return 0.0

}







