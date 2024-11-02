package main

import (
	"container/list"
	"fmt"
)

// Graph представляет ориентированный граф с использованием списка смежности.
type Graph struct {
	vertices int           // количество вершин
	edges    map[int][]int // список смежности
}

// NewGraph создаёт новый граф с заданным количеством вершин.
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		edges:    make(map[int][]int),
	}
}

// AddEdge добавляет ориентированное ребро из вершины u в вершину v.
func (g *Graph) AddEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
}

// ShortestPath возвращает кратчайший путь от начальной вершины start до конечной вершины end.
func (g *Graph) ShortestPath(start, end int) []int {
	// Если начальная и конечная вершины совпадают, путь состоит только из одной вершины
	if start == end {
		return []int{start}
	}

	// Инициализация вспомогательных структур
	visited := make([]bool, g.vertices) // массив посещённых вершин
	prev := make([]int, g.vertices)     // массив для хранения предшественников
	for i := range prev {
		prev[i] = -1
	}

	// Очередь для выполнения BFS
	queue := list.New()
	queue.PushBack(start)
	visited[start] = true

	// Выполнение BFS
	for queue.Len() > 0 {
		node := queue.Remove(queue.Front()).(int)

		for _, neighbor := range g.edges[node] {
			if !visited[neighbor] {
				queue.PushBack(neighbor)
				visited[neighbor] = true
				prev[neighbor] = node

				// Если достигли конечной вершины, то можно прекратить поиск
				if neighbor == end {
					return reconstructPath(prev, start, end)
				}
			}
		}
	}

	// Путь не найден
	return nil
}

// reconstructPath восстанавливает путь по массиву предшественников.
func reconstructPath(prev []int, start, end int) []int {
	var path []int
	for at := end; at != -1; at = prev[at] {
		path = append([]int{at}, path...)
	}
	if path[0] == start {
		return path
	}
	return nil // путь не найден
}

func main() {
	// Пример использования: граф с корректными рёбрами для кратчайшего пути 0 → 1 → 5 → 6
	g := NewGraph(7)
	g.AddEdge(0, 1) // От перекрёстка 0 можно попасть на перекрёсток 1
	g.AddEdge(0, 2) // От перекрёстка 0 можно попасть на перекрёсток 2
	g.AddEdge(0, 3) // От перекрёстка 0 можно попасть на перекрёсток 3
	g.AddEdge(1, 5) // От перекрёстка 1 можно попасть на перекрёсток 5
	g.AddEdge(2, 4) // От перекрёстка 2 можно попасть на перекрёсток 4
	g.AddEdge(3, 4) // От перекрёстка 3 можно попасть на перекрёсток 4
	g.AddEdge(4, 6) // От перекрёстка 4 можно попасть на перекрёсток 6
	g.AddEdge(5, 6) // От перекрёстка 5 можно попасть на перекрёсток 6

	start, end := 0, 6
	path := g.ShortestPath(start, end)
	if path != nil {
		fmt.Printf("Кратчайший путь от перекрёстка %d до перекрёстка %d: %v\n", start, end, path)
	} else {
		fmt.Printf("Путь от перекрёстка %d до перекрёстка %d не найден.\n", start, end)
	}
}
