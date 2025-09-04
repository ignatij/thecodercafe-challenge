package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func findAugmentingPath(source, sink string, residual map[string]map[string]int) map[string]string {
	queue := []string{source}
	visited := map[string]bool{}
	parent := map[string]string{}
	for len(queue) != 0 {
		curr := queue[0]
		// remove the first element
		queue = queue[1:]

		for next, cap := range residual[curr] {
			if cap > 0 && !visited[next] {
				parent[next] = curr
				if next == sink {
					return parent
				}
				visited[next] = true
				queue = append(queue, next)
			}
		}
	}
	return nil
}

func maxFlow(source, sink string, residual map[string]map[string]int) int {
	flow := 0
	for {
		parent := findAugmentingPath(source, sink, residual)
		if parent == nil {
			// all paths covered
			break
		}

		bottleneck := math.MaxInt64
		curr := sink
		for curr != source {
			next := parent[curr]
			bottleneck = min(bottleneck, residual[next][curr])
			curr = next
		}

		curr = sink
		for curr != source {
			next := parent[curr]
			residual[next][curr] -= bottleneck
			residual[curr][next] += bottleneck
			curr = next
		}

		flow += bottleneck
	}
	return flow
}

func addEdge(g map[string]map[string]int, u, v string, cap int) {
	if g[u] == nil {
		g[u] = map[string]int{}
	}
	if g[v] == nil {
		g[v] = map[string]int{}
	}
	g[u][v] = cap
	if _, ok := g[v][u]; !ok {
		g[v][u] = 0
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("error reading input file")
	}
	defer file.Close()
	b := bufio.NewScanner(file)

	graph := map[string]map[string]int{}

	var origin string
	_ = origin
	var destinations []string

	for b.Scan() {
		line := b.Text()
		if strings.Contains(line, "TRANSMISSION") {
			fields := strings.Fields(line)
			node1, node2 := fields[3], fields[5]
			cap, err := strconv.Atoi(fields[8])
			if err != nil {
				panic("can't parse capacity to int")
			}
			addEdge(graph, node1, node2, cap)

		}
		if strings.Contains(line, "ALERT") {
			parts := strings.Split(line, " ")
			origin = parts[5]
		} else if strings.Contains(line, "CRITICAL") {
			parts := strings.Split(line, "ARE")
			allFinishingNodes := strings.Split(parts[1], ",")
			for _, id := range allFinishingNodes {
				destinations = append(destinations, strings.Trim(id, " "))
			}
		}
	}

	superSink := "T"
	for _, destination := range destinations {
		addEdge(graph, destination, superSink, math.MaxInt64)
	}

	maxFlow := maxFlow(origin, superSink, graph)
	fmt.Println(maxFlow)
}
