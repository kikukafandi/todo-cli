package todo

import (
	"fmt"
	"os"
)

type Task struct {
	ID     int
	Name   string
	Status bool
}

func CekFileTodo() {
	var fileName string = "todo.csv"
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		fmt.Printf("File %s tidak ditemukan. apakah kamu ingin membuat file baru? (y/n)", fileName)
		var response string
		fmt.Scanln(&response)

		if response == "y" || response == "Y" {
			file, err := os.Create(fileName)
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

func addTodo() {

}
