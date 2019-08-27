package golang_leetcode

import "math"

type coordinate struct {
	x float64
	y float64
}

type polygon struct {
	coordinates []coordinate
	distances []float64
}

func(p *polygon) getPerimeter() float64 {
	res := 0.0
	n := len(p.coordinates)
	distances := make([]float64, n)
	for i := 0; i < n-1; i++ {
		diffX := p.coordinates[i+1].x - p.coordinates[i].x
		diffY := p.coordinates[i+1].y - p.coordinates[i].y
		distance := math.Sqrt(diffX * diffX + diffY * diffY)
		res += distance
		distances = append(distances, distance)
	}
	diffX := p.coordinates[0].x - p.coordinates[n-1].x
	diffY := p.coordinates[0].y - p.coordinates[n-1].y
	distance := math.Sqrt(diffX * diffX + diffY * diffY)
	res += distance
	distances = append(distances, distance)
	p.distances = distances
	return res
}

func (p *polygon) getKDivideCoordinates(k int) []coordinate {
	res := make([]coordinate, k)
	res[0] = p.coordinates[0]
	perimeter := p.getPerimeter()
	oneLength := perimeter/float64(k)
	nextCoordinate := 1
	calculateLength := oneLength
	for i := 1; i < k; i++ {
		if p.distances[nextCoordinate-1] == calculateLength {
			res = append(res, p.coordinates[nextCoordinate])
			nextCoordinate++
			calculateLength = oneLength
		} else if p.distances[nextCoordinate-1] > calculateLength {
			diffX := p.coordinates[nextCoordinate].x - p.coordinates[nextCoordinate-1].x
			x := p.coordinates[nextCoordinate-1].x + calculateLength/p.distances[nextCoordinate-1] * diffX
			diffY := p.coordinates[nextCoordinate].y - p.coordinates[nextCoordinate-1].y
			y := p.coordinates[nextCoordinate-1].y + calculateLength/p.distances[nextCoordinate-1] * diffY
			res[i] = coordinate{
				x: x,
				y: y,
			}
			calculateLength = p.distances[nextCoordinate-1] - calculateLength
		} else {
			nextCoordinate++
			calculateLength = calculateLength - p.distances[nextCoordinate-1]
		}
	}
	return res
}