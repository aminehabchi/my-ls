package funcs

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Print(result []DIR) {
	if len(result) == 0 {
		return
	}
	lenght := len(result)
	for i := 0; i < lenght; i++ {

		if lenght > 1 || Flag_R {
			fmt.Println("")
		}

		fmt.Println(result[i].Path + ":")
		fmt.Println("total:", result[i].Total)
		if result[i].Err != nil {
			fmt.Println("err:=", result[i].Err)
		}

		for j := 0; j < len(result[i].Files); j++ {
			if result[i].Files[j].Err != nil {
				fmt.Println("err               :", result[i].Files[j].Err)
				continue
			}
			fmt.Println(LFormat(result[i].Files[j], result[i].PInfo))
		}
		Print(result[i].SubDir)
	}
}
func LFormat(file File, prtInfo PrintInfo) string {
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
