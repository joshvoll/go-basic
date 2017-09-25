// getting the data from the us army
package main  


import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)


func main() {

	// get the data from a location
	res, err := http.Get("http://lpo.dt.navy.mil/data/DM/Environmental_Data_Deep_Moor_2015.txt")

	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", bs)
}