package lexical

import (
	"fmt"
	"slices"
	"strconv"
)

type (
	ObjectType string
	JsonObject interface {
		getObjectType() ObjectType
		getStringValue() string
		getIntegerValue() int
		getFloatValue() float64
		getBooleanValue() bool
		getArrayValue() []JsonObject
		getChildren() map[string]JsonObject
	}

	objectNode struct {
		nodeType ObjectType
		value    Token
		array    []JsonObject
		children map[string]JsonObject
	}

	JsonParser struct {
	}
	Stack[T any] struct {
		items []T
	}
)

const (
	POJO  ObjectType = "POJO"
	ARRAY ObjectType = "ARRAY"
	VALUE ObjectType = "VALUE"
)

var (
	LITERAL_TOKEN = []string{STRING, NUMBER, NULL, TRUE, FALSE}
)

func buildNode(tokenType string, value string) (JsonObject, error) {
	return &objectNode{
		nodeType: VALUE,
		value: Token{
			tokenType: tokenType,
			value:     value,
		},
	}, nil
}

func parseObject(tokens []Token, currentPosition *int) (JsonObject, error) {
	*currentPosition++
	jsonNode := &objectNode{
		nodeType: POJO,
		children: make(map[string]JsonObject),
	}

	for *currentPosition < len(tokens) {
		if tokens[*currentPosition].tokenType == STRING {
			key := tokens[*currentPosition].value
			*currentPosition++
			if tokens[*currentPosition].tokenType != COLON {
				return nil, fmt.Errorf("invalid JSON, expected a colon after key %s", key)
			}
			*currentPosition++
			if *currentPosition >= len(tokens) {
				return nil, fmt.Errorf("invalid JSON, expected a value token after colon")
			}
			value, err := parseValue(tokens, currentPosition)
			if err != nil {
				return nil, err
			}
			jsonNode.children[key] = value
			*currentPosition++
			if tokens[*currentPosition].tokenType == COMMA {
				*currentPosition++
			} else if tokens[*currentPosition].tokenType == RIGHT_BRACE {
				continue
			} else {
				return nil, fmt.Errorf("invalid JSON, expected a comma after key: value pair %s", key)
			}
		} else if tokens[*currentPosition].tokenType == RIGHT_BRACE {
			*currentPosition++
			return jsonNode, nil
		} else {
			return nil, fmt.Errorf("invalid JSON, expected } token")
		}
	}

	return jsonNode, nil
}

func parseArray(tokens []Token, currentPosition *int) (JsonObject, error) {
	*currentPosition++
	jsonNode := &objectNode{
		nodeType: ARRAY,
		array:    make([]JsonObject, 0),
	}

	for *currentPosition < len(tokens) {
		if tokens[*currentPosition].tokenType == RIGHT_BRACKET {
			*currentPosition++
			return jsonNode, nil
		} else if isLiteral(&tokens[*currentPosition]) {
			value, err := parseValue(tokens, currentPosition)
			if err != nil {
				return nil, err
			}
			jsonNode.array = append(jsonNode.array, value)
			*currentPosition++
			if tokens[*currentPosition].tokenType == COMMA {
				*currentPosition++
			} else if tokens[*currentPosition].tokenType == RIGHT_BRACKET {
				continue
			} else {
				return nil, fmt.Errorf("invalid JSON, expected a comma or ] token")
			}

		} else {
			return nil, fmt.Errorf("expecting `,` OR `]`token for Json array")
		}
	}
	return nil, fmt.Errorf("expecting `,` OR `]`token for Json array")
}

func parseValue(tokens []Token, currentPosition *int) (JsonObject, error) {
	currentTokenType := tokens[*currentPosition].tokenType
	switch currentTokenType {
	case STRING, NUMBER, TRUE, FALSE:
		node, err := buildNode(currentTokenType, tokens[*currentPosition].value)
		return node, err
	case LEFT_BRACE:
		return parseObject(tokens, currentPosition)
	default:
		return nil, fmt.Errorf("unexpected token type: %s", currentTokenType)
	}
}

func isLiteral(token *Token) bool {
	tokenType := token.tokenType
	return slices.Contains(LITERAL_TOKEN, tokenType)
}

func (jp *JsonParser) ParseTokens(tokens []Token) (JsonObject, error) {
	if len(tokens) < 2 {
		return nil, fmt.Errorf("insufficient tokens to parse")
	}
	i := 0
	if tokens[i].tokenType == LEFT_BRACE {
		return parseObject(tokens, &i)
	}

	if tokens[i].tokenType == LEFT_BRACKET {
		return parseArray(tokens, &i)
	}
	return nil, fmt.Errorf("unexpected token type: %v", i)
}

func (object *objectNode) getObjectType() ObjectType {
	return object.nodeType
}

func (object *objectNode) getStringValue() string {
	return object.value.value
}

func (object *objectNode) getIntegerValue() int {
	i, err := strconv.Atoi(object.value.value)
	if err != nil {
		panic(err)
	}
	return i
}

func (object *objectNode) getFloatValue() float64 {
	i, err := strconv.ParseFloat(object.value.value, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func (object *objectNode) getBooleanValue() bool {
	i, err := strconv.ParseBool(object.value.value)
	if err != nil {
		panic(err)
	}
	return i
}

func (object *objectNode) getArrayValue() []JsonObject {
	return object.array
}

func (object *objectNode) getChildren() map[string]JsonObject {
	return object.children
}

func (stack *Stack[T]) Push(item T) {
	stack.items = append(stack.items, item)
}

func (s *Stack[T]) Pop() {
	if s.IsEmpty() {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *Stack[T]) Top() (*T, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	return &s.items[len(s.items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack[T]) Print() {
	for _, item := range s.items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
