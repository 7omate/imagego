package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func main() {
	// Create small images with random names in targetfolder
	path := flag.String("targetfolder", "./", "Folder to receive images")
	nbimages := flag.Int("nb", 10, "Number of images to create")
	flag.Parse()
	if *nbimages == 1 {
		fmt.Printf("Creating %d image", *nbimages)
	} else {
		fmt.Printf("Creating %d images", *nbimages)
	}
	// handle distracted input w/ "/" suffix
	if !strings.HasSuffix(*path, "/") {
		p2 := strings.Join([]string{*path, "/"}, "")
		path = &p2
	}
	if *path == "./" {
		fmt.Printf(" in current folder.\n")
	} else {
		fmt.Printf(" in %s folder.\n", *path)
	}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < *nbimages; i++ {
		randomimage(*path + randname2(20))
	}
}

func randname(n int) string {
	// basic random letters
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ "
	alpha += "àbcdéèìïù"
	var res string
	for i := 0; i < n; i += 1 {
		r, _ := utf8.DecodeRuneInString(alpha[rand.Intn(len(alpha)):])
		if r != utf8.RuneError {
			res += fmt.Sprintf("%c", r)
		}
	}
	// remove leading space
	return strings.TrimLeft(res, " ")
}

func randname2(desired_length int) string {
	// a lot of characters end up being not printable...
	res := ""
	for i := 0; i < desired_length; {
		r := rune(rand.Intn(utf8.MaxRune))
		if unicode.IsPrint(r) { // Check rune printability
			res += string(r)
			i++
		}
	}
	return res
}

func randomimage(path string) {
	// see imported jpeg
	if !strings.HasSuffix(path, ".jpeg") {
		path += ".jpeg"
	}
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	var r, g, b, a uint8 = uint8(rand.Int()), uint8(rand.Int()), uint8(rand.Int()), 10
	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x += 1 {
		for y := bounds.Min.Y; y < bounds.Max.Y; y += 1 {
			img.SetRGBA(x, y, color.RGBA{r, g, b, a})
		}
	}
	w, err := os.Create(path)
	if err != nil {
		log.Println(err)
		log.Fatal(fmt.Sprintf("Failed to create file %s", path))
	}
	err = jpeg.Encode(w, img, nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to save image %s", path))
	}
}
