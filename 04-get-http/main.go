/* example how to use http class for go */
package main 


import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


func main() {
	// getting the data of the query
	res, err := http.Get("http://www.sandals.com/")

	if err != nil {
		log.Fatal(err)
	}

	// read the pages
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", page)
}