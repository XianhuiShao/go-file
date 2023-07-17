package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {

	type info struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	var inf info
	inf.ID = "HA"
	inf.Name = "订单"

	data, err := json.Marshal(inf)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	} else {
		fmt.Println(string(data))
	}

	file := "file.txt"
	//if err := os.WriteFile("file1.txt", []byte("Hello GOSAMPLES!"), 0666); err != nil {
	if err := os.WriteFile(file, []byte(data), 0666); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Write file Successful")

	err = os.Remove(file)
	if err != nil {

		//如果删除失败则输出 file remove Error!

		fmt.Println("file remove Error!")

		//输出错误详细信息

		fmt.Printf("%s", err)

	} else {

		//如果删除成功则输出 file remove OK!

		fmt.Print("file remove OK!")

	}
}
