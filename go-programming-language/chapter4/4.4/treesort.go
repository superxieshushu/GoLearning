package main

type Tree struct {
	value       int
	left, right *Tree
}

func Sort(values []int) {
	var root *Tree
	for _, v := range values {
		root = Add(root, v)
	}
	AppendValues(values[:0], root)
}

func AppendValues(values []int, t *Tree) []int {
	if t != nil {
		values = AppendValues(values, t.left)
		values = append(values, t.value)
		values = AppendValues(values, t.right)
	}
	return values
}

func Add(t *Tree, value int) *Tree {
	if t == nil {
		t = new(Tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = Add(t.left, value)
	} else {
		t.right = Add(t.right, value)
	}
	return t
}
