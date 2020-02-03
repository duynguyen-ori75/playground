package btree

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func newLeafNode(keys []int, values []int, parent *InternalNode) *LeafNode {
	result := &LeafNode{keys: keys, values: values, parent: parent}
	return result
}

func newTestTree() *BPlusTree {
	root := &InternalNode{}
	root.keys = append(root.keys, 2, 5, 9)
	root.children = append(
		root.children,
		newLeafNode([]int{1, 2}, []int{6, 9}, root),
		newLeafNode([]int{3, 4, 5}, []int{10, 4, 7}, root),
		newLeafNode([]int{6, 9}, []int{12, 2}, root),
		newLeafNode([]int{14, 15}, []int{2, 8}, root),
	)
	root.children[0].(*LeafNode).rightSibling, root.children[1].(*LeafNode).rightSibling, root.children[2].(*LeafNode).rightSibling =
		root.children[1].(*LeafNode), root.children[2].(*LeafNode), root.children[3].(*LeafNode)
	return &BPlusTree{degree: 2, root: root}
}

func arrayToString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ",", -1), "[]")
}

func TestHelper(t *testing.T) {
	a := []int{}
	a = insertElement(a, 0, 5)
	a = insertElement(a, 1, 4)
	a = insertElement(a, 0, 2)
	a = insertElement(a, 2, 1)
	if !reflect.DeepEqual(a, []int{2, 5, 1, 4}) {
		t.Errorf("Expected array should be %s instead of %s", arrayToString(a), arrayToString([]int{2, 5, 1, 4}))
	}
}

func TestLeafNodeGetValue(t *testing.T) {
	node := newLeafNode([]int{2, 5, 7}, []int{1, 2}, nil)
	val, err := node.getValue(5)
	if err == nil {
		t.Error("Should raise invalid state exception")
	}
	node.values = append(node.values, 7)
	val, err = node.getValue(4)
	if err == nil || err.Error() != "Key 4 not found" {
		t.Error(fmt.Sprintf("Should raise key not found instead of: %s", err.Error()))
	}
	val, err = node.getValue(5)
	if err != nil || val != 2 {
		t.Error("Should get value correctly without any problem")
	}
}

func TestSearch(t *testing.T) {
	tree := newTestTree()
	node, err := tree.search(7)
	if err == nil || err.Error() != "Key 7 not found" {
		t.Error(fmt.Sprintf("Should raise key 7 not found instead of: %s", err.Error()))
	}
	node, err = tree.search(1)
	if err != nil || node.keys[0] != 1 {
		t.Error("Search does not return correct LeafNode")
	}
}

func TestLeafNodeInsert(t *testing.T) {

}