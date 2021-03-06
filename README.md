# bingo-tool
Tool that consumes an existing bingo board jpg and generates any number of new, randomly shuffled boards

The tool takes in one input: the filename for a config file which may or may not yet exist. If the config file does not exist (or if it's missing any fields) bingo-tool will ask the user questions and create/update the config file. The config file is simple json and can be edited by hand as needed.

Most BINGO boards have a 5x5 layout with a freespace in the middle. This tool can work on boards with arbitrary numbers of rows and columns. If there are an odd number of tiles, then the middle will be assumed to be a freespace.

Bingo-tool requires several parameters:
1) The filename of the bingo board to be used as a template.

2) A list of names - A new name.jpg file will be generated for each name in the list. 

3) The coordinates of the top left box (top left corner and bottom right corner), given that <0,0> is the top left corner of the image.

4) The coordinates of the box down and right from the first box (only the top left corner needed), given that <0,0> is the top left corner of the image.

5) Number of horizontal rows (typically five)

6) Number of vertical columns (typically five)

7) Seed Integer - This will be automatically generated on first run, and can be changed or kept by the user on subsequent runs. This allows you to have a deterministic output for subsequent runs, meaning the boards that are generated will be the same random boards each time. To generate new boards, just change or remove the number in the config.