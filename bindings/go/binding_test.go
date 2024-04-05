package tree_sitter_terra_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-terra"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_terra.Language())
	if language == nil {
		t.Errorf("Error loading Terra grammar")
	}
}
