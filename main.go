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
	//resp, err := http.Get("https://dulich.tuoitre.vn/ngam-tuyet-trang-anh-dao-o-nhat-ban-thang-2-20190212121811143.htm")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//doc, err := goquery.NewDocumentFromReader(resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
}
