package p130

type Coordinate struct {
	Row int
	Col int
}

func Invoke(maze [][]string, startRow int, startCol int, CoordinateRow []int, CoordinateCol []int) int {
	moves := make([]Coordinate, len(CoordinateRow))
	for i, _ := range CoordinateRow {
		moves[i] = Coordinate{Row: CoordinateRow[i], Col: CoordinateCol[i]}
	}

	graph := make([][][]Coordinate, len(maze))
	for i, _ := range graph {
		graph[i] = make([][]Coordinate, len(maze[0]))
	}
	for rowI, row := range maze {
		for colI, _ := range row {
			for _, move := range moves {
				graph[rowI][colI] = append(graph[rowI][colI], Coordinate{
					Row: rowI + move.Row,
					Col: colI + move.Col,
				})
			}
		}
	}

	return 0
}
