package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"math/rand"
	"os"
)

var config Config

func main() {
	//parse arguments from json file
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config = readConfig(pwd)
	//stop here and ask to continue, else end

	//import image as go image
	board := importJPG(config.Filepath)
	//generate rectangles from input points
	firstTile := Tile{config.FirstRect.Origin, config.FirstRect.OppositeCorner}
	tileArray := generateTileArray(firstTile, config.NextRectOrigin, config.NumRows, config.NumColumns)

	//create window showing image with rectangles highlighted, ask to continue
	//save subsets of bingo board into array
	subImageArray := generateSubImageArray(tileArray, board)

	r := rand.New(rand.NewSource(config.Seed))
	permutations := generatePermutation(r, config.NumRows, config.NumColumns, len(config.Names))

	//loop over list of names
	for perm, person := range config.Names {
		shuffledArr := permutations[perm]
		fmt.Println(person)
		fmt.Println(shuffledArr)
		//shuffle board
		newBoard := shuffleBoard(board, subImageArray, tileArray, shuffledArr)

		//save new copy of image
		writeImage(newBoard, fmt.Sprintf("%s.jpg", person))
	}
}

//function to read config from file (filename string) Config
func readConfig(pwd string) Config {
	configPath := pwd + "/bingo-config.json"
	dat, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Writing template config file to your current directory")
		err = ioutil.WriteFile(configPath, []byte(configTemplate), 0644)
		if err != nil {
			panic(err.Error())
		}
		os.Exit(0)
	}
	// fmt.Println(string(dat))
	return parseConfig(string(dat))
}

//function to get photo (filename string) go image
func importJPG(filename string) draw.Image {
	existingImageFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer existingImageFile.Close()

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := jpeg.Decode(existingImageFile)
	if err != nil {
		panic(err)
	}

	b := loadedImage.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(m, m.Bounds(), loadedImage, b.Min, draw.Src)
	return m

}

//function to create array of rectangles (first Rectangle, nextCorner Point) []Tiles
func generateTileArray(first Tile, nextCorner image.Point, rows, columns int) []Tile {
	origin := first.Origin

	tileWidth := first.OppositeCorner.X - first.Origin.X
	tileHeight := first.OppositeCorner.Y - first.Origin.Y

	gapWidth := nextCorner.X - first.OppositeCorner.X
	gapHeight := nextCorner.Y - first.OppositeCorner.Y

	tileArray := []Tile{}
	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			// index = 5*row + column
			tileOriginX := origin.X + column*(tileWidth+gapWidth)
			tileOriginY := origin.Y + row*(tileHeight+gapHeight)

			tileOppositeX := tileOriginX + tileWidth
			tileOppositeY := tileOriginY + tileHeight

			tileArray = append(tileArray, newTile(tileOriginX, tileOriginY, tileOppositeX, tileOppositeY))
		}
	}
	return tileArray
}

func generateSubImageArray(tiles []Tile, img draw.Image) []draw.Image {
	imgArray := []draw.Image{}
	for _, tile := range tiles {
		imgArray = append(imgArray, getSubImage(img, tile))
	}
	return imgArray
}

//function to create subset of image (image goimage, bounds Rectangle) goimage
func getSubImage(img draw.Image, bounds Tile) (subImage *image.RGBA) {
	x, y, _ := bounds.getDimensions()
	subImage = image.NewRGBA(image.Rect(0, 0, x, y))

	for row := 0; row < y; row++ {
		for column := 0; column < x; column++ {
			subImage.Set(column, row, img.At(column+bounds.Origin.X, row+bounds.Origin.Y))
		}
	}
	return
}

func writeImage(img draw.Image, filename string) {
	// fmt.Println("Attempting to write " + filename)

	// outputFile is a File type which satisfies Writer interface
	outputFile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	jpeg.Encode(outputFile, img, nil)

	// Don't forget to close files
	outputFile.Close()

}

// takes in a Rand object and dimensions of the board and number of Names, calculates all random permutation of tile indices
func generatePermutation(r *rand.Rand, rows, columns, num int) [][]int {
	indices := []int{}
	permutations := [][]int{}
	useFreespace := false
	freespace := 0
	for ii := 0; ii < columns*rows; ii++ {
		indices = append(indices, ii)
		useFreespace = !useFreespace
	}

	if useFreespace {
		freespace = (columns*rows - 1) / 2
		fmt.Sprintf("freespace is...%d\n", freespace)
		indices = append(indices[:freespace], indices[freespace+1:]...) //slice out the freespace before shuffling
	}

	for ii := 1; ii <= num; ii++ {
		shuffledIndices := []int{}
		for _, i := range r.Perm(len(indices)) {
			shuffledIndices = append(shuffledIndices, indices[i])
		}

		if useFreespace {
			shuffledIndices = append(shuffledIndices, 0) //lengthen the array
			copy(shuffledIndices[freespace+1:], shuffledIndices[freespace:])
			shuffledIndices[freespace] = freespace
		}
		permutations = append(permutations, shuffledIndices)
	}

	return permutations
}

//function to create new image from subsets (main goimage, tiles []goimage, rects []Rectangle) goimage
func shuffleBoard(board draw.Image, images []draw.Image, tiles []Tile, newIndices []int) draw.Image {

	//	//loop over array
	for newIndex, shuffledIndex := range newIndices {
		tile := tiles[newIndex]
		subImage := images[shuffledIndex]
		//place subimages in new locations
		_, _, sr := tile.getDimensions()
		r := image.Rectangle{tile.Origin, tile.Origin.Add(sr.Size())}
		draw.Draw(board, r, subImage, sr.Min, draw.Src)

	}

	return board
}
