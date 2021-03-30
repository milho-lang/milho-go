package parser

import (
	"fmt"

	"github.com/danfragoso/milho/tokenizer"
)

func Parse(tokens []*tokenizer.Token) (*Node, error) {
	tokenList := CreateTokenList(tokens)
	currentToken := tokenList.FirstToken()

	var currentNode *Node

	for currentToken != nil {
		switch currentToken.Type {
		case tokenizer.OParen:
			if currentNode == nil {
				currentNode = createEmptyNode()
			} else {
				childNode := createEmptyNode()
				childNode.Parent = currentNode

				currentNode.Nodes = append(currentNode.Nodes, childNode)
				currentNode = childNode
			}

		case tokenizer.OBrack:
			if currentNode == nil {
				currentNode = createListNode()
			} else {
				childNode := createListNode()
				childNode.Parent = currentNode

				currentNode.Nodes = append(currentNode.Nodes, childNode)
				currentNode = childNode
			}

		case tokenizer.CParen:
			if currentNode == nil {
				return nil, fmt.Errorf("unexpected token '('")
			} else if currentNode.Parent != nil {
				currentNode = currentNode.Parent
			}

		case tokenizer.CBrack:
			if currentNode == nil {
				return nil, fmt.Errorf("unexpected token '['")
			} else if currentNode.Parent != nil {
				currentNode = currentNode.Parent
			}

		case tokenizer.Symbol:
			if currentNode == nil {
				return nil, fmt.Errorf("unexpected token '%s'", currentToken.Value)
			} else {
				switch currentNode.Type {
				case Nil:
					if isMacro(currentToken.Value) {
						currentNode.Type = Macro
					} else {
						currentNode.Type = Function
					}

					currentNode.Identifier = currentToken.Value

				default:
					childNode := createEmptyNode()
					childNode.Parent = currentNode
					childNode.Type = Identifier
					childNode.Identifier = currentToken.Value

					if currentNode == nil {
						currentNode = childNode
					} else {
						currentNode.Nodes = append(currentNode.Nodes, childNode)
					}
				}

			}

		case tokenizer.Number:
			childNode := createEmptyNode()
			childNode.Parent = currentNode
			childNode.Type = Number
			childNode.Identifier = currentToken.Value

			if currentNode == nil {
				currentNode = childNode
			} else {
				currentNode.Nodes = append(currentNode.Nodes, childNode)
			}
		}

		currentToken = tokenList.NextToken()
	}

	return currentNode, nil
}

func createEmptyNode() *Node {
	return &Node{}
}

func createListNode() *Node {
	return &Node{Type: List}
}

func isMacro(macroCandidate string) bool {
	return macroCandidate == "defn"
}
