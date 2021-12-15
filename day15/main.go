package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

func run() error {
	m, err := parse("day15/input.txt")
	if err != nil {
		return err
	}
	gold(m)
	return nil
}

func parse(filename string) ([][]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	m := make([][]int, 0)
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		raw := scan.Text()
		line := make([]int, 0, len(raw))
		for _, ch := range raw {
			val, err := strconv.Atoi(string(ch))
			if err != nil {
				return nil, err
			}
			line = append(line, val)
		}
		m = append(m, line)
	}

	return m, nil
}

type Point struct {
	Pos Position
	Val int
	Distance int
	HeapIndex int
	Discovered bool
}

func (p *Point) GetNeighbours(d [][]*Point) []*Point {
	i,j := p.Pos.I, p.Pos.J
	points := make([]*Point, 0)
	if i > 0 {
		points = append(points, d[i-1][j])
	}
	if i < len(d) - 1 {
		points = append(points, d[i+1][j])
	}
	if j > 0 {
		points = append(points, d[i][j-1])
	}
	if j < len(d[i]) - 1 {
		points = append(points, d[i][j+1])
	}

	return points
}

func silver(m [][]int) {
	d := make([][]*Point, 0 , len(m))
	for i := range m {
		line := make([]*Point, 0, len(m[i]))
		for j := range m[i]{
			line = append(line, &Point{
				Pos: Position{I:i,J:j},
				Val:      m[i][j],
				Distance: math.MaxInt,
				HeapIndex: -1,
			})
		}
		d = append(d, line)
	}

	pq := NewPriorityQueue()
	start := d[0][0]
	start.Distance = 0
	pq.Push(start)

	end := d[len(m)-1][len(m[0])-1]
	for pq.Len() > 0 {
		point := pq.Pop()
		point.Discovered = true
		if point == end {
			break
		}

		currDist := point.Distance
		for _, nextPoint := range point.GetNeighbours(d) {
			minDistance := min(nextPoint.Distance, currDist + nextPoint.Val)
			if minDistance != nextPoint.Distance {
				nextPoint.Distance = minDistance
				if pq.Has(nextPoint.Pos) {
					pq.Update(nextPoint.Pos)
				}
			}
			if !nextPoint.Discovered && !pq.Has(nextPoint.Pos) {
				pq.Push(nextPoint)
			}
		}
	}

	fmt.Printf("\n final distance is %d", end.Distance)
}

func gold(m [][]int) {
	iLen := len(m)
	jLen := len(m[0])
	newMap := make([][]int, iLen*5)

	for i := 0; i < len(newMap); i++ {
		newMap[i] = make([]int, jLen*5)
	}

	for i := range m {
		for j := range m[i] {
			newMap[i][j] = m[i][j]
		}
	}


	// grow first line
	for c := 1; c <= 4; c++ {
		for i := 0; i < iLen; i++ {
			for j := c*jLen; j < c*jLen + jLen; j++ {
				nextVal := newMap[i][j-jLen]
				nextVal++
				if nextVal > 9 {
					nextVal = 1
				}
				newMap[i][j] = nextVal
			}
		}
	}

	for l := 1; l <= 4; l++ {
		for i := l * iLen; i < l*iLen + iLen; i++ {
			for j := 0; j < len(newMap[i]); j++ {
				nextVal := newMap[i - iLen][j]
				nextVal++
				if nextVal > 9 {
					nextVal = 1
				}
				newMap[i][j] = nextVal
			}
		}
	}

	fmt.Printf("\n\n")
	for i := range newMap{
		fmt.Printf("\n")
		for j := range newMap[i] {
			fmt.Printf("%d", newMap[i][j])
		}
	}

	silver(newMap)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

type Position struct {
	I int
	J int
}

func NewPriorityQueue() *PriorityQueue {
	h := &PointHeap{}
	heap.Init(h)
	return &PriorityQueue{
		h:     h,
		nodes: make(map[Position]*Point),
	}
}

type PriorityQueue struct {
	h *PointHeap
	nodes map[Position]*Point
}

func (p *PriorityQueue) Len() int {
	return p.h.Len()
}

func (p *PriorityQueue) Push(point *Point) {
	if _, ok := p.nodes[point.Pos]; !ok {
		heap.Push(p.h, point)
		p.nodes[point.Pos] = point
	}
}

func (p *PriorityQueue) Pop() *Point {
	point := heap.Pop(p.h).(*Point)
	delete(p.nodes, point.Pos)
	return point
}

func (p *PriorityQueue) Has(pos Position) bool {
	return p.nodes[pos] != nil
}

func (p *PriorityQueue) Update(pos Position) {
	point := p.nodes[pos]
	heap.Fix(p.h, point.HeapIndex)
}

type PointHeap []*Point

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i].Distance < h[j].Distance }
func (h PointHeap) Swap(i, j int)      {
	h[i], h[j] = h[j], h[i]
	h[i].HeapIndex = i
	h[j].HeapIndex = j
}

func (h *PointHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	val := x.(*Point)
	val.HeapIndex = len(*h)
	*h = append(*h, val)
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	x.HeapIndex = -1
	old[n-1] = nil // avoid mem leak
	*h = old[0 : n-1]
	return x
}