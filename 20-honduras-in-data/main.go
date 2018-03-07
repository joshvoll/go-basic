// main package
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// CountryInformation type of all the columns
type CountryInformation struct {
	columns map[string]int
}

// Country the columns we're going to use
type Country struct {
	CountryCode string
	CountryName string
	Year        int
}

// set the columns from the csv file, this is the first record from the loop
func (info *CountryInformation) setColumns(record []string) {
	// set the column with a loop
	for idx, column := range record {
		info.columns[column] = idx
	}

}

func (info *CountryInformation) parseCountry(record []string) (*Country, error) {
	// defining the coulumn will be add it o the struct
	countryCode := record[info.columns["Code"]]
	countryName := record[info.columns["Entity"]]

	// converting the string to inter
	column := info.columns["Year"]
	year, err := strconv.Atoi(record[column])
	if err != nil {
		return nil, err
	}

	return &Country{
		CountryCode: countryCode,
		CountryName: countryName,
		Year:        year,
	}, nil
}

func main() {
	// read the file
	file, err := os.Open("cancer-death-by-type/cancer-death-by-type.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	info := &CountryInformation{columns: make(map[string]int, 0)}
	countryLookup := map[string]*Country{}

	for rowCount := 0; ; rowCount++ {

		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if rowCount == 0 {

			info.setColumns(record)

		} else {
			// parsing the country
			country, err := info.parseCountry(record)
			if err != nil {
				log.Fatal(err)
			}

			countryLookup[country.CountryCode] = country

		}

	}

	if len(os.Args) < 2 {
		fmt.Println("you need to pass the country name example honduras or usa")
	}

	countryNameArgs := strings.ToUpper(os.Args[1])

	countr, ok := countryLookup[countryNameArgs]
	if !ok {
		log.Fatal("Something when wrong, we could not found your country name")
	}

	fmt.Println(countr)
}
