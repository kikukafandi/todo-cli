package todo

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	ID     int
	Name   string
	Status bool
}

var FileName string = "todo.csv"

func CekFileTodo() {
	_, err := os.Stat(FileName)

	if os.IsNotExist(err) {
		fmt.Printf("File %s tidak ditemukan. apakah kamu ingin membuat file baru? (y/n)\n", FileName)
		var response string
		fmt.Scanln(&response)

		if response == "y" || response == "Y" {
			file, err := os.Create(FileName)
			if err != nil {
				fmt.Println("Membuat file gagal:", err)
				return
			}
			defer file.Close()
			fmt.Println("File berhasil dibuat.")
		} else {
			fmt.Println("File tidak dibuat.")
			return
		}
	} else {
		fmt.Println("File ditemukan, lanjut proses")
	}

}
func GetLastId() int {
	file, err := os.Open(FileName)
	if err != nil {
		return 0
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return 0
	}
	if len(records) == 0 {
		return 0
	}

	lastRecord := records[len(records)-1]
	lastId, err := strconv.Atoi(lastRecord[0])
	if err != nil {
		return 0
	}
	return lastId
}
func AddTodo(name string) {
	file, err := os.OpenFile(FileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		return
	}
	defer file.Close()

	id := GetLastId() + 1

	writer := csv.NewWriter(file)
	defer writer.Flush()

	task := []string{
		strconv.Itoa(id),
		name,
		"false",
	}

	if err := writer.Write(task); err != nil {
		return
	}

	fmt.Println("Task berhasil ditambahkan")

}

func ShowTodo() {
	file, err := os.Open(FileName)
	if err != nil {
		fmt.Println("Error membuka file:", err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error membaca csv")
		return
	}
	if len(record) == 0 {
		fmt.Println("Belum ada todo")
	}
	fmt.Println("\n==== Daftar Todo List ====")
	for i, rec := range record {
		if len(rec) > 2 {

			fmt.Printf("%d. %s\n", i+1, rec[1])
		}
	}
}
