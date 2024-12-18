package funcs

import (
	"fmt"
	"strings"
)

func Print(result []DIR) {
	if len(result) == 0 {
		return
	}
	for i := 0; i < len(result); i++ {
		Print(result[i].SubDir)
		fmt.Println("")
		fmt.Println(result[i].ParentDir + "/" + result[i].Name + ":")
		fmt.Println("total:", result[i].Total)

		if result[i].Err != nil {
			fmt.Println("err:=", result[i].Err)
		}

		for j := 0; j < len(result[i].Files); j++ {
			if result[i].Files[j].Err != nil {
				fmt.Println("err               :", result[i].Files[j].Err)
			}
			fmt.Println(LFormat(result[i].Files[j], result[i].PInfo))
		}
	}
}
func LFormat(file File, prtInfo PrintInfo) string {
	arr := []string{file.Mode, padStart(file.Hlink, prtInfo.MaxHlink, " "), padStart(file.GroupName, prtInfo.MaxGrName, " "), padStart(file.UserName, prtInfo.MaxUName, " "), padStart(file.Size, prtInfo.MaxSize, " "), file.Name}
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
