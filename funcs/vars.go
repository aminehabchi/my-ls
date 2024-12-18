package funcs

import "time"

var (
	Flag_l bool
	Flag_R bool
	Flag_a bool
	Flag_r bool
	Flag_t bool
)

type DIR struct {
	Name   string
	Path      string
	ParentDir string
	Total     int
	PInfo     PrintInfo
	SubDir    []DIR
	Files     []File
	Err       error
}

type File struct {
	Name      string
	Path      string
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
