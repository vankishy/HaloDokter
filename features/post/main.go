package post

import (
	"fmt"
	"projek/common"
	pasienMenu "projek/features/pasien/menu"
	postFunc "projek/features/post/functions"
	postStruct "projek/features/post/structs"
)

func Main(arrPost *postStruct.TabPost) {
	var input int
	pasienMenu.ShowHomePasienMenu(arrPost)
	fmt.Print("Pilih Menu : ")
	fmt.Scan(&input)
	for input != 3 {
		postFunc.ShowPost(arrPost)
		if input == 1 {
			postFunc.PoChat(arrPost)
		} else if input == 2 {
			postFunc.RepPost(arrPost)
		} else if input == 3 {
		}
		common.ResetConsole()
		pasienMenu.ShowHomePasienMenu(arrPost)
		fmt.Print("Pilih Menu : ")
		fmt.Scan(&input)
	}
}
