package main

import (
	"fmt"
	"image"
	"io/fs"
	"strconv"

	"image/jpeg"

	"path/filepath"
	"strings"

	"log"
	"os"

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
			var img_sizes Img_Bounds = create_img_struct(bounds)
			all_resized := create_resized(img, img_sizes)

			for r := 0; r < len(all_resized); r++ {
				resized := all_resized[r]
				size_str := bounds_to_string(resized.Bounds())
				err := save_image(resized, output_dir, new_img_name, size_str)
				if err != nil {
					log.Fatal(err)
				}
			}
			fmt.Println("--------------------")
		}
	}
}

func create_resized(img image.Image, img_sizes Img_Bounds) []image.Image {
	tiny := resize.Resize(uint(img_sizes.bounds_x_tiny), uint(img_sizes.bounds_y_tiny), img, resize.Lanczos3)
	smallest := resize.Resize(uint(img_sizes.bounds_x_smallest), uint(img_sizes.bounds_y_smallest), img, resize.Lanczos3)
	small := resize.Resize(uint(img_sizes.bounds_x_small), uint(img_sizes.bounds_y_small), img, resize.Lanczos3)
	medium := resize.Resize(uint(img_sizes.bounds_x_med), uint(img_sizes.bounds_y_med), img, resize.Lanczos3)
	big := resize.Resize(uint(img_sizes.bounds_x_big), uint(img_sizes.bounds_y_big), img, resize.Lanczos3)
	maximum := resize.Resize(uint(img_sizes.bounds_x_max), uint(img_sizes.bounds_y_max), img, resize.Lanczos3)

	all_resized := []image.Image{tiny, smallest, small, medium, big, maximum}
	return all_resized
}

func create_img_struct(bounds image.Rectangle) Img_Bounds {
	img_sizes := Img_Bounds{
		bounds_y_tiny:     bounds.Dy() / 6,
		bounds_x_tiny:     bounds.Dx() / 6,
		bounds_y_smallest: bounds.Dy() / 5,
		bounds_x_smallest: bounds.Dx() / 5,
		bounds_y_small:    bounds.Dy() / 4,
		bounds_x_small:    bounds.Dx() / 4,
		bounds_y_med:      bounds.Dy() / 3,
		bounds_x_med:      bounds.Dx() / 3,
		bounds_y_big:      bounds.Dy() / 2,
		bounds_x_big:      bounds.Dx() / 2,
		bounds_y_max:      bounds.Dy(),
		bounds_x_max:      bounds.Dx(),
	}

	return img_sizes
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
