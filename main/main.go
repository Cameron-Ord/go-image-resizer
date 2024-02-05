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
		if strings.HasSuffix(file.Name(), ".jpg") || strings.HasSuffix(file.Name(), ".jpeg") {
			err = do_jpg(file, input_dir)
			if err != nil {
				fmt.Printf("Image processing failed: %s\n", err)
			}
		}

		if strings.HasSuffix(file.Name(), ".png") {
			err = do_png(file, input_dir)
			if err != nil {
				fmt.Printf("Image processing failed: %s\n", err)
			}
		}
	}
}

func do_png(file fs.DirEntry, input_dir string) error {
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

	img, err = png.Decode(img_file)
	if err != nil {
		return err
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

func do_jpg(file fs.DirEntry, input_dir string) error {
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

	img, err = jpeg.Decode(img_file)
	if err != nil {
		return err
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

func create_resized(img image.Image, x, y int) image.Image {
	resized := resize.Resize(uint(x), uint(y), img, resize.Lanczos3)
	return resized
}

func extract_file_extension(filename string) (string, string) {
	var prefix_substr string = ""
	var extension_string string = ""
	for i := 0; i < len(filename); i++ {
		char := filename[i]
		if char != '.' {
			prefix_substr += string(char)
		}

		if char == '.' {
			for ci := i; ci < len(filename); ci++ {
				if ci >= 0 && ci < len(filename) {
					extension_string += string(filename[ci])
				}
			}
			return prefix_substr, extension_string
		}
	}
	return "", ""
}

func save_image(img image.Image, output_dir, file_name, size_suffix string) error {
	var err error
	var output_path string
	prefix_substr, ext_substr := extract_file_extension(file_name)
	if len(prefix_substr) == 0 || len(ext_substr) == 0 {
		return errors.New("No existing prefix or file extension")
	}

	output_path = filepath.Join(output_dir, fmt.Sprintf("%s_%s_%s", prefix_substr, size_suffix, ext_substr))
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
