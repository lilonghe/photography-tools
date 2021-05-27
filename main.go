package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	path := flag.String("path", "", "photography path")
	width := flag.Int("width", 300, "resize to width")
	flag.Parse()

	if *path == "" {
		panic("nil path")
	}
	imgPaths := getImagesPath(*path)
	if len(imgPaths) == 0 {
		panic("can not find any jpg image")
	}
	resizeImages(imgPaths, *width)

	fmt.Println("EOF")
}

func getImagesPath(p string) []string {
	paths := make([]string, 0)
	p = strings.TrimRight(p, "/")
	files, err := ioutil.ReadDir(p)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			if strings.HasSuffix(file.Name(), ".thumbnail.jpg") {
				continue
			}

			suffix := strings.ToLower(path.Ext(file.Name()))
			if suffix == ".jpg" || suffix == ".jpeg" {
				paths = append(paths, p+"/"+file.Name())
			}
		}
	}
	return paths
}

func resizeImages(arr []string, width int) {
	for _, path := range arr {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}

		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		m := resize.Resize(uint(width), 0, img, resize.Lanczos3)

		out, err := os.Create(path + ".thumbnail.jpg")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
		fmt.Println("==>", path)
	}

}
