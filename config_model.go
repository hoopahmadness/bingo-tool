package main

import (
	"encoding/json"
	"fmt"
	"image"
)

type Config struct {
	Filepath       string      `json:"filepath"`
	Names          []string    `json:"names"`
	NumRows        int         `json:"numRows"`
	NumColumns     int         `json:"numColumns"`
	NextRectOrigin image.Point `json:"nextRectOrigin"`
	Seed           int64       `json:"seedInteger"`
	Test           bool        `json:"testing"`
	FirstRect      struct {
		Origin         image.Point `json:"origin"`
		OppositeCorner image.Point `json:"oppositeCorner"`
	} `json:"firstRect"`
	ExtraSquares []ExtraSquare `json:"extraSquares"`
}

type ExtraSquare struct {
	Filepath     string `json:"filepath"`
	NumOfSquares int    `json:"numOfSquares"`
}

func testConfig() Config {
	test := Config{}
	//add test data
	return test
}

func parseConfig(jsonStr string) Config {
	var newConfig Config
	newConfig.ExtraSquares = []ExtraSquare{}
	err := json.Unmarshal([]byte(jsonStr), &newConfig)
	if err != nil {
		fmt.Println("ERROR UNMARSHALLING")
		panic(err.Error())
	}
	return newConfig

}
