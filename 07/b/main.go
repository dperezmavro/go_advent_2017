package main

import (
	"log"
)

// var rootName = "tknk"

var rootName = "azqje"
var pTree programTree
var counts = make(map[string]int)
var nodes = make(map[string]program)

func (p *programTree) getWeight() int {
	total := 0
	for _, v := range p.Children {
		total += v.getWeight()
	}

	return total + p.Weight
}

func buildTree(startNode string) programTree {
	var result programTree

	result.Name = startNode
	result.Weight = nodes[startNode].Weight

	if len(nodes[startNode].Children) > 1 {
		for _, p := range nodes[startNode].Children {
			result.Children = append(
				result.Children,
				buildTree(p),
			)
		}

	}

	return result
}

func buildWeights(root *programTree) {
	root.TotalWeight = root.getWeight()
	for i := range root.Children {
		buildWeights(&root.Children[i])
	}
}

func isSubTreeBalanced(n programTree) bool {
	if len(n.Children) > 0 {
		w := n.Children[0].TotalWeight
		for i := 1; i < len(n.Children); i++ {
			if n.Children[i].TotalWeight != w {
				return false
			}
		}
	}

	return true
}

func getUnbalancedNode(children []programTree) programTree {
	var res programTree
	var targetWeight int
	var counts = make(map[int]int)

	for _, v := range children {
		counts[v.TotalWeight]++
	}

	for k, v := range counts {
		if v == 1 {
			targetWeight = k
		}
	}

	for _, v := range children {
		if v.TotalWeight == targetWeight {
			return v
		}
	}

	return res
}

func inspectNode(node programTree) {
	if isSubTreeBalanced(node) {
		log.Printf("Subtree for %s is balanced, need to decrease node's weight\n", node.Name)
	} else {
		log.Printf("Subtree for %s is not balanced\n", node.Name)
		problemNode := getUnbalancedNode(node.Children)
		log.Printf("Found problem Node %s\n", problemNode.Name)
		inspectNode(problemNode)
	}
}

func main() {

	input := readInput("input")
	processInput(input)

	tree := buildTree(rootName)
	buildWeights(&tree)

	inspectNode(tree)
}
