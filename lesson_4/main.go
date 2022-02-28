package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Point struct {
	X int `mapstructure:"xx"`
	Y int
}

func (p Point) method() {
	fmt.Println("call Point method")
}

func main() {
	pointsMap := map[string]int{
		"xx": 2,
		"y":  3,
	}
	fmt.Println(pointsMap)
	p1 := Point{}
	mapstructure.Decode(pointsMap, &p1)
	fmt.Println(p1)

	for k, v := range pointsMap {
		fmt.Println(k)
		fmt.Println(v)
	}
}

func maps() {
	pointsMap := map[string]Point{
		"b": {
			Y: 12,
			X: 13,
		},
	}
	pointsMap["a"] = Point{
		X: 1,
		Y: 2,
	}

	fmt.Println(pointsMap)
	fmt.Println(pointsMap["a"])

	value, ok := pointsMap["a"]

	if ok {
		fmt.Printf("key=%s exist in map\n", "a")
		fmt.Println(value)
	} else {
		fmt.Printf("key=%s does exist in map\n", "a")
		fmt.Println(value)
	}
}
