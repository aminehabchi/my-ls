package funcs

import "time"

var (
	Flag_l bool
	Flag_R bool
	Flag_a bool
	Flag_r bool
	Flag_t bool
)

const blockSize = 512

type DIR struct {
	Err     error
	DirName string
	Total   int
	PInfo   PrintInfo
	SubDir  []DIR
	Files   []File
}

type File struct {
	Name      string
	Time      time.Time
	Mode      string
	UserName  string
	GroupName string
	Size      string
	Hlink     int
	Err       error
}

type PrintInfo struct {
	MaxSize   int
	MaxGrName int
	MaxUName  int
}
