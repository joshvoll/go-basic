/* example of a file path for the images */
package main   

import (
	"os"
	"fmt"
	"path/filepath"
)

// main walk
func main() {
    // loading the file path
	filepath.Walk("../images", func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		fmt.Println(path)

		return nil
	})
}