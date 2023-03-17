package main

import (
	"fmt"
	"sync"
)

type DataInterface interface {
	print(int, *sync.WaitGroup)
}

type DataStruct struct {
	val string
}

func (ds *DataStruct) print(data interface{}, urutan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ds.val = fmt.Sprintf("%v %d", data, urutan)
	fmt.Println(ds.val)
}

func main() {
	var data1 DataStruct
	var data2 DataStruct
	var coba interface{} = []string{"coba1", "coba2", "coba3"}
	var bisa interface{} = []string{"bisa1", "bisa2", "bisa3"}
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		func(urutan int) {
			go data1.print(coba, urutan, &wg)
		}(i)
	}
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		func(urutan int) {
			go data2.print(bisa, urutan, &wg)
		}(i)
	}
	wg.Wait()
}
