//Tree is a slice of pointers to Node objects
type Tree []*Node

//Node contains two pointers to children nodes
//(one or both may be nil), and a pointer to a parent node (which may be nil)
type Node struct {
	child1, child2 *Node
	parent         *Node
}

//Insert your SetParents() function here, along with any subroutines that you need.
func SetParents(t Tree) Tree {
	for _, node := range t {
		if node.child1 != nil {
			node.child1.parent = node
		}
		if node.child2 != nil {
			node.child2.parent = node
		}
	}
	return t

}