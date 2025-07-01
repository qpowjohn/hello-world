package main

import (
	"fmt"
	"os" // 導入 os 套件以讀取環境變數
)

func main() {
	// 嘗試從環境變數讀取 ENV
	env := os.Getenv("ENV")

	// 如果 ENV 變數沒有設定，給它一個預設值
	if env == "" {
		env = "development" // 例如，設定為 "development"
	}

	// 輸出 Hello, World! 和讀取到的 ENV
	fmt.Printf("Hello, World! Running in environment: %s\n", env)
}