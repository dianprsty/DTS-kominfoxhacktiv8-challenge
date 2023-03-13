package main

import "fmt"

func main() {
	var kalimat = "selamat malam"
	var letterCount = map[string]int{}

	for _, v := range kalimat {
		value, exist := letterCount[string(v)]
		if exist {
			letterCount[string(v)] = value + 1
		} else {
			letterCount[string(v)] = 1
		}
		fmt.Println(string(v))
	}
	fmt.Printf("%+v\n", letterCount)
}