package main  


import (
	"encoding/csv"
	"os"
	"fmt"
)


func main() {
	/* 1.- reading a file from golang */
	f, err := os.Open("../Environmental_Data_Deep_Moor_2015.txt")

	if err != nil {
		panic(err)
	}

	/* 2.- defer the file to save some memory */
	defer f.Close()

	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()

	if err != nil {
		panic(err)
	}

	// loop throgh the columns
	// for _, row := range rows {
	// 	fmt.Println(i, row)
	// }

	// loop and print the first data
	for i, row := range rows {

		fmt.Println(i, row)

		if i == 1 {
			fmt.Println(rows[0])
			break
		}
	}
}