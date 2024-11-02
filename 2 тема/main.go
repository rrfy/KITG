package main

import (
	"fmt"
	"math"
)

// Dijkstra принимает на вход матрицу смежности graph и начальную вершину start.
// Возвращает массив кратчайших расстояний от start до каждой вершины.
func Dijkstra(graph [][]int, start int) []int {
	numVertices := len(graph)
	dist := make([]int, numVertices)
	visited := make([]bool, numVertices)

	// Инициализируем все расстояния бесконечностью, а начальную вершину - нулём
	for i := range dist {
		dist[i] = math.MaxInt64 // "бесконечность" для отсутствующих путей
	}
	dist[start] = 0

	// Основной цикл алгоритма Дейкстры
	for i := 0; i < numVertices; i++ {
		u := -1

		// Найдём вершину u с минимальным расстоянием, которая ещё не посещена
		for v := 0; v < numVertices; v++ {
			if !visited[v] && (u == -1 || dist[v] < dist[u]) {
				u = v
			}
		}

		// Если не осталось непосещённых вершин, до которых можно добраться, выходим
		if dist[u] == math.MaxInt64 {
			break
		}

		// Помечаем вершину как посещённую
		visited[u] = true

		// Обновляем расстояния до смежных вершин
		for v := 0; v < numVertices; v++ {
			if graph[u][v] != -1 && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}

	return dist
}

func main() {
	// Пример графа в виде матрицы смежности
	// graph[i][j] представляет вес пути от вершины i до вершины j (-1 означает отсутствие пути)
	graph := [][]int{
		{0, 10, 3, -1, -1},
		{-1, 0, 1, 2, -1},
		{-1, 4, 0, 8, 2},
		{-1, -1, -1, 0, 7},
		{-1, -1, -1, 9, 0},
	}

	start := 0
	distances := Dijkstra(graph, start)

	// Вывод кратчайших расстояний от начальной вершины до каждой другой вершины
	fmt.Printf("Кратчайшие расстояния от вершины %d:\n", start)
	for i, d := range distances {
		if d == math.MaxInt64 {
			fmt.Printf("До вершины %d: нет пути\n", i)
		} else {
			fmt.Printf("До вершины %d: %d\n", i, d)
		}
	}
}
