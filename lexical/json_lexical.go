package lexical

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	LEFT_BRACE    = "LEFT_BRACE"
	RIGHT_BRACE   = "RIGHT_BRACE"
	LEFT_BRACKET  = "LEFT_BRACKET"
	RIGHT_BRACKET = "RIGHT_BRACKET"
	COLON         = "COLON"
	COMMA         = "COMMA"
	STRING        = "STRING"
	NUMBER        = "NUMBER"
	TRUE          = "TRUE"
	FALSE         = "FALSE"
	NULL          = "NULL"
)

type (
	Token struct {
		tokenType string
		value     string
	}

	JsonLexicalParser struct {
	}
)

func (j *JsonLexicalParser) getTokens(s []rune) ([]Token, *ParseError) {
	currentPosition := 0
	tokens := []Token{}
	for currentPosition < len(s) {
		character := s[currentPosition]
		switch character {
		case '{':
			tokens = append(tokens, Token{tokenType: LEFT_BRACE, value: "{"})
			currentPosition++
		case '}':
			tokens = append(tokens, Token{tokenType: RIGHT_BRACE, value: "}"})
			currentPosition++
		case '[':
			tokens = append(tokens, Token{tokenType: LEFT_BRACKET, value: "["})
			currentPosition++
		case ']':
			tokens = append(tokens, Token{tokenType: RIGHT_BRACKET, value: "]"})
			currentPosition++
		case ':':
			tokens = append(tokens, Token{tokenType: COLON, value: ":"})
			currentPosition++
		case ',':
			tokens = append(tokens, Token{tokenType: COMMA, value: ","})
			currentPosition++
		case '"':
			stringSequenceBuilder := strings.Builder{}
			currentPosition++
			for currentPosition < len(s) && string(s[currentPosition]) != `"` {
				stringSequenceBuilder.WriteRune(s[currentPosition])
				currentPosition++
			}
			currentPosition++
			tokens = append(tokens, Token{tokenType: STRING, value: stringSequenceBuilder.String()})
		default:
			if unicode.IsSpace(character) {
				currentPosition++
			} else if unicode.IsNumber(character) {
				numberSequenceBuilder := strings.Builder{}
				for currentPosition < len(s) && (unicode.IsNumber(s[currentPosition]) || ',' == s[currentPosition] || '.' == s[currentPosition]) {
					numberSequenceBuilder.WriteRune(s[currentPosition])
					currentPosition++
				}
				tokens = append(tokens, Token{tokenType: NUMBER, value: numberSequenceBuilder.String()})
			} else if string(s[currentPosition:currentPosition+4]) == "true" {
				tokens = append(tokens, Token{tokenType: TRUE, value: "true"})
				currentPosition += 4
			} else if string(s[currentPosition:currentPosition+5]) == "false" {
				tokens = append(tokens, Token{tokenType: FALSE, value: "false"})
				currentPosition += 5
			} else if string(s[currentPosition:currentPosition+4]) == "null" {
				tokens = append(tokens, Token{tokenType: NULL, value: "null"})
				currentPosition += 4
			} else {
				return nil, &ParseError{
					err:     fmt.Errorf("unexpected charcater at position %d", currentPosition),
					message: fmt.Sprintf("Unexpected charcater at position %d", currentPosition),
				}
			}
		}
	}
	return tokens, nil
}
