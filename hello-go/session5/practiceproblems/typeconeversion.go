package main

import (
	"fmt"
	"strconv"
)


func IntToString(i int) string {
	return strconv.Itoa(i)
}


func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func FloatToString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}


func StringToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}


func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

func StringToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

func main() {
	
	i := 42
	str := IntToString(i)
	fmt.Println("Int to String:", str)

	intVal, _ := StringToInt("123")
	fmt.Println("String to Int:", intVal)

	f := 3.14
	strFloat := FloatToString(f)
	fmt.Println("Float to String:", strFloat)

	floatVal, _ := StringToFloat("2.718")
	fmt.Println("String to Float:", floatVal)

	b := true
	strBool := BoolToString(b)
	fmt.Println("Bool to String:", strBool)

	boolVal, _ := StringToBool("true")
	fmt.Println("String to Bool:", boolVal)
}
