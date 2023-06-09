package post

import (
	"fmt"
	"projek/common"
	postStruct "projek/features/post/structs"
)

func ShowPost(tPost *postStruct.TabPost) {
	//Menampilkan isi postingan yang telah diinputkan pengguna
	common.ResetConsole()

	for i := 0; i < tPost.N; i++ {
		fmt.Printf("ID Postingan %s: %s \nTag : %s \n", tPost.ArrPost[i].ID, tPost.ArrPost[i].TxtAddPost, tPost.ArrPost[i].TagPost)
		for j := 0; j < tPost.ArrPost[i].Nreply; j++ {
			fmt.Println("Username : ", tPost.ArrPost[i].User.Username)
			fmt.Println("Balasan : ", tPost.ArrPost[i].ArrReply[j])
		}
		fmt.Println()

	}

}
