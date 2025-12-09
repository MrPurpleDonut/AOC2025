package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"time"

	aoc "github.com/MrPurpleDonut/aoc-functions"
)

type point struct {
	x int
	y int
	z int
}

type pair struct {
	point1   *point
	point2   *point
	distance float64
}

type collect struct {
	vals []*point
	tag  int
}

func (c *collect) merge(b *collect) {
	for _, v := range b.vals {
		if !c.contains(v) {
			c.vals = append(c.vals, v)
		}
	}
}

func (c *collect) contains(v *point) bool {
	for _, p2 := range c.vals {
		if v.x == p2.x && v.y == p2.y && v.z == p2.z {
			return true
		}
	}
	return false
}

func (c *collect) isOverlap(b *collect) bool {
	for _, v := range b.vals {
		if c.contains(v) {
			return true
		}
	}
	return false
}

func main() {
	start := time.Now()

	rows, err := aoc.MakeRows(os.Args[1])
	aoc.HandleError(err)
	points := make([]point, 0)
	for _, p := range rows {
		nums, err := aoc.ParseAllInts(p)
		aoc.HandleError(err)
		points = append(points, point{nums[0], nums[1], nums[2]})
	}

	pairs := make([]pair, 0)
	for i := 0; i < len(points); i++ {
		for j := 0; j < i; j++ {
			p1 := &points[i]
			p2 := &points[j]
			x := (p1.x - p2.x) * (p1.x - p2.x)
			y := (p1.y - p2.y) * (p1.y - p2.y)
			z := (p1.z - p2.z) * (p1.z - p2.z)
			distance := math.Sqrt(float64(x + y + z))
			pairs = append(pairs, pair{p1, p2, distance})
		}
	}

	slices.SortFunc(pairs, func(a, b pair) int {
		if a.distance < b.distance {
			return -1
		}
		return 1
	})
	allCollections := make([]collect, 0)

	tag := 0

	for i := range len(pairs) {
		if i == 1000 {

			slices.SortFunc(allCollections, func(a, b collect) int { return len(b.vals) - len(a.vals) })

			product := len(allCollections[0].vals) * len(allCollections[1].vals) * len(allCollections[2].vals)

			fmt.Println(product)
		}
		p := pairs[i]
		contained := false

		collectionsToMerge := make([]*collect, 0)
		for idx := range allCollections {
			c := &allCollections[idx]
			if c.contains(p.point1) || c.contains(p.point2) {
				c.merge(&collect{[]*point{p.point1, p.point2}, tag})
				tag++
				contained = true
				collectionsToMerge = append(collectionsToMerge, c)
			}
		}

		if !contained {
			allCollections = append(allCollections, collect{[]*point{p.point1, p.point2}, tag})
			tag++
			continue
		}

		// Merge all collections into the first one
		parent := collectionsToMerge[0]
		badIndex := make([]int, 0)
		for in, c := range collectionsToMerge {
			if in == 0 {
				continue
			}
			parent.merge(c)
			index := index(allCollections, *c)
			if index == -1 {
				panic("AHH")
			}
			badIndex = append(badIndex, index)
		}

		// Remove merged collections
		slices.Sort(badIndex)
		for idx := len(badIndex) - 1; idx >= 0; idx-- {
			allCollections = slices.Delete(allCollections, badIndex[idx], badIndex[idx]+1)
		}

		if len(collectionsToMerge) > 1 {
			// After merging, resort to find the largest
			slices.SortFunc(allCollections, func(a, b collect) int {
				return len(b.vals) - len(a.vals)
			})
		}

		if len(allCollections[0].vals) == 1000 {
			fmt.Println(p.point1.x * p.point2.x)
			break
		}

	}

	fmt.Println(time.Since(start))
}

func index(col []collect, c collect) int {
	for i, v := range col {
		if c.tag == v.tag {
			return i
		}
	}
	return -1
}
