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

func UpdateTodo(id int, newName string, newStatus bool) error {
	file, err := os.Open(FileName)
	if err != nil {
		return fmt.Errorf("gagal membuka file : %d", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	record, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("gagal membaca csv : %d", err)
	}
	updated := false

	for i, rec := range record {
		if len(rec) < 3 {
			continue
		}
		recordId, err := strconv.Atoi(rec[0])
		if err != nil {
			continue
		}
		if recordId == id {
			record[i][1] = newName
			record[i][2] = strconv.FormatBool(newStatus)
			updated = true
			break
		}
	}

	if !updated {
		return fmt.Errorf("todo dengan ID %d tidak ditemukan", id)
	}

	fileWriter, err := os.Create(FileName)
	writer := csv.NewWriter(fileWriter)
	defer writer.Flush()
	if err != nil {
		return fmt.Errorf("gagal menulis data csv: %w", err)
	}
	defer fileWriter.Close()
	return nil
}
