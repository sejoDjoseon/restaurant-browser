package coordinates

import (
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
	var x, y float64
	var err error
	point := Point{X: 0, Y: 0}
	x, err = strconv.ParseFloat(points[0], 64)
	if err != nil {
		return point, nil
	}
	y, err = strconv.ParseFloat(points[1], 64)
	if err != nil {
		return point, nil
	}

	return Point{X: x, Y: y}, nil
}
