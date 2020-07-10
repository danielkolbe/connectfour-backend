package game


// hasFour returns the color (red or blue) that is present
// at least four contiguous items (like i_k, i_k+1, i_k+2, i_k3) of the given slice.
// If non-existent none will be returned. 
// It is assumed that the condition applies for at most one of both colors.
func hasFour (c *[]color) color {
	current := none
	count := 0
	for _, color := range *c {
		 if(current != color || none == color) {
			 count = 0
		 }
		 count ++
		 if 4 == count {
			 return color
		 }
		 current = color	
	}
	return none
}
