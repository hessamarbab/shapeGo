package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"os/exec"
)

func heart(i, j int) int {
	var y float32 = float32(128-i) / 100
	var x float32 = float32(j-127) / 100
	var p float32 = float32(x*x + y*y - 1)
	var d float32 = float32(x * x * y * y * y)
	var r int = int(((p * p * p) - d) * 50000)
	if r > 0 {
		return 0
	}
	return r / 2
	//return 90
}

func block(i, j int) int {
	var r int = int(i%50+j%50) * 2
	return r
}

func circle(i, j int) int {
	var r int = int((i*i)+(j*j)/2) * 2
	return r
}

func cross(i, j int) int {
	var r int = int(i*j) / 10
	return r
}

func createImage(f func(i, j int) int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 255, 255))
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			r := uint8(f(y, x))
			//a := uint8(f(y, x))
			//r := 255 - a
			//g := a / 2
			//b := (255 + a) / 2
			pixel := color.RGBA{r, 0, 0, 200}
			img.Set(x, y, pixel)
		}
	}

	return img
}

func saveImage(path string, img image.Image) {
	out, _ := os.Create(path)
	png.Encode(out, img)
	out.Close()
}

func openImage(path string) {
	err := exec.Command("/bin/display", "./_golang_image_creator.png").Run()

	if err != nil {
		log.Fatal(err)
	}
}

func removeImage(path string) {
	error := exec.Command("/bin/rm", "./_golang_image_creator.png").Run()

	if error != nil {
		log.Fatal(error)
	}
}

func main() {
	fmt.Println("Press : \n\t h for heart \n\t b for block \n\t x for cross \n\t c for circle")

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		log.Fatal(err)
	}

	// print out the unicode value i.e. A -> 65, a -> 97
	fmt.Println(char)

	var shape func(x, y int) int
	switch char {
	case 'h':
		shape = heart
		break
	case 'b':
		shape = block
		break
	case 'x':
		shape = cross
		break
	case 'c':
		shape = circle
		break
	default:
		fmt.Println("you can't add another char")
		return
	}

	img := createImage(shape)

	path := "./_golang_image_creator.png"

	saveImage(path, img)
	openImage(path)
	removeImage(path)
}
