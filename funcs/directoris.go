package funcs

import (
	"os"
)

func FitchDir(dir DIR) DIR {
	files, err := os.ReadDir(dir.Path)
	if err != nil {
		dir.Err = err
		return dir
	}
	if Flag_a {
		var file File

		file.Path = dir.Path
		file.Name = "."
		dir = FileInfo(dir, file)

		file = File{}
		file.Path = dir.Path
		file.Name = ".."
		dir = FileInfo(dir, file)

	}
	for _, file := range files {
		if file.Name()[0] != '.' || (file.Name()[0] != '.' || Flag_a) {
			var File File
			File.Name = file.Name()
			File.Path = dir.Path
			dir = FileInfo(dir, File)
			
		}
		if file.IsDir() && Flag_R {
			var subDir DIR
			subDir.Path = dir.Path + "/" + file.Name()
			subDir.Name = file.Name()
			subDir = FitchDir(subDir)
			dir.SubDir = append(dir.SubDir, subDir)
		}
	}

	dir.Total /= 2

	return dir
}
