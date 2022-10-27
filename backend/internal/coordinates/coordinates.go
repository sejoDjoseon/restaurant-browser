package coordinates

import (
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X float64 `json:"lat"`
	Y float64 `json:"lng"`
}

func DecodePoint(data []byte) (Point, error) {
	r := string(data)
	rs := strings.Split(r, "(")
	rs = strings.Split(rs[1], ")")

	points := strings.Split(rs[0], ",")

	return PointFromStrings(points[0], points[1])
}

func PointFromStrings(xs, ys string) (Point, error) {
	point := Point{X: 0, Y: 0}
	x, err := strconv.ParseFloat(xs, 64)
	if err != nil {
		return point, nil
	}
	y, err := strconv.ParseFloat(ys, 64)
	if err != nil {
		return point, nil
	}

	return Point{X: x, Y: y}, nil
}

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// distance in meters
func Distance(p1, p2 Point) float64 {
	var la1, lo1, la2, lo2, r float64
	la1 = p1.X * math.Pi / 180
	lo1 = p1.Y * math.Pi / 180
	la2 = p2.X * math.Pi / 180
	lo2 = p2.Y * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}
