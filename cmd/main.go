package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"

	"github.com/romanzaycev/yet-another-tiny-raytracer/src/scenes"
	"github.com/romanzaycev/yet-another-tiny-raytracer/src/tinyraytracer"
)

var (
	outWidth           int
	outHeight          int
	outFile            string
	background         string
	checkerboardColorA string
	checkerboardColorB string
)

func init() {
	flag.IntVar(&outWidth, "iw", 1000, "output image width")
	flag.IntVar(&outHeight, "ih", 1000, "output image height")
	flag.StringVar(&outFile, "o", "out.ppm", "output filename")
	flag.StringVar(&background, "b", "f2f0fc", "environment color (background)")
	flag.StringVar(&checkerboardColorA, "cca", "ecf0f1", "checkerboard color A")
	flag.StringVar(&checkerboardColorB, "ccb", "2c3e50", "checkerboard color B")
}

func main() {
	flag.Parse()

	if err := _main(os.Stdin, os.Stdout, os.Args); err != nil {
		log.Fatal(err)
	}
}

func _main(in io.Reader, out io.Writer, args []string) error {
	config := tinyraytracer.NewConfig()
	config.OutImageHeight = outHeight
	config.OutImageWidth = outWidth
	config.OutImageName = outFile

	envColor, err := tinyraytracer.HexRgbToComponent(background)
	if err != nil {
		return err
	}

	cColorA, err := tinyraytracer.HexRgbToComponent(checkerboardColorA)
	if err != nil {
		return err
	}

	cColorB, err := tinyraytracer.HexRgbToComponent(checkerboardColorB)
	if err != nil {
		return err
	}

	config.EnvironmentColor = envColor
	config.CheckerboardColorA = cColorA
	config.CheckerboardColorB = cColorB

	scene := scenes.SpheresScene()
	renderer := tinyraytracer.NewRaytracer(config)

	result, err := renderer.Render(scene)
	if err != nil {
		return err
	}

	file, err := os.Create(config.OutImageName)
	if err != nil {
		return err
	}

	defer func() {
		err := file.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	writeBuffer := bufio.NewWriter(file)
	err = tinyraytracer.WritePpmFile(config.OutImageWidth, config.OutImageHeight, result, writeBuffer)
	if err != nil {
		return err
	}

	return nil
}
