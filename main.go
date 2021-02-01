package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cch12313/synthesizing-file/helper"
)

func main() {
	var sourcePath string
	fmt.Print("輸入檔案來源資料夾：")
	fmt.Scanf("%s", &sourcePath)
	if !helper.ValidateDir(sourcePath) {
		panic("檔案來源不存在")
	}

	var targetPath string
	fmt.Print("產出檔案路徑：")
	fmt.Scanf("%s", &targetPath)
	if !helper.ValidateDir(sourcePath) {
		panic("產出檔案路徑不存在")
	}

	synthesizingFile(sourcePath)

	fmt.Println("結束")
}

func synthesizingFile(sourcePath string) {
	file, _ := ioutil.ReadFile(sourcePath + "test.xlsx")
	fmt.Printf("%v", &file)
}

func openFile(sourcePath string) {
	file, err := os.Open(sourcePath)

	if err != nil {
		panic(err)
	}
	defer file.Close()
}
