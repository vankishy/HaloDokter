package post

import (
	"fmt"
	common "projek/common"
	postStruct "projek/features/post/structs"
)

func RepPost(tPost *postStruct.TabPost) {
	//Pengguna menginputkan balasan pada postingan
	var IDPost string
	var idx int
	fmt.Println("Pilih ID postingan yang ingin dibalas: ")
	fmt.Scan(&IDPost)
	idx = SearchPost(*tPost, IDPost)
	if idx != -1 {
		for 
		fmt.Println("Balasan anda: ")
		common.InputMultipleString(&tPost.ArrPost[idx].ArrReply[0])
	} else {
		fmt.Println("Postingan tidak ditemukan!")
		return
	}
	fmt.Println("Balasan anda: ")
	common.InputMultipleString(&tPost.ArrPost[idx].ArrReply[0])
	fmt.Println("Balasan anda telah diposting!")
}
