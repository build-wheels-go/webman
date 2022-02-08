package framework

import (
	"errors"
	"strings"
)

// Tree 树结构
type Tree struct {
	root *node
}

// 节点
type node struct {
	isLast   bool
	segment  string
	handler  ControllerHandler
	children []*node
}

func newNode() *node {
	return &node{
		isLast:   false,
		segment:  "",
		children: []*node{},
	}
}

func NewTree() *Tree {
	root := newNode()
	return &Tree{root}
}

// 判断是否为通用节点(:开头)
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

func (n *node) filterChildNodes(segment string) []*node {
	if len(n.children) == 0 {
		return nil
	}
	//如果是通用节点，则所有子节点满足要求
	if isWildSegment(segment) {
		return n.children
	}
	nodes := make([]*node, 0, len(n.children))

	for _, cnode := range n.children {
		if isWildSegment(n.segment) {
			nodes = append(nodes, cnode)
		} else if cnode.segment == segment {
			nodes = append(nodes, cnode)
		}
	}
	return nodes
}

func (n *node) matchNode(uri string) *node {
	segments := strings.SplitN(uri, "/", 2)

	segment := segments[0]
	if !isWildSegment(segment) {
		segment = strings.ToUpper(segment)
	}

	cnodes := n.filterChildNodes(segment)
	if cnodes == nil || len(cnodes) == 0 {
		return nil
	}
	// 只有一个segment，则标记最后一个
	if len(segments) == 1 {
		for _, tn := range cnodes {
			if tn.isLast {
				return tn
			}
		}
		return nil
	}

	// 如果有2个segment，递归继续查找
	for _, tn := range cnodes {
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}

func (tree *Tree) AddRouter(uri string, handler ControllerHandler) error {
	n := tree.root
	//检测路由是否已存在
	if n.matchNode(uri) != nil {
		return errors.New("router exist: " + uri)
	}

	segments := strings.Split(uri, "/")
	for index, segment := range segments {
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *node

		cnodes := n.filterChildNodes(segment)
		if len(cnodes) > 0 {
			for _, cn := range cnodes {
				if cn.segment == segment {
					objNode = cn
					break
				}
			}
		}

		if objNode != nil {

			cnode := newNode()
			cnode.segment = segment
			if isLast {
				cnode.isLast = true
				cnode.handler = handler
			}
			n.children = append(n.children, cnode)
			objNode = cnode
		}

		n = objNode
	}
	return nil
}

func (tree *Tree) FindHandler(uri string) ControllerHandler {
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handler
}
