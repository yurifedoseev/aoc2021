package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func run() error {
	graph, err := parse("day12/input.txt")
	if err != nil {
		return err
	}
	fmt.Printf("\n found %d vertixes", len(graph.Vertexes))

	gold(graph)
	return nil
}

func silver(g *Graph) {
	t := &Traverse{
		Visited: make(map[string]bool),
		Path:    make([]string, 0),
		Result: make([]string, 0),
	}

	findAllPaths(t, g.Start, g.End)
	for _, p := range t.Result {
		fmt.Printf("\n%s", p)
	}
	fmt.Printf("\n found %d path!", len(t.Result))
}

func findAllPaths(t *Traverse, s *Node, d *Node) {
	if !s.IsBig() {
		t.Visited[s.Name] = true
	}
	t.Path = append(t.Path, s.Name)

	if s == d {
		path := strings.Join(t.Path, ",")
		t.Result = append(t.Result, path)
	} else {
		for _, v := range s.Vertexes {
			if !t.Visited[v.Name] {
				findAllPaths(t, v, d)
			}
		}
	}

	t.Path = t.Path[:len(t.Path) - 1]
	t.Visited[s.Name] = false
}

type TraverseV2 struct {
	Graph *Graph
	Visited map[string]int
	Path []string
	Result []string
	DoubleNode string
	StartVisited bool
	EndVisited bool
}

func gold(g *Graph) {
	t := &TraverseV2{
		Graph: g,
		Visited: make(map[string]int),
		Path:    make([]string, 0),
		Result: make([]string, 0),
		DoubleNode: "",
	}

	findAllPathsGold(t, g.Start, g.End)
	for _, p := range t.Result {
		fmt.Printf("\n%s", p)
	}
	fmt.Printf("\n found %d path!", len(t.Result))
}

func findAllPathsGold(t *TraverseV2, s *Node, d *Node) {
	if !s.IsBig() {
		t.Visited[s.Name]++
		if t.Visited[s.Name] == 2 {
			t.DoubleNode = s.Name
		}
	}
	t.Path = append(t.Path, s.Name)

	if s == t.Graph.Start {
		t.StartVisited = true
	}

	if s == t.Graph.End {
		t.EndVisited = true
	}

	if s == d {
		path := strings.Join(t.Path, ",")
		t.Result = append(t.Result, path)
	} else {
		for _, v := range s.Vertexes {
			hasVisit := true
			if v.IsBig() {
				hasVisit = true
			} else {
				hasVisit = t.Visited[v.Name] == 0 || (t.Visited[v.Name] == 1 && t.DoubleNode == "")
			}

			if v == t.Graph.End && t.EndVisited {
				hasVisit = false
			}

			if v == t.Graph.Start && t.StartVisited {
				hasVisit = false
			}

			if hasVisit {
				findAllPathsGold(t, v, d)
			}
		}
	}

	t.Path = t.Path[:len(t.Path) - 1]
	t.Visited[s.Name]--
	if t.Visited[s.Name] < 0 {
		t.Visited[s.Name] = 0
	}
	if t.DoubleNode == s.Name {
		t.DoubleNode = ""
	}

	if s == t.Graph.Start {
		t.StartVisited = false
	}

	if s == t.Graph.End {
		t.EndVisited = false
	}
}

type Node struct {
	Name string
	Vertexes []*Node
}

func (n *Node) IsBig() bool {
	return n.Name == strings.ToUpper(n.Name)
}

type Graph struct {
	Start *Node
	End *Node
	Vertexes map[string]*Node
}

type Traverse struct {
	Visited map[string]bool
	Path []string
	Result []string
}

func parse(filename string) (*Graph, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	graph := &Graph{
		Start:    nil,
		End:      nil,
		Vertexes: make(map[string]*Node),
	}
	graph.End = &Node{Name: "end", Vertexes: make([]*Node, 0)}
	graph.Vertexes[graph.End.Name] = graph.End
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		raw := strings.Split(scanner.Text(), "-")
		if _, ok := graph.Vertexes[raw[0]]; !ok {
			graph.Vertexes[raw[0]] = &Node{
				Name: raw[0],
				Vertexes: make([]*Node, 0),
			}
		}
		nodeA := graph.Vertexes[raw[0]]
		if nodeA.Name == "start" {
			graph.Start = nodeA
		}

		if _, ok := graph.Vertexes[raw[1]]; !ok {
			graph.Vertexes[raw[1]] = &Node{
				Name: raw[1],
				Vertexes: make([]*Node, 0),
			}
		}

		nodeB := graph.Vertexes[raw[1]]
		nodeA.Vertexes = append(nodeA.Vertexes,nodeB)
		nodeB.Vertexes = append(nodeB.Vertexes, nodeA)
	}

	return graph, nil
}

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}
