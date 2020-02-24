package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Params struct {
	FileName string
	SkipCase bool
	DropUnUnique bool
	ReverseSort bool
	OutputFile string
	SortNums bool
	ColumnNum int
}

type TextToSort struct {
	StrArr []string
	columnToSort int
	SkipCase bool
}

var LessThan = func(first []string, second []string, columnToSort int) bool {
	return first[columnToSort-1] <  second[columnToSort-1]
}

func (t TextToSort) Len() int {
	return len(t.StrArr)
}

func (t TextToSort) Swap(i int, j int) {
	temp := t.StrArr[i]
	t.StrArr[i] = t.StrArr[j]
	t.StrArr[j] = temp
}

func (t TextToSort) Less(i int, j int) bool {
	var first, second []string
	if t.SkipCase {
		first = strings.Split(strings.ToLower(t.StrArr[i]), " ")
		second = strings.Split(strings.ToLower(t.StrArr[j]), " ")
	} else {
		first = strings.Split(t.StrArr[i], " ")
		second = strings.Split(t.StrArr[j], " ")
	}
	if t.columnToSort != 1 {
		if t.columnToSort >= len(first) || t.columnToSort >= len(second) || t.columnToSort < 1 {
			fmt.Println("bad column length")
			os.Exit(1)
		}
	}
	return LessThan(first, second, t.columnToSort)
}

func (t TextToSort) Unique(params Params) TextToSort {
	for i := 0; i < len(t.StrArr); i++ {
		j := i + 1
		for {
			if j < len(t.StrArr) {
				if params.SkipCase {
					if strings.ToLower(t.StrArr[i]) != strings.ToLower(t.StrArr[j]) {
						t.StrArr = append(t.StrArr[:i+1], t.StrArr[j:]...)
						break
					}
				} else {
					if t.StrArr[i] != t.StrArr[j] {
						t.StrArr = append(t.StrArr[:i+1], t.StrArr[j:]...)
						break
					}
				}
				j++
			} else {
				break
			}
		}
	}
	return t
}

func (t TextToSort) Printer() {
	for _, element := range t.StrArr {
		fmt.Println(element)
	}
}

func Sorter(text TextToSort, params Params) []string {
	if params.SkipCase {
		text.SkipCase = true
	}

	if params.ReverseSort {
		LessThan = func(first []string, second []string, columnToSort int) bool {
			return first[columnToSort-1] > second[columnToSort-1]
		}
	}

	sort.Sort(text)

	if params.DropUnUnique {
		text = text.Unique(params)
	}

	if params.OutputFile != "" {
		f, err := os.Create(params.OutputFile)
		if err != nil {
			fmt.Println(err)
			f.Close()
			os.Exit(1)
		}
		defer f.Close()
		for _, str := range text.StrArr {
			fmt.Fprintln(f, str)
		}
	} else {
		text.Printer()
	}
	return text.StrArr
}

func ParseArgs(args []string, params Params) Params {
	argsQuantity := len(args)
	if argsQuantity < 2 {
		fmt.Println("no file to sort")
		os.Exit(1)
	}
	params.FileName = args[argsQuantity-1]
	if params.FileName[0] == '-' {
		fmt.Println("filename must be in the end of query")
		os.Exit(1)
	}
	for i := 1; i < argsQuantity-1; i++ {
		switch args[i] {
		case "-f": // не учитываем регистр
			params.SkipCase = true
		case "-u": // выводим первое среди нескольких равных
			params.DropUnUnique = true
		case "-r": // сортируем по убыванию
			params.ReverseSort = true
		case "-k": // сортируем по столбцу
			params.ColumnNum, _ = strconv.Atoi(args[i+1])
		case "-o": // выводим в файл
			if argsQuantity < i+1 {
				os.Exit(1)
			}
			params.OutputFile = args[i+1]
		case "-n": // сортируем числа
			params.SortNums = true
		}
	}
	return params
}

func ParseFile(FileName string, arr []string) []string {
	data, _ := ioutil.ReadFile(FileName)
	lastInd := 0
	for i, elem := range data {
		if elem == 10 {
			arr = append(arr, string(data[lastInd:i]))
			lastInd = i + 1
		}
	}
	arr = append(arr, string(data[lastInd:]))
	return arr
}

func UniqueInts(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		j := i + 1
		for {
			if j < len(arr) {
				if arr[i] != arr[j] {
					arr = append(arr[:i+1], arr[j:]...)
					break
				}
				j++
			} else {
				break
			}
		}
	}
	return arr
}

func IntSorter(arr []string, params Params) []int {
	var result []int
	for _, elem := range arr {
		num, err := strconv.Atoi(elem)
		if err == nil {
			result = append(result, num)
		} else {
			continue
		}
	}
	sort.Ints(result)
	if params.DropUnUnique {
		result = UniqueInts(result)
	}
	if params.ReverseSort {
		lent := len(result)
		for i := 0; i < lent/2; i++ {
			result[i], result[lent-1-i] = result[lent-1-i], result[i]
		}
	}
	if params.OutputFile != "" {
		f, err := os.Create(params.OutputFile)
		if err != nil {
			fmt.Println(err)
			f.Close()
			os.Exit(1)
		}
		defer f.Close()
		for _, str := range result {
			fmt.Fprintln(f, str)
		}
	} else {
		for _, num := range result { fmt.Println(num) }
	}
	return result
}

func main() {
	params := Params {ColumnNum: 1}
	params = ParseArgs(os.Args, params)
	var arr []string
	arr = ParseFile(params.FileName, arr)

	if params.SortNums {
		IntSorter(arr, params)
	} else {
		text := TextToSort{arr[:], params.ColumnNum, false}
		Sorter(text, params)
	}
}
