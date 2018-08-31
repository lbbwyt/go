// tire_test
package logic

import (
	"fmt"
	"school_suggestion/logic"
	"testing"
)

func TestPrefixSearch(t *testing.T) {
	var trieTree *logic.Trie
	trieTree = logic.NewTrie()
	trieTree.Add("abcdefg", "abcdefg")
	trieTree.Add("abcde", "abcde")
	trieTree.Add("abc", "abc")
	trieTree.Add("abcgf", "abcgf")
	var nodes []*logic.Node
	nodes = trieTree.PrefixSearch("abc", 10)
	fmt.Println(len(nodes))
	for _, node := range nodes {
		fmt.Println(node.Data)
	}
}
