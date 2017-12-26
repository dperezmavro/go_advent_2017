package main

type program struct {
	Name     string
	Weight   int
	Children []string
}

type programTree struct {
	Name        string
	Weight      int
	Children    []programTree
	TotalWeight int
}
