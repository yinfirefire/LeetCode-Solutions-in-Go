package main

func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {
	//find the closet point on rectangle to the circle
	cx := x_center
	cy := y_center
	if x1 < x_center {
		cx = min(x_center, x2)
	} else {
		cx = x1
	}
	if y1 < y_center {
		cy = min(y_center, y2)
	} else {
		cy = y1
	}
	distx := x_center - cx
	disty := y_center - cy
	return (distx*distx + disty*disty) <= (radius * radius)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
