package main

import (
	"fmt"
	aStar "./algorithm/a-star"
)


const I_FORMATION = `............................
............................
.............X..............
.............X..............
.............X..............
.......E.....X.......B......
.............X..............
.............X..............
.............X..............
.............X..............
............................`


const DIAGONAL = `............................
............................
...XX................B......
.....XX.....................
.......X....................
........XX..................
......E...X.................
...........X................
............X...............
.............X..............
............................`

const NO_ROUTE = `............................
............................
.............X..............
.............X.X............
.......E.....X..X.X.........
.............X..............
........XXXXXX..............
........X....X..............
........X..B.X..............
........XXXXXX..............
............................`

const C_TYPE = `............................
............................
............................
............XX..............
..........XX................
.........X..................
........X...................
.........X..................
..........XX................
............XX...E...XX.....
..B...........XX...XX.......
................XXX.........
............................`


var routes = [...]string{C_TYPE,NO_ROUTE,I_FORMATION,DIAGONAL}

func main() {



	for _, route := range routes {
		data := aStar.ReadData(route)
		result := aStar.Astar(aStar.NewGraph(aStar.ReadData(route)))
		fmt.Println(aStar.ShowResult(data, result))
	}


}

