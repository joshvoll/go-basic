package main  

import (
	"fmt"
	"os"
	"log"
	"image/jpeg"
)


func main() {
	// first lest open a image file
	f, err := os.Open("../images/78771293.jpg")

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

	img, err := jpeg.Decode(f)

	if err != nil {
		log.Fatal(err)
	}

	// print the type of the images
	fmt.Printf("IMAGE TYPE: %T \n", img)

	// image bounds
	bounds := img.Bounds()

	fmt.Println("Widht x Height", bounds.Dx(), "x", bounds.Dy())
	fmt.Println("Total Pixel", bounds.Dx() * bounds.Dy())

	r,g,b,a := img.At(0, 0).RGBA()

	fmt.Println("First pixel", r,g,b,a)

	r,g,b,a = img.At(bounds.Dx(), bounds.Dy()).RGBA()

	fmt.Println("Last Pixel", r,g,b,a)

}



