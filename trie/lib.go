package trie

import (
	"strings"
  . "github.com/bdreece/go-structs/queue"
)

type Node struct {
	Children 		[]Node
  IsTerminal 	bool
  Value				string
}

type Trie struct {
  Node
}

func NewTrie() Trie {
  return Trie{
    Node{
    	Children: nil,
  		IsTerminal: true,
    	Value: "",
    },
  }
}

func (t *Trie) AddWord(word string) {
	if len(word) <= 0 {
    return
  }

  chars := strings.Split(word, "")
  var current *Node = &t.Node

  for _, char := range chars {
    var nextNode *Node
    for _, child := range current.Children {
      if child.Value == char {
        nextNode = &child
        break
      }
    }

    if nextNode == nil {
      current.Children = append(current.Children, Node{
        Children: nil,
        IsTerminal: true,
        Value: char,
      })
    }
		current.IsTerminal = false
    current = nextNode
  }
}

func bfs(node *Node) []string {
	var values []string
  for child := range node.Children {
  	
  }
}

func (t *Trie) FindPrefix(prefix string) []string {
	
}
