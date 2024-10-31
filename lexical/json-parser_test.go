package lexical

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	jsonParser = JsonParser{}
)

func TestParsingEmptyJsonArraySuccessfully(t *testing.T) {
	//given
	token := []Token{
		{tokenType: LEFT_BRACKET, value: "["},
		{tokenType: RIGHT_BRACKET, value: "]"},
	}
	//when
	jsonValue, err := jsonParser.ParseTokens(token)

	//then
	assert.NoError(t, err)
	assert.NotNil(t, jsonValue)
	assert.Empty(t, jsonValue.getArrayValue())
	assert.Equal(t, ARRAY, jsonValue.getObjectType())
}

func TestParsingJsonArraySuccessfully(t *testing.T) {
	//given
	token := []Token{
		{tokenType: LEFT_BRACKET, value: "["},
		{tokenType: STRING, value: "test"},
		{tokenType: RIGHT_BRACKET, value: "]"},
	}
	//when
	jsonValue, err := jsonParser.ParseTokens(token)

	//then
	assert.NoError(t, err)
	assert.NotNil(t, jsonValue)
	assert.Equal(t, ARRAY, jsonValue.getObjectType())
	assert.NotEmpty(t, jsonValue.getArrayValue())
	assert.Equal(t, "test", jsonValue.getArrayValue()[0].getStringValue())
}

func TestParsingEmptyJsonObjectSuccessfully(t *testing.T) {
	//given
	token := []Token{
		{tokenType: LEFT_BRACE, value: "{"},
		{tokenType: RIGHT_BRACE, value: "}"},
	}
	//when
	jsonValue, err := jsonParser.ParseTokens(token)

	//then
	assert.NoError(t, err)
	assert.NotNil(t, jsonValue)
	assert.Equal(t, POJO, jsonValue.getObjectType())
	assert.Empty(t, jsonValue.getChildren())
}

func TestParsingJsonObjectSuccessfully(t *testing.T) {
	//given
	token := []Token{
		{tokenType: LEFT_BRACE, value: "{"},
		{tokenType: STRING, value: "name"},
		{tokenType: COLON},
		{tokenType: STRING, value: "test"},
		{tokenType: RIGHT_BRACE, value: "}"},
	}
	//when
	jsonValue, err := jsonParser.ParseTokens(token)

	//then
	assert.NoError(t, err)
	assert.NotNil(t, jsonValue)
	assert.Equal(t, POJO, jsonValue.getObjectType())
	assert.NotEmpty(t, jsonValue.getChildren())
	assert.Equal(t, 1, len(jsonValue.getChildren()))
	assert.Equal(t, "test", jsonValue.getChildren()["name"].getStringValue())
}
