package main 

import (
	"fmt"
	"os"
	"log"
	"image"
	"image/jpeg"
)

func main() {
	// call to load the iages
	img := loadImage("../images/78771293.jpg")

	// loading all the data for images
	r,g,b,a := img.At(0,0).RGBA()

	// print the informaton
	fmt.Printf("%d %d %d %d \n", r,g,b,a)

}

func loadImage(filename string) image.Image {
	
	// get the files and return and images
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	img, err := jpeg.Decode(f)

	if err != nil {
		log.Fatal(err)
	}

	return img
}