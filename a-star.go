package a_star

import (
	"strings"
)


//////////////////////////////////// VARIABLES
const (
	UNKNOWN int = iota - 1
	LAND
	WALL
	BEGIN
	END
)

type Graph struct {
	start, stop *Node
	nodes       []*Node
	data        *MatrixData
}

type Node struct {
	X, Y   int
	parent *Node
	H      int //aproximate distance
	cost   int //Path cost for this node
}

type MatrixData [][]int

//////////////////////////////////// INPUT
func createMatrixData(rows, cols int) *MatrixData {
	result := make(MatrixData, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]int, cols)
	}
	return &result
}

//////////////////////////////////// NODES
func createNode(x, y int) *Node {
	node := &Node{
		X:      x,
		Y:      y,
		parent: nil,
		H:      0,
		cost:   0,
	}
	return node
}

func (self *Graph) Node(x, y int) *Node {

	for _, n := range self.nodes {
		if n.X == x && n.Y == y {
			return n
		}
	}
	map_data := *self.data
	if map_data[x][y] == LAND || map_data[x][y] == END {

		n := createNode(x, y)
		self.nodes = append(self.nodes, n)
		return n
	}
	return nil
}

func removeNode(nodes []*Node, node *Node) []*Node {
	ith := -1
	for i, n := range nodes {
		if n == node {
			ith = i
			break
		}
	}
	if ith != -1 {
		copy(nodes[ith:], nodes[ith+1:])
		nodes = nodes[:len(nodes)-1]
	}
	return nodes
}

func hasNode(nodes []*Node, node *Node) bool {
	for _, n := range nodes {
		if n == node {
			return true
		}
	}
	return false
}

//////////////////////////////////// PATH
func rebuildPath(current_node *Node) []*Node {
	var path []*Node
	path = append(path, current_node)
	for current_node.parent != nil {
		path = append(path, current_node.parent)
		current_node = current_node.parent
	}
	//Reverse path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func (self *Graph) adjust(node *Node) []*Node {
	var result []*Node
	map_data := *self.data
	rows := len(map_data)
	cols := len(map_data[0])

	if node.X <= rows && node.Y+1 < cols {
		if new_node := self.Node(node.X, node.Y+1); new_node != nil {
			result = append(result, new_node)
		}
	}
	if node.X <= rows && node.Y-1 >= 0 {
		new_node := self.Node(node.X, node.Y-1)
		if new_node != nil {
			result = append(result, new_node)
		}
	}
	if node.Y <= cols && node.X+1 < rows {
		new_node := self.Node(node.X+1, node.Y)
		if new_node != nil {
			result = append(result, new_node)
		}
	}
	if node.Y <= cols && node.X-1 >= 0 {
		new_node := self.Node(node.X-1, node.Y)
		if new_node != nil {
			result = append(result, new_node)
		}
	}
	return result
}

//////////////////////////////////// CALCULATION H

func minH(nodes []*Node) *Node {
	if len(nodes) == 0 {
		return nil
	}
	result_node := nodes[0]
	minH := result_node.H
	for _, node := range nodes {
		if node.H < minH {
			minH = node.H
			result_node = node
		}
	}
	return result_node
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//////////////////////////////////// API
func Astar(graph *Graph) []*Node {
	var path, openSet, closedSet []*Node

	openSet = append(openSet, graph.start)
	for len(openSet) != 0 {

		current := minH(openSet)
		if current.parent != nil {
			current.cost = current.parent.cost + 1
		}
		if current == graph.stop {
			return rebuildPath(current)
		}
		openSet = removeNode(openSet, current)
		closedSet = append(closedSet, current)
		for _, tile := range graph.adjust(current) {
			if tile != nil && graph.stop != nil && !hasNode(closedSet, tile) {
				tile.H = Heuristic(graph, tile) + current.cost
				if !hasNode(openSet, tile) {
					openSet = append(openSet, tile)
				}
				tile.parent = current
			}
		}
	}
	return path
}

func ReadData(map_str string) *MatrixData {
	rows := strings.Split(map_str, "\n")
	if len(rows) == 0 {
		panic("Matrix needs to have at least 1 row")
	}

	row_count := len(rows)
	col_count := len(rows[0])

	result := *createMatrixData(row_count, col_count)
	for i := 0; i < row_count; i++ {
		for j := 0; j < col_count; j++ {
			char := rows[i][j]
			switch char {
			case '.':
				result[i][j] = LAND
			case 'X':
				result[i][j] = WALL
			case 'B':
				result[i][j] = BEGIN
			case 'E':
				result[i][j] = END
			}
		}
	}
	return &result
}

func ShowResult(data *MatrixData, nodes []*Node) string {
	var result string
	for i, row := range *data {
		for j, cell := range row {
			added := false
			for _, node := range nodes {
				if node.X == i && node.Y == j {
					result += "*"
					added = true
					break
				}
			}
			if !added {
				switch cell {
				case LAND:
					result += "."
				case WALL:
					result += "X"
				case BEGIN:
					result += "B"
				case END:
					result += "E"
				default: //Unknown
					result += "?"
				}
			}
		}
		result += "\n"
	}
	return result
}

func NewGraph(map_data *MatrixData) *Graph {
	var start, stop *Node
	var nodes []*Node
	for i, row := range *map_data {
		for j, _type := range row {
			if _type == BEGIN || _type == END {
				node := createNode(i, j)
				nodes = append(nodes, node)
				if _type == BEGIN {
					start = node
				}
				if _type == END {
					stop = node
				}
			}
		}
	}
	g := &Graph{
		nodes: nodes,
		start: start,
		stop:  stop,
		data:  map_data,
	}
	return g
}

func Heuristic(graph *Graph, tile *Node) int {
	return abs(graph.stop.X-tile.X) + abs(graph.stop.Y-tile.Y)
}

