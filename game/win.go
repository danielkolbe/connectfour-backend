package game


// winH returns a color if four contiguous fields of the
// board with the same color in a horizontal line exist. It returns
// none otherwise. 
// It is assumed (but not checked) that the condition applies for at most one of both colors.
func winH(b *Board) color {
	color := none
	for _, row := range b.Fields {
		color = hasFour(row[:])
		if none != color {
			break
		}
	}
	return color
}

// hasFour is the working horse of the win detection algorithm.
// It returns the (first) color (red or blue) that is value of
// four contiguous items (like i_k, i_k+1, i_k+2, i_k+3) of the given slice.
// If non-existent none will be returned.
// It is assumed (but not checked) that the condition applies for at most one of both colors.
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
