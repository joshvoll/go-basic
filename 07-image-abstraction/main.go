
package main 

import (
	"fmt"
	"os"
	"log"
	"image"
	"image/jpeg"
)

type loadimage interface {
	LoadImage() image.Image
}

type imagen struct {
	filename string
}

func (fil imagen) LoadImage() image.Image {

	fmt.Println("loading interfaces")
	f, err := os.Open(fil.filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// get the information of the file
	fi, _ := f.Stat()

	fmt.Println("File name: \t\t", fi.Name())
	fmt.Println("File Size: \t\t", fi.Size())
	fmt.Println("File Mode: \t\t", fi.Mode())
	fmt.Println("File ModTime: \t\t", fi.ModTime())
	fmt.Println("File Is Directory: \t\t", fi.IsDir())
	fmt.Println("File Sys: \t\t", fi.Sys())

	imgn, err := jpeg.Decode(f)

	if err != nil {
		log.Fatal(err)
	}

	return imgn
}

func main() {

	// call to load the iages
	img := imagen{ filename: "../images/78771293.jpg" }

	 // loading all the data for images
    r,g,b,a := img.LoadImage().At(0,0).RGBA()

    // print the informaton
    fmt.Printf("%d %d %d %d \n", r,g,b,a)

}


