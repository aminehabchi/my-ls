package funcs

func Resevre(files []File) []File {
	length := len(files)
	for j := 0; j < length/2; j++ {
		files[j], files[length-j-1] = files[length-1-j], files[j]
	}
	return files
}
