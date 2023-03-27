package main

const (
	configTemplate = `{
	"filepath": "./your_board.png",
	"names": [
		"Steve",
		"Caleb"
	],
	"numColumns": 5,
	"numRows": 5,
	"firstRect": {
		"origin": {
			"X":57,
			"Y":410
		},
		"oppositeCorner": {
			"X": 374,
			"Y": 726
		}
	},
	"nextRectOrigin": {
		"X": 427,
		"Y": 778
	},
	"testing": false,
	"seedInteger": 39599,
	"extraSquares": [
		{
			"filepath": "./your_extra_board.png",
			"numOfSquares": 20	
		}
	]
}`
)
