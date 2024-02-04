package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	input_dir := "to_process"
	output_dir := "processed"
	var err error
	var files []fs.DirEntry
	files, err = os.ReadDir(input_dir)
	if err != nil {
		log.Fatal(err)
	}

	for f, file := range files {
		if strings.HasSuffix(file.Name(), ".jpg") || strings.HasSuffix(file.Name(), ".jpeg") {

			var input_path string = filepath.Join(input_dir, file.Name())
			var img_file *os.File
			fmt.Println()
			fmt.Println("INPUT ->", input_path)
			fmt.Println("--------------------")
			img_file, err = os.Open(input_path)
			if err != nil {
				log.Fatal(err)
			}

			var new_img_name string = strconv.Itoa(f)
			var img image.Image

			img, err = jpeg.Decode(img_file)
			if err != nil {
				log.Fatal(err)
			}

			var bounds image.Rectangle = img.Bounds()
			var reduction_amount int = 2
			for i := 0; i < 4; i++ {
				bounds_y := bounds.Dy() / reduction_amount
				bounds_x := bounds.Dx() / reduction_amount
				var resized image.Image = create_resized(img, bounds_y, bounds_x)
				size_str := bounds_to_string(resized.Bounds())
				err := save_image(resized, output_dir, new_img_name, size_str)
				if err != nil {
					log.Fatal(err)
				}
				reduction_amount += 2
			}
			fmt.Println("--------------------")
		}
	}
}

func create_resized(img image.Image, x, y int) image.Image {
	resized := resize.Resize(uint(x), uint(y), img, resize.Lanczos3)
	return resized
}

func save_image(img image.Image, output_dir, file_name, size_suffix string) error {
	output_path := filepath.Join(output_dir, fmt.Sprintf("%s_%s.jpg", strings.TrimSuffix(file_name, ".jpg"), size_suffix))
	output_file, err := os.Create(output_path)
	if err != nil {
		return err
	}
	defer output_file.Close()

	err = jpeg.Encode(output_file, img, nil)
	if err != nil {
		return err
	}

	fmt.Println("SAVED ->", output_path)
	return nil
}

func bounds_to_string(bounds image.Rectangle) string {
	return fmt.Sprintf("%dx%d", bounds.Dx(), bounds.Dy())
}
