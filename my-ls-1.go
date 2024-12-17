package main

import (
	"fmt"
	"funcs/funcs"
	"os"
	"strings"
)

func main() {
	lsArgs := os.Args
	var result []funcs.DIR

	for i := 1; i < len(lsArgs); i++ {
		if lsArgs[i][0] == '-' && CheckFlagIsValid(lsArgs[i]) {
			flag := lsArgs[i]

			if strings.Contains(flag, "l") {
				funcs.Flag_l = true
			}
			if strings.Contains(flag, "R") {
				funcs.Flag_R = true
			}
			if strings.Contains(flag, "a") {
				funcs.Flag_a = true
			}
			if strings.Contains(flag, "r") {
				funcs.Flag_r = true
			}
			if strings.Contains(flag, "t") {
				funcs.Flag_t = true
			}
		} else if lsArgs[i][0] != '-' {
			dir := funcs.DIR{DirName: lsArgs[i]}
			result = append(result, dir)
		}
	}
	if len(result) == 0 {
		dir := funcs.DIR{DirName: "."}
		result = append(result, dir)
	}

	for i := 0; i < len(result); i++ {

		result[i] = funcs.FitchDir(result[i])
	}
	for i:=0;i<len(result);i++{
		for j:=0;j<len(result[i].Files);j++{
			fmt.Println(result[i].Files[j])
		}
	}
	fmt.Println(result[0].Total)
	
}

func CheckFlagIsValid(flag string) bool {
	l := len(flag)
	if l < 2 || l > 6 {
		return false
	}
	for i := 0; i < l; i++ {
		if i == 0 && flag[i] != '-' {
			return false
		} else if i != 0 {
			if !(flag[i] == 'a' || flag[i] == 'R' || flag[i] == 'r' || flag[i] == 't' || flag[i] == 'l') {
				return false
			}
		}
	}
	return true
}
