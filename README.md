# bingo-tool
Tool that consumes an existing bingo board jpg and zero or more "partial boards" and generates any number of new, randomly shuffled boards from all of the tiles.

The tool takes in one input: the filename for a config file which may or may not yet exist. If the config file does not exist, the bingo-tool create an example config file. The config file is simple json and can be edited by hand as needed.

Most BINGO boards have a 5x5 layout with a freespace in the middle. This tool can work on boards with arbitrary numbers of rows and columns. If there are an odd number of tiles, then the middle will be assumed to be a freespace.

## Features
- Turn one bingo board into as many as you want.

- Generates new board for each person on your list, using their names in the filename.

- Requires a bit of setup but supports any template* regardless of image resolution, colors and theming, column and row size, or even number of columns and rows.

- Set an RNG seed so that you can recreate boards over and over with the same permutations (useful if you want to make a change to a board without randomizing all the generated boards every time) or set a new seed to start fresh.

- Assumes the middle space is Free Space and doesn't move it around

- Testing mode that helps you visualize how spaces are being calculated and fine-time your config as needed.

- Extra spaces: if you don't want every board to have the same boring old set of tiles, you can add more spaces to the RNG pool by adding them to a second, third, etc board. Your generated boards will all have some subset of the larger tile set.

## Config
Bingo-tool requires several parameters in the config:
- The filename of the bingo board to be used as a template.

- A list of names - A new name.png file will be generated for each name in the list. Make boards for all your friends!

- The coordinates of the top left box (top left corner and bottom right corner), given that <0,0> is the top left corner of the image.

- The coordinates of the box down and right from the first box (only the top left corner needed), given that <0,0> is the top left corner of the image.

- Number of horizontal rows (typically five)

- Number of vertical columns (typically five)

- Seed Integer - This will be automatically generated on first run, and can be changed or kept by the user on subsequent runs. This allows you to have a deterministic output for subsequent runs, meaning the boards that are generated will be the same random boards each time. To generate new boards, just change or remove the number in the config.

- Extra Squares: if you would like to add extra squares to the shuffle (such that each board will no longer be gauranteed to have the same squares as the other boards), you can do this by adding one or more Extra Squares options to the config. You should put your extra squares on an identical copy of your main template board (that is, the squares are in all the same positions), filling them in from left to right, top to bottom. The extra board does not need to be filled in all the way, you can have an arbitrary number of squares on each board. Unlike main template boards, extra square boards don't have a "Free space"; it's just a space like any other. In the config, each option should include a filename for the partial board and the number of squares the board contains.

## Usage
1) Install or update Go according to the [install instructions](https://go.dev/doc/install)
2) Download or clone this repository to a new directory
3) Open that directory in your terminal and run `go install`
	- If you're having trouble finding the installed binary in your PATH then instead run `go build` and make sure to reference the local binary when executing it
4) Run the command `bingo-tool` to create a blank config file
5) Open the config file and fill it out according to the README
6) Run the command `bingo-tool` again to consume the config and generate your boards.

## Known pain points
- Right now this is a simple cli tool that I'm not packaging, so to use it you'll need to install Go and build my tool yourself, then run it from the terminal. Or ask a friend to do it for you!

- Configuration for a board includes painstakingly recording the locations of 3 points on your template board. You only have to do it once unless your template changes.

- This tool does NOT work with bingo boards that have column/row widths/heights that change. Each tile can be any sized rectangle but they all have to be the same size; same goes for the margins between tiles.

- User has to manually open a text file and add the appropriate details while maintaining the JSON format

- Right now we're PNG only because that's what my friends and I use /shrug I don't think any of these is very difficult to work around but I acknowledge that the lack of packaging and user experience polish might be a turn off for some people.