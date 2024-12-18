package funcs

import (
	"fmt"
	"math"
	"os"
)

func FitchDir(dir DIR) DIR {
	files, err := os.ReadDir(dir.DirName)
	if err != nil {
		dir.Err = err
		return dir
	}

	for _, file := range files {
		fmt.Println(file.Name())
		if file.IsDir() && Flag_R {
			var subDir DIR
			subDir.DirName = file.Name()
			FitchDir(subDir)
			dir.SubDir = append(dir.SubDir, subDir)
		}
		if file.Name()[0] != '.' || (file.Name()[0] != '.' || Flag_a) {
			File, total := FileInfo(file.Name())
			dir.Total += (int(math.Ceil(float64(total) / blockSize)))
			dir.Files = append(dir.Files, File)
		}
	}
	if Flag_a  {
		File, total := FileInfo(dir.DirName)
		dir.Total += (int(math.Ceil(float64(total) / blockSize)))
		dir.Files = append(dir.Files, File)

		File, total = FileInfo("."+dir.DirName)
		dir.Total += (int(math.Ceil(float64(total) / blockSize)))
		dir.Files = append(dir.Files, File)
	}

	return dir
}
