package main

import "fmt"

func main() {
	i := 0
	//var letter = []rune{'\u0421', ' ', '\u0410', ' ', '\u0428', ' ', '\u0410', ' ', '\u0420', ' ', '\u0412', ' ', '\u041e'}

	for i <= 5 {
		if i == 5 {
			for j:= 0; j <= 10; j++{
				if j == 5 {
					for i,v := range "САШАРВО" {
						fmt.Printf("character %#U starts at byte position %d\n", v, i)
					}
					continue
				}
				fmt.Println("Nilai j =", j)
			}
			break
		}

		fmt.Println("Nilai i =", i)
		i++
	}
}