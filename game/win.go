package game

// winH returns a color if four contiguous fields of the
// board of that color in a horizontal line exist. It returns none otherwise.
// If this condition applies for both colors one of both will be returned.
func winHorizontal(b *Board) color {
	color := none
	for _, row := range b.Fields {
		color = hasFour(row[:])
		if none != color {
			break
		}
	}
	return color
}

// winVertical returns a color if four contiguous fields of the
// board of that color in a vertical line exist. It returns none otherwise. 
// If this condition applies for both colors one of both will be returned.
func winVertical(b *Board) color {
	color := none
	for index := range b.Fields[0] {
		color = hasFour(column(&b.Fields, index))
		if none != color {
			break
		}
	}
	return color
}

// winDiagonal returns a color if four contiguous fields of the
// board of that color in a vertical line exist. It returns none otherwise. 
// If this condition applies for both colors one of both will be returned.
func winDiagonal(b *Board) color {
	color := none
	for y := 0; y <= len(b.Fields)-4; y++ {
		color = hasFour(diagonal(&b.Fields, y, 0))
		if none != color {
			return color
		}
	}
	for x := 0; x <= len(b.Fields[0])-4; x++ {
		color = hasFour(diagonal(&b.Fields, 0, x))
		if none != color {
			return color
		}
	}
	
	return none
}


// hasFour is the working horse of the win detection algorithm.
// It returns the (first) color (red or blue) that is value of
// four contiguous items (like i_k, i_k+1, i_k+2, i_k+3) of the given slice.
// If non-existent none will be returned.
// If this condition applies for both colors one of both will be returned.
func hasFour(c []color) color {
	current := none
	count := 0
	for _, color := range c {
		if current != color || none == color {
			count = 0
		}
		count++
		if 4 == count {
			return color
		}
		current = color
	}
	return none
}

// column returns the column with given index of the given array
func column(fields *[nRows][nCols] color, index int) []color {
	column := make([]color, nRows)
	for _, row := range fields {
		column = append(column, row[index])
	}
	return column
}

// diagonal returns the diagonal of the given array starting at (row, column)
func diagonal(fields *[nRows][nCols] color, row int, column int) []color {
	diagonal := make([]color, 0)
	x := column
	for y := row; y < len(fields); y++ {
		if x < len(fields[0]) {
			diagonal = append(diagonal, fields[y][x])
		} 
		x++
	}
	return diagonal
}