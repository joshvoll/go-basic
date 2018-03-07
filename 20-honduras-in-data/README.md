package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// CountryInformation for the columns of the csv
type CountryInformation struct {
	columns map[string]int
}

// Country the struct of the code
type Country struct {
	CountryName string
	CountryCode string
	Year        int
}

func (info *CountryInformation) setColumns(record []string) {
	for idx, column := range record {
		info.columns[column] = idx

	}
}

func (info *CountryInformation) parseCountry(record []string) (*Country, error) {
	// defining the columns
	column := info.columns["Year"]
	countryName := record[info.columns["Entity"]]
	countryCode := record[info.columns["Code"]]
	year, err := strconv.Atoi(record[column])
	if err != nil {
		return nil, err
	}

	return &Country{
		CountryName: countryName,
		CountryCode: countryCode,
		Year:        year,
	}, nil

}

func main() {
	// 1.- open the file
	file, err := os.Open("cancer-death-by-type/cancer-death-by-type.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// local properties
	csvReader := csv.NewReader(file)
	info := &CountryInformation{columns: make(map[string]int, 0)}

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

			country, err := info.parseCountry(record)
			if err != nil {
				log.Fatal(err)
			}

			log.Println(country)
		}

	}

}
