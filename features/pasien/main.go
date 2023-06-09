package pasien

import (
	//
	"fmt"
	"projek/common"

	//
	pasienFunc "projek/features/pasien/functions"
	pasienMenu "projek/features/pasien/menu"
	pasienStruct "projek/features/pasien/structs"

	//

	post "projek/features/post"
	postStruct "projek/features/post/structs"
)

func Main(arrPasien *pasienStruct.TabPasien, arrPost *postStruct.TabPost) {
	var input int
	pasienMenu.ShowAuthPasienMenu()
	// Terima inputan dari user
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&input)

	for input != 3 {

		if input == 1 {
			// Register
			common.ResetConsole()
			pasienFunc.RegistrasiPasien(arrPasien)
			common.ResetConsole()
		} else if input == 2 {
			// Login
			if pasienFunc.LoginPasien(arrPasien) {
				// Jika login berhasil
				common.ResetConsole()

				post.Main(arrPost)
			} else {
				// Jika login gagal
				common.ResetConsole()

				fmt.Println("=======================================================================================")
				fmt.Println("                  Mohon inputkan username dan password dengan benar!                             ")
				fmt.Println("=======================================================================================")
				fmt.Println()

				common.ShowEndAction()
			}
		} else {
			fmt.Println("Menu Salah, coba lagi!")
		}
		pasienMenu.ShowAuthPasienMenu()
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&input)
	}

	// Reset console
	common.ResetConsole()
}
