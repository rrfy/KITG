package main

import (
	"fmt"
	"math"
)

// Edge представляет ребро графа с указанием мощности и потока
type Edge struct {
	to, rev   int
	cap, flow int
}

// Graph представляет сеть с рёбрами
type Graph struct {
	n   int
	adj [][]Edge
}

// NewGraph инициализирует граф с n вершинами
func NewGraph(n int) *Graph {
	return &Graph{n: n, adj: make([][]Edge, n)}
}

// AddEdge добавляет направленное ребро от вершины u к вершине v с мощностью cap
func (g *Graph) AddEdge(u, v, cap int) {
	g.adj[u] = append(g.adj[u], Edge{v, len(g.adj[v]), cap, 0})
	g.adj[v] = append(g.adj[v], Edge{u, len(g.adj[u]) - 1, 0, 0}) // Обратное ребро с мощностью 0
}

// bfs строит уровневую сеть
func (g *Graph) bfs(s, t int, level []int) bool {
	for i := range level {
		level[i] = -1
	}
	level[s] = 0
	queue := []int{s}
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, e := range g.adj[u] {
			if level[e.to] < 0 && e.flow < e.cap {
				level[e.to] = level[u] + 1
				queue = append(queue, e.to)
			}
		}
	}
	return level[t] >= 0
}

// dfs находит блокирующий поток
func (g *Graph) dfs(u, t int, flow int, level []int, start []int) int {
	if u == t {
		return flow
	}
	for start[u] < len(g.adj[u]) {
		e := &g.adj[u][start[u]]
		if level[e.to] == level[u]+1 && e.flow < e.cap {
			currFlow := min(flow, e.cap-e.flow)
			tempFlow := g.dfs(e.to, t, currFlow, level, start)
			if tempFlow > 0 {
				e.flow += tempFlow
				g.adj[e.to][e.rev].flow -= tempFlow
				return tempFlow
			}
		}
		start[u]++
	}
	return 0
}

// MaxFlow вычисляет максимальный поток из s в t
func (g *Graph) MaxFlow(s, t int) int {
	if s == t {
		return 0
	}
	level := make([]int, g.n)
	maxFlow := 0
	for g.bfs(s, t, level) {
		start := make([]int, g.n)
		for flow := g.dfs(s, t, math.MaxInt32, level, start); flow > 0; flow = g.dfs(s, t, math.MaxInt32, level, start) {
			maxFlow += flow
		}
	}
	return maxFlow
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Пример использования
	n := 6 // Количество вершин
	graph := NewGraph(n)

	// Добавляем рёбра (u, v, cap)
	graph.AddEdge(0, 1, 10)
	graph.AddEdge(0, 2, 10)
	graph.AddEdge(1, 2, 2)
	graph.AddEdge(1, 3, 4)
	graph.AddEdge(2, 4, 9)
	graph.AddEdge(1, 4, 8)
	graph.AddEdge(4, 3, 6)
	graph.AddEdge(4, 5, 10)
	graph.AddEdge(3, 5, 10)

	// Находим максимальный поток от источника (0) до стока (5)
	fmt.Println("Максимальный поток:", graph.MaxFlow(0, 5))
}
