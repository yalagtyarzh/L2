package strategy

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

// PrintStrategy - интерфейс, реализации которого умеют рисовать
type PrintStrategy interface {
	Print() error
}

// ConsoleSquare - класс, рисующий квадрат в консоль
type ConsoleSquare struct {
}

// ImageSquare - класс, рисующий квадрат в формате jpg по пути, определенному в DestinationFilePath
type ImageSquare struct {
	DestinationFilePath string
}

// Print рисует квадрат в консоль
func (c *ConsoleSquare) Print() error {
	fmt.Println("square")
	return nil
}

// Print рисует квадрат в .jpg файл
func (i *ImageSquare) Print() error {
	width := 800
	height := 600

	origin := image.Point{}

	bgImage := image.NewRGBA(
		image.Rectangle{
			Min: origin,
			Max: image.Point{X: width, Y: height},
		},
	)

	bgColor := image.Uniform{C: color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	quality := &jpeg.Options{Quality: 75}

	bgRectangle := image.NewRGBA(
		image.Rectangle{
			Min: origin,
			Max: image.Point{X: width, Y: height},
		},
	)

	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{C: color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareWidth)
	square = square.Add(
		image.Point{
			X: (width / 2) - (squareWidth / 2),
			Y: (height / 2) - (squareHeight / 2),
		},
	)

	squareImg := image.NewRGBA(square)

	draw.Draw(bgImage, squareImg.Bounds(), &squareColor, origin, draw.Src)

	f, err := os.Create(i.DestinationFilePath)
	if err != nil {
		return fmt.Errorf("error opening image")
	}
	defer f.Close()

	if err = jpeg.Encode(f, bgRectangle, quality); err != nil {
		return fmt.Errorf("error writing image to disk")
	}

	return nil
}
