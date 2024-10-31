package lexical

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	jsonLexicalParser = JsonLexicalParser{}
)

func TestParseEmptyJsonSuccessfully(t *testing.T) {
	tokens, err := jsonLexicalParser.getTokens([]rune("{}"))
	assert.Nil(t, err)
	assert.Equal(t, len(tokens), 2)
	assert.Equal(t, LEFT_BRACE, tokens[0].tokenType)
	assert.Equal(t, RIGHT_BRACE, tokens[1].tokenType)
}

func TestParseEmptyJsonArraySuccessfully(t *testing.T) {
	tokens, err := jsonLexicalParser.getTokens([]rune("[]"))
	assert.Nil(t, err)
	assert.Equal(t, len(tokens), 2)
	assert.Equal(t, LEFT_BRACKET, tokens[0].tokenType)
	assert.Equal(t, RIGHT_BRACKET, tokens[1].tokenType)
}

func TestParseWrongJsonFailed(t *testing.T) {
	tokens, err := jsonLexicalParser.getTokens([]rune(""))
	assert.Nil(t, err)
	assert.Empty(t, tokens)
}

func TestParseNoneNumericAsNumberJsonFailed(t *testing.T) {
	tokens, err := jsonLexicalParser.getTokens([]rune("[tttt]"))
	assert.NotNil(t, err)
	assert.Nil(t, tokens)
}

func TestParseJsonObjectSuccessfully(t *testing.T) {
	tokens, err := jsonLexicalParser.getTokens([]rune("{\"name\":\"my-name\"}"))
	assert.Nil(t, err)
	assert.NotNil(t, tokens)
	assert.Equal(t, LEFT_BRACE, tokens[0].tokenType)
	assert.IsTypef(t, STRING, tokens[1].tokenType, "Wrong token type")
	assert.Equal(t, "name", tokens[1].value)
	assert.Equal(t, COLON, tokens[2].tokenType)
	assert.IsTypef(t, STRING, tokens[3].tokenType, "Wrong token type")
	assert.IsTypef(t, "my-name", tokens[3].value, "Wrong token type")
	assert.Equal(t, RIGHT_BRACE, tokens[4].tokenType)
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
