package funcs

// func print() {

// 	for _, dir := range result {
// 		if len(result) > 1 {
// 			fmt.Println("\n" + dir.DirName + ":")
// 		}
// 		if Flag_t {
// 			sort.Slice(dir.Files, func(i, j int) bool {
// 				return dir.Files[i].Time.After(dir.Files[j].Time)
// 			})
// 		} else {
// 			sort.Slice(dir.Files, func(i, j int) bool {
// 				return strings.ToLower(dir.Files[i].Name) < strings.ToLower(dir.Files[j].Name)
// 			})
// 		}
// 		if Flag_r {
// 			resevre()
// 		}
// 		for i, file := range dir.Files {
// 			if Flag_l {
// 				if !Flag_a && file.Name[0] == '.' {
// 					continue
// 				}
// 				if i == 0 {
// 					fmt.Printf("total %v\n", Total)
// 				}
// 				// need to hundel spacses
// 				fmt.Printf("%v %v %v %v %*v %v %v", file.Mode, file.Hlink, file.UserName, file.GroupName, file.Size, file.Time.Format("Jan _2 15:04"), file.Name)
// 				if i != len(dir.Files)-1 {
// 					fmt.Println()
// 				}
// 			} else {
// 				if !Flag_a && file.Name[0] == '.' {
// 					continue
// 				}
// 				fmt.Print(file.Name + " ")
// 			}
// 		}
// 	}
// 	fmt.Println("")
// }

// func resevre() {
// 	for i := 0; i < len(result); i++ {
// 		length := len(result[i].Files)
// 		for j := 0; j < length/2; j++ {
// 			result[i].Files[j], result[i].Files[length-1-j] = result[i].Files[length-1-j], result[i].Files[j]
// 		}
// 	}
// }
