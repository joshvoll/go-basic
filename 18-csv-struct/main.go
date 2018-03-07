package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// StateInformation for state information
type StateInformation struct {
	columns map[string]int
}

// State defining the state that i need from the code
type State struct {
	ID               int
	Name             string
	Abbreviation     string
	CensusRegionName string
}

func (info *StateInformation) setColumns(record []string) {
	for idx, column := range record {
		info.columns[column] = idx
	}
}

func (info *StateInformation) parseState(record []string) (*State, error) {
	// adding the columnt to the struct
	column := info.columns["id"]
	id, err := strconv.Atoi(record[column])
	if err != nil {
		return nil, err
	}
	name := record[info.columns["name"]]
	abbreviation := record[info.columns["abbreviation"]]
	censusRegionName := record[info.columns["census_region_name"]]

	// return everything to the mehtod
	return &State{
		ID:               id,
		Name:             name,
		Abbreviation:     abbreviation,
		CensusRegionName: censusRegionName,
	}, nil
}

func main() {
	// open the file
	file, err := os.Open("state_table.csv")
	if err != nil {
		log.Fatal(err)
	}

	// always defer the memory
	defer file.Close()

	// declaring a state loop up to search by abbreviation
	info := &StateInformation{columns: make(map[string]int, 0)}
	stateLookup := map[string]*State{}

	// read the file
	csvReader := csv.NewReader(file)
	// loop throw the file to read
	for rowCount := 0; ; rowCount++ {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if rowCount == 0 {
			// loop throw the index
			info.setColumns(record)

		} else {
			// getting the data from the method
			state, err := info.parseState(record)
			if err != nil {
				log.Fatal(err)
				break
			}

			stateLookup[state.Abbreviation] = state
		}

	}

	// get the input from the user
	if len(os.Args) < 2 {
		log.Fatal("expected state code like AL or FL")
	}

	abbreviation := os.Args[1]

	state, ok := stateLookup[abbreviation]
	if !ok {
		log.Fatal("problem getting the state")
	}

	fmt.Println(`
		<html>
			<head></head>
			<body>
				<table>
					<tr>
						<th>Abbreviation</th>
						<th>Name</th>
					</tr>`)
	fmt.Println(`
		<tr>
			<td>` + state.Abbreviation + `</td>
			<td>` + state.Name + `</td>
		</tr> 	
	`)
	fmt.Println(`
				</table>
			</body>
		</html>
	`)
}
