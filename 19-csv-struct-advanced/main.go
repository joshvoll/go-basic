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

// Shares of the records struct
type Shares struct {
	Open  float64
	High  float64
	Low   float64
	Close float64
}

// InfoShares is a type of the data
type InfoShares struct {
	columns map[string]int
}

func (info *InfoShares) setColumns(record []string) {
	for idx, column := range record {
		info.columns[column] = idx
	}
}

func (info *InfoShares) parseShare(record []string) (*Shares, error) {

	open, _ := strconv.ParseFloat(record[1], 64)
	high, _ := strconv.ParseFloat(record[info.columns["High"]], 64)
	low, _ := strconv.ParseFloat(record[info.columns["Low"]], 64)
	close, _ := strconv.ParseFloat(record[info.columns["Close"]], 64)

	// return the values
	return &Shares{
		Open:  open,
		High:  high,
		Low:   low,
		Close: close,
	}, nil

}

func main() {
	// open file
	file, err := os.Open("appl.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	info := &InfoShares{columns: make(map[string]int, 0)}

	fmt.Println(`
		<html>
			<head>
			<script src="http://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha256-3edrmyuQ0w65f8gfBsqowzjJe2iM6n0nKciPUp8y+7E=" crossorigin="anonymous"></script>
			<script src="https://code.highcharts.com/highcharts.js"></script>
			<script src="https://code.highcharts.com/modules/series-label.js"></script>
			<script src="https://code.highcharts.com/modules/exporting.js"></script>
			</head>
			<body>
				<div id="container"></div>
				<table>
					<tr>
						<th>Open</th>
						<th>High</th>
						<th>Low</th>
						<th>Close</th>
					</tr>	`)

	openValues := []string{}
	highValues := []string{}
	lowValues := []string{}
	closeValues := []string{}

	for rowCount := 0; ; rowCount++ {
		// definging locals
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if rowCount == 0 {

			info.setColumns(record)

		} else {

			share, err := info.parseShare(record)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(`
				<tr>
					<td>` + fmt.Sprintf("%.2f", share.Open) + `</td>
					<td>` + fmt.Sprintf("%.2f", share.High) + `</td>
				    <td>` + fmt.Sprintf("%.2f", share.Low) + `</td>
					<td>` + fmt.Sprintf("%.2f", share.Close) + `</td>
				</tr>
			`)

			openValues = append(openValues, fmt.Sprintf("%.2f", share.Open))
			highValues = append(highValues, fmt.Sprintf("%.2f", share.High))
			lowValues = append(lowValues, fmt.Sprintf("%.2f", share.Low))
			closeValues = append(closeValues, fmt.Sprintf("%.2f", share.Close))

		}

	}

	fmt.Println(`
			</table>
		</body>
		<script>
			Highcharts.chart('container', {

				title: {
					text: 'apple inc share values, 2017-2018'
				},
			
				subtitle: {
					text: 'Source: https://finance.yahoo.com/quote/AAPL/history?ltr=1'
				},
			
				yAxis: {
					title: {
						text: 'time'
					}
				},
				legend: {
					layout: 'vertical',
					align: 'right',
					verticalAlign: 'middle'
				},
			
				plotOptions: {
					series: {
						label: {
							connectorAllowed: false
						},
						pointStart: 2010
					}
				},
			
				series: [{
					name: 'Open',
					data: [` + strings.Join(openValues, ",") + `]
				}, {
					name: 'High',
					data: [` + strings.Join(highValues, ",") + `]
				}, {
					name: 'Low',
					data: [` + strings.Join(lowValues, ",") + `]
				}, {
					name: 'Close',
					data: [` + strings.Join(closeValues, ",") + `]
				}],
			
				responsive: {
					rules: [{
						condition: {
							maxWidth: 500
						},
						chartOptions: {
							legend: {
								layout: 'horizontal',
								align: 'center',
								verticalAlign: 'bottom'
							}
						}
					}]
				}
			
			});
		</script>
	</html>	

	`)

}
