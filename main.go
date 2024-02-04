package main

import (
	"fmt"
	"image"
	"strconv"

	"image/jpeg"

	"path/filepath"
	"strings"

	"log"
	"os"

	"github.com/nfnt/resize"
)

type img_bounds struct {
	bounds_y_small int
	bounds_x_small int
	bounds_y_med   int
	bounds_x_med   int
	bounds_y_big   int
	bounds_x_big   int
	bounds_y_max   int
	bounds_x_max   int
}

type resized_images struct {
	small   image.Image
	medium  image.Image
	big     image.Image
	maximum image.Image
}

func main() {
	input_dir := "to_process"
	output_dir := "processed"
	files, err := os.ReadDir(input_dir)
	if err != nil {
		log.Fatal(err)
	}

	for f, file := range files {
		if strings.HasSuffix(file.Name(), ".jpg") {
			input_path := filepath.Join(input_dir, file.Name())
			img_file, err := os.Open(input_path)
			if err != nil {
				log.Fatal(err)
			}

			new_img_name := strconv.Itoa(f)

			img, err := jpeg.Decode(img_file)
			if err != nil {
				log.Fatal(err)
			}

			bounds := img.Bounds()
			img_sizes := img_bounds{
				bounds_y_small: bounds.Dy() / 4,
				bounds_x_small: bounds.Dx() / 4,
				bounds_y_med:   bounds.Dy() / 3,
				bounds_x_med:   bounds.Dx() / 3,
				bounds_y_big:   bounds.Dy() / 2,
				bounds_x_big:   bounds.Dx() / 2,
				bounds_y_max:   bounds.Dy(),
				bounds_x_max:   bounds.Dx(),
			}

			fmt.Println(file.Name(), f, img_sizes)

			resized_image_class := resized_images{
				small:   resize.Resize(uint(img_sizes.bounds_x_small), uint(img_sizes.bounds_y_small), img, resize.Lanczos3),
				medium:  resize.Resize(uint(img_sizes.bounds_x_med), uint(img_sizes.bounds_y_med), img, resize.Lanczos3),
				big:     resize.Resize(uint(img_sizes.bounds_x_big), uint(img_sizes.bounds_y_big), img, resize.Lanczos3),
				maximum: resize.Resize(uint(img_sizes.bounds_x_max), uint(img_sizes.bounds_y_max), img, resize.Lanczos3),
			}

			save_image(resized_image_class.small, output_dir, new_img_name, "small")
			save_image(resized_image_class.medium, output_dir, new_img_name, "medium")
			save_image(resized_image_class.big, output_dir, new_img_name, "big")
			save_image(resized_image_class.maximum, output_dir, new_img_name, "maximum")

		}
	}
}

func save_image(img image.Image, output_dir, file_name, size_suffix string) {
	output_path := filepath.Join(output_dir, fmt.Sprintf("%s_%s.jpg", strings.TrimSuffix(file_name, ".jpg"), size_suffix))
	output_file, err := os.Create(output_path)
	if err != nil {
		log.Fatal(err)
	}
	defer output_file.Close()

	err = jpeg.Encode(output_file, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}