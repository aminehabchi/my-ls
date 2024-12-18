package funcs

import "fmt"

func Print(result []DIR) {
	if len(result) == 0 {
		return
	}
	for i := 0; i < len(result); i++ {
		Print(result[i].SubDir)
		fmt.Println("")
		fmt.Println(result[i].ParentDir + "/"+result[i].Name + ":")
		for j := 0; j < len(result[i].Files); j++ {
			fmt.Println(result[i].Files[j])
		}
	}
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
