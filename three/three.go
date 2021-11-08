package main

import (
	"fmt"
	"strings"
)

type node struct {
	val   string
	left  *node
	right *node
}

/*
 * Given the root to a binary tree, implement serialize(root),
 * which serializes the tree into a string, and deserialize(s),
 * which deserializes the string back into the tree.
 */
func serialize(current *node) string {
	if current == nil {
		return ""
	}

	return fmt.Sprintf("%v|%v|%v", current.val, serialize(current.left), serialize(current.right))
}

func split(serialized string) (string, string) {
	parts := strings.SplitN(serialized, "|", 2)

	switch len(parts) {
	case 0:
		return "", ""
	case 1:
		return parts[0], ""
	default:
		return parts[0], parts[1]
	}
}

func deserialize(serialized string) (*node, string) {

	value, remainder := split(serialized)

	if value == "" {
		return nil, remainder
	}

	curr := &node{val: value}
	if len(remainder) > 0 {
		curr.left, remainder = deserialize(remainder)
	}

	if len(remainder) > 0 {
		curr.right, remainder = deserialize(remainder)
	}

	return curr, remainder
}

func main() {
	root := &node{
		val: "root",
		left: &node{
			val: "left",
			left: &node{
				val: "left.left",
			},
			right: &node{
				val: "right",
			},
		},
	}

	serialized := serialize(root)
	fmt.Println(serialized)
	deserialized, _ := deserialize(serialized)
	fmt.Println(serialize(deserialized))
}
