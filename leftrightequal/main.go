package main

import (
	"fmt"
	"strings"
)

func main() {
	// รับข้อความเข้ารหัสจากคีย์บอร์ด
	var encoded string
	fmt.Print("Enter encoded string: ")
	fmt.Scanln(&encoded)
	encoded = strings.TrimSpace(encoded)

	fmt.Println(decode(encoded))
}
