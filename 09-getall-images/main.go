/* 
	example to get all the images and process the files

*/

package main 


import (
	"fmt"
	"os"
	"image"
	"image/jpeg"
	"path/filepath"
	"log"
)

// defining the interfaces
type loadimage interface {
	LoadImage() image.Image
}

// defining the struct
type imagen struct {
	filename string
}

// definding the procedure call
func (fil imagen) LoadImage() image.Image {
    // first let load the files
    f, err := os.Open(fil.filename)

    if err != nil {
    	log.Fatal(err)
    }

    defer f.Close()

    // get the images information
    imgn, err := jpeg.Decode(f)

    if err != nil {
    	log.Fatal(err)
    }

    return imgn

}

func getImages() []image.Image {

	// get all the images
	var images []image.Image

	// walk to the folder to look the images
	filepath.Walk("../images", func(path string, info os.FileInfo, err error) error {

	    if info.IsDir() {
	    	return nil
	    }

	    img := imagen{ filename: path}
	    images = append(images, img.LoadImage())

	    return nil
	})

	return images

}

func main() {
	// load all teh images
	images := getImages()

	for _, v := range images {
		r, g, b, a := v.At(0, 0).RGBA()
		fmt.Printf("%d %d %d %d \n", r, g, b, a) 
	}
}




