package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct{
	absen int
	nama, alamat, pekerjaan, alasan string
}

func (s *student) introduce(){
	fmt.Printf("Nama\t\t: %s\nAlamat\t\t: %s\nPekerjaan\t: %s\nAlasan\t\t: %s\n", s.nama, s.alamat, s.pekerjaan, s.alasan)
}

func getStudent(students []student, args string){

	absen, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
	} else {
		found := false
		for _, student := range students {
			if student.absen == absen {
				student.introduce()
				found = true
				break
			}
		}
		if !found {
			fmt.Printf("Data siswa dengan absen %d tidak ditemukan\n", absen)
		} 
	}
}

func main(){
	studentsList := []student{
		{1, "Dian", "Jawa Tengah", "Programmer", "Mau mendalami golang"},
		{2, "Prasetyo", "DKI Jakarta", "Fullstack", "Beralih ke backend"},
	}

	args := os.Args

	getStudent(studentsList, args[1])

}