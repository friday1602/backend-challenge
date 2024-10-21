package main

import (
	"strconv"
	"strings"
)

func decode(input string) string {
	// init result slice with 0 value
	result := make([]int, len(input)+1)

	// loop and check each case
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case 'L':
			for result[i] <= result[i+1] {
				result[i]++
				// ในเคส L จะมีปัญหาเพราะไม่ได้เปลี่ยนตัวขวา แต่เปลี่ยนตัวปัจจุบันทำให้กระทบตัวก่อนหน้า
				if i > 0 {
					index := i
					// important: เช็ค index ย้อนหลังทีละตัวว่าเป็น L หรือ = ที่ไม่เช็ค R เพราะยังไงขวาก้มากกว่าถ้าเพิ่มขวา
					for index > 0 {
						if input[index-1] == 'L' && result[index-1] <= result[index] {
							result[index-1]++
						} else if input[index-1] == '=' && result[index-1] != result[index] {
							result[index-1] = result[index]
						}
						index--
					}
				}

			}
		case 'R':
			for result[i] >= result[i+1] {
				result[i+1]++

			}
		case '=':
			if result[i] != result[i+1] {
				result[i+1] = result[i]
			}

		}
	}
	// create slice string result to store each num then join later
	resultString := make([]string, len(result))
	for i, num := range result {
		resultString[i] = strconv.Itoa(num)
	}
	return strings.Join(resultString, "")
}
