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

	"github.com/disintegration/imaging"
)

func main() {
	path := flag.String("path", "", "photography path")
	width := flag.Int("width", 1000, "resize to width")
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
		img, err := imaging.Open(path, imaging.AutoOrientation(true))
		if err != nil {
			log.Fatal(err)
		}

		m := imaging.Resize(img, width, 0, imaging.Lanczos)

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
