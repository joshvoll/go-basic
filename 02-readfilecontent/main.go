/* changin the delimiler to read the file on better way */

package main  


import (
	"fmt"
	"os"
	"encoding/csv"
)


func main() {

	// read the file
	f, err := os.Open("../Environmental_Data_Deep_Moor_2015.txt")

	if err != nil {
		panic(err)
	}
	// defer the files
	defer f.Close()

	// read the files
	rdr := csv.NewReader(f)
	
	// create the delimieter here
	rdr.Comma = '\t'
	fmt.Println(rdr.TrimLeadingSpace)
	rdr.TrimLeadingSpace = true
	fmt.Println(rdr.TrimLeading)

	rows, err := rdr.ReadAll()

	if err != nil {
		panic(err)
	} 

	// loop throu the file
	for i, row := range rows {
		fmt.Println(row)

		// get the second row
		if i == 1 {
			fmt.Printf("%T %T %T\n", row[1], row[2], row[7])
			fmt.Println(row[1], row[2], row[7])
			break
		}
	}

}