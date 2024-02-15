package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	input_dir := "to_process"
	var err error
	var files []fs.DirEntry
	files, err = os.ReadDir(input_dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err = process(file, input_dir)
		if err != nil {
			fmt.Println("An error has occured: ", err)
		}
	}
}

func process(file fs.DirEntry, input_dir string) error {
	output_dir := "processed"
	var err error
	var input_path string = filepath.Join(input_dir, file.Name())
	var img_file *os.File

	fmt.Println()
	fmt.Println("INPUT ->", input_path)
	fmt.Println("--------------------")

	img_file, err = os.Open(input_path)
	if err != nil {
		return err
	}

	var img image.Image

	if strings.HasSuffix(file.Name(), "png") {
		img, err = png.Decode(img_file)
		if err != nil {
			return err
		}
	}

	if strings.HasSuffix(file.Name(), "jpg") || strings.HasSuffix(file.Name(), "jpeg") {
		img, err = jpeg.Decode(img_file)
		if err != nil {
			return err
		}
	}

	if img == nil {
		return errors.New("Img is nil")
	}

	var bounds image.Rectangle = img.Bounds()
	var reduction_amount int = 2
	for i := 0; i < 4; i++ {
		bounds_y := bounds.Dy() / reduction_amount
		bounds_x := bounds.Dx() / reduction_amount
		var resized image.Image = create_resized(img, bounds_y, bounds_x)
		size_str := bounds_to_string(resized.Bounds())
		err := save_image(resized, output_dir, file.Name(), size_str)
		if err != nil {
			return err
		}
		reduction_amount += 2
	}

	fmt.Println("--------------------")
	return nil
}

func create_resized(img image.Image, y, x int) image.Image {
	resized := resize.Resize(uint(x), uint(y), img, resize.Lanczos3)
	return resized
}

func combine(s string) string {
	return "." + s
}

func save_image(img image.Image, output_dir, file_name, size_suffix string) error {
	var err error
	var output_path string

	split_slice := strings.Split(file_name, ".")

	concat := fmt.Sprintf("%s%s", split_slice[0], split_slice[1])
	if (len(split_slice[0]) == 0 || len(split_slice[1]) == 0) || len(concat) < len(file_name) {
		return errors.New("No existing prefix or file extension")
	}

	output_path = filepath.Join(output_dir, fmt.Sprintf("%s_%s_%s", split_slice[0], size_suffix, combine(split_slice[1])))
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
