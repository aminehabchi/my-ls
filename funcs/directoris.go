package funcs

import (
	"fmt"
	"os"
)

func FitchDir(dir DIR) DIR {
	fmt.Println(dir.Path)
	files, err := os.ReadDir(dir.Path)
	fmt.Println(len(files))
	if err != nil {
		dir.Err = err
		return dir
	}

	for _, file := range files {
		//fmt.Println("aa", file.Name())
		if file.IsDir() && Flag_R {
			var subDir DIR
			subDir.Path = dir.Path + "/" + file.Name()
			subDir.Name = file.Name()
			subDir = FitchDir(subDir)
			dir.SubDir = append(dir.SubDir, subDir)
		}
		if file.Name()[0] != '.' || (file.Name()[0] != '.' || Flag_a) {
			var File File
			File.Name = file.Name()
			File.Path = dir.Path
			dir = FileInfo(dir, File)

		}
	}
	if Flag_a {
		var file File

		file.Path = dir.Path
		dir = FileInfo(dir, file)

		file = File{}
		file.Name = dir.ParentDir
		dir = FileInfo(dir, file)

	}
	dir.Total /= 2

	return dir
}
