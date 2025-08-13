package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kikukafandi/todo-cli/todo"
)

func main() {
	todo.CekFileTodo()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n==========================")
		fmt.Println("      TODOLIST APP      ")
		fmt.Println("==========================")
		fmt.Println("1. Tambahkan ToDo")
		fmt.Println("2. Update ToDo")
		fmt.Println("3. Lihat ToDo")
		fmt.Println("4. Hapus ToDo")
		fmt.Println("5. Keluar Program")
		fmt.Println("--------------------------")
		fmt.Print("Pilih menu (1-5): ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Println("\n--- Tambah Todo ---")
			fmt.Print("Masukkan Todo: ")
			scanner.Scan()
			item := strings.TrimSpace(scanner.Text())
			todo.AddTodo(item)
		case "2":
			fmt.Println("\n--- Update Todo ---")
			fmt.Printf("Masukkan ID todo yang udah diupdate (0 untuk batal): ")
			scanner.Scan()
			idStr := strings.TrimSpace(scanner.Text())

			if idStr == "0" {
				fmt.Println("Update dibatalkan")
				break
			}

			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Id Harus berupa angka.")
				continue
			}

			fmt.Println("Masukkan nama todo baru (0 untuk batal): ")
			scanner.Scan()
			newName := strings.TrimSpace(scanner.Text())
			if newName == "0" {
				fmt.Println("Update dibatalkan")
				break
			}

			fmt.Println("Apakah status Todo sudah selesai? (y/n 0 untuk membatalkan): ")
			scanner.Scan()
			statusStr := strings.TrimSpace(scanner.Text())
			if statusStr == "0" {
				fmt.Println("Update dibatalkan")
				break
			}

			var newStatus bool
			if strings.ToLower(statusStr) == "y" {

				newStatus = true
			} else {
				newStatus = false
			}
			err = todo.UpdateTodo(id, newName, newStatus)
			if err != nil {

				fmt.Println("Error update Todo: ", err)
			} else {
				fmt.Println("Todo berhasil diupdate!")
			}

		case "3":
			for {
				fmt.Println("\n === Daftar Todo ===")
				todo.ShowTodo()
				fmt.Println("0. untuk kembali ke menu utama")
				fmt.Println("Pilih :")

				scanner.Scan()
				backChoice := strings.TrimSpace(scanner.Text())

				if backChoice == "0" {
					break
				} else {
					fmt.Println("Pilihan tidak valid")

				}
			}
		case "4":
			fmt.Println("\n--- Hapus Todo ---")
			fmt.Print("Masukkan ID todo yang akan dihapus (0 untuk batal): ")
			scanner.Scan()
			idStr := strings.TrimSpace(scanner.Text())

			if idStr == "0" {
				fmt.Println("Hapus dibatalkan.")
				break
			}

			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("ID harus berupa angka.")
				break
			}

			fmt.Printf("Yakin ingin menghapus Todo ID %d? (y/n): ", id)
			scanner.Scan()
			confirm := strings.ToLower(strings.TrimSpace(scanner.Text()))

			if confirm != "y" {
				fmt.Println("Hapus dibatalkan.")
				break
			}

			todo.DeleteTodo(id)
			fmt.Println("Todo berhasil dihapus.")

		case "5":
			fmt.Println("Keluar dari program...")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
