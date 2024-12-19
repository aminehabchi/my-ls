package funcs

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var IsMoreThenOne bool = false

func Print(result []DIR) {
	if len(result) == 0 {
		return
	}
	lenght := len(result)
	for i := 0; i < lenght; i++ {
		if IsMoreThenOne {
			fmt.Println(result[i].Path + ":")
		}
		if Flag_l {
			fmt.Println("total:", result[i].Total)
		}

		if result[i].Err != nil {
			fmt.Println("err Dir:", result[i].Err)
			continue
		}

		for j := 0; j < len(result[i].Files); j++ {
			if result[i].Files[j].Err != nil {
				fmt.Println("err  File:", result[i].Files[j].Err)
				continue
			}
			
			fmt.Print(LFormat(result[i].Files[j], result[i].PInfo))
			
			if j == len(result[i].Files)-1 || Flag_l {
				fmt.Println()
			}
			if !Flag_l {
				fmt.Print(" ")
			}
		}
		if IsMoreThenOne {
			fmt.Println()
		}
		Print(result[i].SubDir)
	}
}
func LFormat(file File, prtInfo PrintInfo) string {
	if !Flag_l {
		return file.Name
	}
	var time string = timeFormat(file.Time)
	arr := []string{file.Mode, padStart(file.Hlink, prtInfo.MaxHlink, " "), padStart(file.GroupName, prtInfo.MaxGrName, " "), padStart(file.UserName, prtInfo.MaxUName, " "), padStart(file.Size, prtInfo.MaxSize, " "), time, file.Name}
	return strings.Join(arr, " ")

}
func padStart(input string, targetLength int, padChar string) string {
	if len(input) >= targetLength {
		return input
	}
	for len(input) < targetLength {
		input = padChar + input
	}
	return input
}
func timeFormat(t time.Time) string {
	Month := t.Month().String()[:3]
	Day := strconv.Itoa(t.Day())
	hour := strconv.Itoa(t.Hour())
	min := strconv.Itoa(t.Minute())

	return Month[:3] + " " + padStart(Day, 2, " ") + " " + padStart(hour, 2, "0") + ":" + padStart(min, 2, "0")
}

// func PrintWithColor(text, color string) string {
// 	var colorCode string
// 	switch color {
// 	case "red":
// 		colorCode = "\033[31m"
// 	case "green":
// 		colorCode = "\033[32m"
// 	case "yellow":
// 		colorCode = "\033[33m"
// 	case "blue":
// 		colorCode = "\033[34m"
// 	case "magenta":
// 		colorCode = "\033[35m"
// 	case "cyan":
// 		colorCode = "\033[36m"
// 	case "white":
// 		colorCode = "\033[37m"
// 	case "reset":
// 		colorCode = "\033[0m"
// 	default:
// 		colorCode = "\033[0m"
// 	}

// 	return colorCode + text + "\033[0m"
// }

// func resevre() {
// 	for i := 0; i < len(result); i++ {
// 		length := len(result[i].Files)
// 		for j := 0; j < length/2; j++ {
// 			result[i].Files[j], result[i].Files[length-1-j] = result[i].Files[length-1-j], result[i].Files[j]
// 		}
// 	}
// }

// if Flag_t {
// 	sort.Slice(dir.Files, func(i, j int) bool {
// 		return dir.Files[i].Time.After(dir.Files[j].Time)
// 	})
// } else {
// 	sort.Slice(dir.Files, func(i, j int) bool {
// 		return strings.ToLower(dir.Files[i].Name) < strings.ToLower(dir.Files[j].Name)
// 	})
// }
