//тут используется алгоритм Дейкстры, а граф задается json файлом

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
)

type Graph struct {
	Nodes map[string]map[string]float64 `json:"graph"`
}

func loadGraphFromFile(filePath string) (*Graph, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var graph Graph
	err = json.Unmarshal(data, &graph)
	if err != nil {
		return nil, err
	}

	return &graph, nil
}

func dijkstra(graph *Graph, startNode string) (map[string]float64, map[string]string) {
	distances := make(map[string]float64)
	previous := make(map[string]string)
	visited := make(map[string]bool)
	queue := make(map[string]float64)

	for node := range graph.Nodes {
		distances[node] = math.Inf(1)
		previous[node] = ""
		visited[node] = false
	}

	distances[startNode] = 0
	queue[startNode] = 0

	for len(queue) != 0 {
		var currentMinNode string
		currentMinDistance := math.Inf(1)

		for node, distance := range queue {
			if distance < currentMinDistance {
				currentMinNode = node
				currentMinDistance = distance
			}
		}

		delete(queue, currentMinNode)

		for neighbor, weight := range graph.Nodes[currentMinNode] {
			distance := distances[currentMinNode] + weight
			if distance < distances[neighbor] {
				distances[neighbor] = distance
				previous[neighbor] = currentMinNode
				queue[neighbor] = distance
			}
		}

		visited[currentMinNode] = true
	}

	return distances, previous
}

func printPath(previous map[string]string, startNode, endNode string) {
	path := []string{endNode}
	currentNode := endNode

	for previous[currentNode] != "" {
		currentNode = previous[currentNode]
		path = append([]string{currentNode}, path...)
	}

	fmt.Printf("Минимальный путь от %s до %s: %v\n", startNode, endNode, path)
}

func main() {
	graphFilePath := "graph.json"
	startNode := "A" //начальная вершина
	endNode := "D"   //конечная

	graph, err := loadGraphFromFile(graphFilePath)
	if err != nil {
		fmt.Println("Проблема подгрузки графа:", err)
		return
	}

	distances, previous := dijkstra(graph, startNode)

	fmt.Printf("Минимальное расстояние от точки %s:\n", startNode)
	for node, distance := range distances {
		fmt.Printf("%s: %.2f\n", node, distance)
	}

	printPath(previous, startNode, endNode)
}
