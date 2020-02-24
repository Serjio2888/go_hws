package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsingArguments(t *testing.T) {
	args := []string {"way", "filename"}
	params := Params{}
	paramsCheck := Params{
		FileName:     "filename",
	}
	assert.Equal(t, paramsCheck, ParseArgs(args, params))
	args = []string {"something", "-k", "4", "-o", "output.txt", "filename"}
	params = Params{}
	paramsCheck = Params{
		FileName:     "filename",
		OutputFile:   "output.txt",
		ColumnNum:    4,
	}
	assert.Equal(t, paramsCheck, ParseArgs(args, params))
	args = []string {"something", "-k", "334", "-o", "output.txt", "-u", "-r", "-n", "-f", "filename.txt"}
	params = Params{}
	paramsCheck = Params{
		FileName:     "filename.txt",
		SkipCase:     true,
		DropUnUnique: true,
		ReverseSort:  true,
		OutputFile:   "output.txt",
		SortNums:     true,
		ColumnNum:    334,
	}
	assert.Equal(t, paramsCheck, ParseArgs(args, params))
}

func TestIgnorRegister(t *testing.T) {
	args := []string {"something", "-f", "test_files/example.txt"}
	answer := []string {
		"AAA aaa ppp l",
		"AAA bbb ppp f",
		"aaa aaa ppp h",
		"aaa aaa ppp h",
		"aaa AAA ppp g",
		"BBB bbb ccc j",
		"CCC ccc ccc i",
		"ddd ddd ddd k",
		"DDD DDD DDD e",
	}
	params := Params{}
	params = ParseArgs(args, params)
	arr := []string{}
	arr = ParseFile(params.FileName, arr)
	text := TextToSort {
		StrArr:       arr[:],
		columnToSort: 1,
		SkipCase:     false,
	}
	assert.Equal(t, answer, Sorter(text, params))
}

func TestColumns(t *testing.T) {
	args := []string {"something", "-k", "2", "test_files/example.txt"}
	answer := []string{
		"aaa AAA ppp g",
		"DDD DDD DDD e",
		"aaa aaa ppp h",
		"aaa aaa ppp h",
		"AAA aaa ppp l",
		"AAA bbb ppp f",
		"BBB bbb ccc j",
		"CCC ccc ccc i",
		"ddd ddd ddd k",
	}
	params := Params{}
	params = ParseArgs(args, params)
	var arr []string
	arr = ParseFile(params.FileName, arr)
	text := TextToSort {
		StrArr:       arr[:],
		columnToSort: params.ColumnNum,
		SkipCase:     false,
	}
	assert.Equal(t, answer, Sorter(text, params))
}

func TestMainFunctions(t *testing.T) {
	args := []string {"something", "test_files/example.txt"}
	answer := []string{
		"AAA aaa ppp l",
		"AAA bbb ppp f",
		"BBB bbb ccc j",
		"CCC ccc ccc i",
		"DDD DDD DDD e",
		"aaa aaa ppp h",
		"aaa aaa ppp h",
		"aaa AAA ppp g",
		"ddd ddd ddd k",
	}
	params := Params{}
	params = ParseArgs(args, params)
	var arr []string
	arr = ParseFile(params.FileName, arr)
	text := TextToSort {
		StrArr:       arr[:],
		columnToSort: 1,
		SkipCase:     false,
	}
	assert.Equal(t, answer, Sorter(text, params))

	args = []string {"something", "-u", "test_files/example.txt"}
	answer = []string {
		"AAA aaa ppp l",
		"AAA bbb ppp f",
		"BBB bbb ccc j",
		"CCC ccc ccc i",
		"DDD DDD DDD e",
		"aaa aaa ppp h",
		"aaa AAA ppp g",
		"ddd ddd ddd k",
	}
	params = Params{}
	params = ParseArgs(args, params)
	arr = []string{}
	arr = ParseFile(params.FileName, arr)
	text = TextToSort {
		StrArr:       arr[:],
		columnToSort: 1,
		SkipCase:     false,
	}
	assert.Equal(t, answer, Sorter(text, params))

	args = []string {"something", "-r", "test_files/example.txt"}
	answer = []string {
		"ddd ddd ddd k",
		"aaa AAA ppp g",
		"aaa aaa ppp h",
		"aaa aaa ppp h",
		"DDD DDD DDD e",
		"CCC ccc ccc i",
		"BBB bbb ccc j",
		"AAA aaa ppp l",
		"AAA bbb ppp f",
	}
	params = Params{}
	params = ParseArgs(args, params)
	arr = []string{}
	arr = ParseFile(params.FileName, arr)
	text = TextToSort {
		StrArr:       arr[:],
		columnToSort: 1,
		SkipCase:     false,
	}
	assert.Equal(t, answer, Sorter(text, params))
}

func TestIntSorter(t *testing.T) {
	args := []string {"something", "-n", "test_files/numbers.txt"}
	answer := []int {
		-150, -150, 0, 0, 0, 11, 11, 12, 13, 14, 113,
	}
	params := Params{}
	params = ParseArgs(args, params)
	var arr []string
	arr = ParseFile(params.FileName, arr)
	assert.Equal(t, answer, IntSorter(arr, params))

	args = []string {"something", "-n", "-r", "test_files/numbers.txt"}
	answer = []int {
		113, 14, 13, 12, 11, 11, 0, 0, 0, -150, -150,
	}
	params = Params{}
	params = ParseArgs(args, params)
	arr = []string{}
	arr = ParseFile(params.FileName, arr)
	assert.Equal(t, answer, IntSorter(arr, params))

	args = []string {"something", "-n", "-r", "-u", "test_files/numbers.txt"}
	answer = []int {
		113, 14, 13, 12, 11, 0, -150,
	}
	params = Params{}
	params = ParseArgs(args, params)
	arr = []string{}
	arr = ParseFile(params.FileName, arr)
	assert.Equal(t, answer, IntSorter(arr, params))
}

func TestWorkWithFiles(t *testing.T) {
	assert.FileExists(t, "strings.txt")
	args := []string {"something", "-n", "-o", "answer.txt", "test_files/numbers.txt"}
	answer := []string {
		"-150", "-150", "0", "0", "0", "11", "11", "12", "13", "14", "113", "",
	}
	params := Params{}
	params = ParseArgs(args, params)
	var arr []string
	arr = ParseFile(params.FileName, arr)
	IntSorter(arr, params)
	assert.FileExists(t, "answer.txt")
	arr = []string {}
	answArr := ParseFile("answer.txt", arr)
	assert.Equal(t, answer, answArr)
}
