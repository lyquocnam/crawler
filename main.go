package main

import (
	"crawler/models"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Vui lòng nhập địa chỉ bài viết từ báo Tuổi Trẻ")
	} else {
		url := os.Args[1]
		if len(url) < 2 {
			fmt.Println("Vui lòng nhập địa chỉ bài viết từ báo Tuổi Trẻ")
			return
		}

		doc := models.TuoiTreDocument{}
		err := doc.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = doc.SaveCSV()
		if err != nil {
			fmt.Println(err)
			return
		}

		doc.Print()
	}
}
