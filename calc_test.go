package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHardCore(t *testing.T) {
	query := "10*(2*((((1*2)*3)-((5-3/3)-44/2*10)+1)+23)/3)"
	assert.Equal(t, 1640, Evaluate([]byte(Tokenize(query))))
	query = "1440/2/2/(2*2*2)*(2*2*2)+1000000000000000+0-0+(0+1-1+0-0)"
	assert.Equal(t, 1000000000000360, Evaluate([]byte(Tokenize(query))))
}

func TestSimple(t *testing.T) {
	query_list := []string {"0", "2+2", "0-19", "(1+8)", "(14/2)"}
	answer_list := []int {0, 4, -19, 9, 7}
	for ind, elem := range query_list {
		assert.Equal(t, answer_list[ind], Evaluate([]byte(Tokenize(elem))))
	}
}

func TestOperations(t *testing.T) {
	query := "10+101000"
	assert.Equal(t, 101010, Evaluate([]byte(Tokenize(query))))
	query = "10-100"
	assert.Equal(t, -90, Evaluate([]byte(Tokenize(query))))
	query = "100/10"
	assert.Equal(t, 10, Evaluate([]byte(Tokenize(query))))
	query = "10*100"
	assert.Equal(t, 1000, Evaluate([]byte(Tokenize(query))))
}

func TestInvalidData(t *testing.T) {
	query := "10+10.1000"
	assert.Equal(t, false, Validation(query))
	query = "help me"
	assert.Equal(t, false, Validation(query))
	query = "2+2=5"
	assert.Equal(t, false, Validation(query))
	// wrong brackets
	query = "{2+2}"
	assert.Equal(t, false, Validation(query))
	query = "[5-2]"
	assert.Equal(t, false, Validation(query))
	// wrong number order
	query = "42435/-123"
	assert.Equal(t, false, Validation(query))
	query = "42435++123"
	assert.Equal(t, false, Validation(query))
	query = "-42435*123"
	assert.Equal(t, false, Validation(query))
}

func TestBracketsValidation(t *testing.T) {
	query := "(((10+5+14-19999)))"
	assert.Equal(t, true, CheckBrackets(query))
	query = "((((1+2)+3)-4)/2)+2*(15+(23-1))"
	assert.Equal(t, true, CheckBrackets(query))
	query = ""
	assert.Equal(t, true, CheckBrackets(query))
	query = "((33-2)/2)-2)"
	assert.Equal(t, false, CheckBrackets(query))
	query = "(5+((11-0)/2)"
	assert.Equal(t, false, CheckBrackets(query))
	query = "(((14-88)/2)/3)/3)/4/1)+1)/2)/1)/1)/6)/1)/4))))"
	assert.Equal(t, false, CheckBrackets(query))
	query = "((()))((()))((())(())))"
	assert.Equal(t, false, CheckBrackets(query))
	query = "("
	assert.Equal(t, false, CheckBrackets(query))
	query = ")"
	assert.Equal(t, false, CheckBrackets(query))
	query = "(1+1"
	assert.Equal(t, false, CheckBrackets(query))
}

