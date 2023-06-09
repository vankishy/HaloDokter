package main

import (
	// External Package
	fmt "fmt"

	// Internal Package
	common "projek/common"

	// Feature Home
	homeMenu "projek/features/home/menu"

	// Feature Post
	postStruct "projek/features/post/structs"

	// Feature Pasien
	pasien "projek/features/pasien"
	pasienStruct "projek/features/pasien/structs"
)

func main() {
	var arrPost postStruct.TabPost
	var arrPasien pasienStruct.TabPasien

	var input int
	homeMenu.ShowHomeMenu()
	// Terima inputan dari user
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&input)

	for input != 4 {
		// Cek apakah menu tersedia
		// Pasien
		if input == 1 {
			// Reset console
			common.ResetConsole()
			pasien.Main(&arrPasien, &arrPost)
		} else if input == 2 {
			// Dokter
			fmt.Println("Menu Dokter")
		} else if input == 3 {
			// Forum Konsultasi
			fmt.Println("Menu Konsultasi")
		} else {
			fmt.Println("Menu Salah, coba lagi!")
		}
		homeMenu.ShowHomeMenu()
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&input)
	}

	// Exit Aplikasi
	fmt.Println("Terima Kasih Telah Menggunakan Aplikasi Kami")
}
