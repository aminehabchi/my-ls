package main

import (
	"fmt"
	"os"
	"os/user"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var (
	l          bool
	R          bool
	a          bool
	r          bool
	t          bool
	directorys []string
	result     = []DIR{}
)

type DIR struct {
	DirName string
	Files   []File
}
type File struct {
	Name      string
	Time      time.Time
	Mode      string
	UserName  string
	GroupName string
	Size      string
}

func main() {
	lsArgs := os.Args
	for i := 1; i < len(lsArgs); i++ {
		switch lsArgs[i] {
		case "-l":
			l = true
		case "-R":
			R = true
		case "-a":
			a = true
		case "-r":
			r = true
		case "-t":
			t = true
		default:
			directorys = append(directorys, lsArgs[i])
		}
	}
	if len(directorys) == 0 {
		directorys = append(directorys, ".")
	}

	for i := 0; i < len(directorys); i++ {
		run(directorys[i])
	}
	print()
}
func print() {
	for _, dir := range result {
		if t {
			sort.Slice(dir.Files, func(i, j int) bool {
				return dir.Files[i].Time.After(dir.Files[j].Time)
			})
		} else {
			sort.Slice(dir.Files, func(i, j int) bool {
				return strings.ToLower(dir.Files[i].Name) < strings.ToLower(dir.Files[j].Name)
			})
		}
		if r {
			resevre()
		}
		for _, file := range dir.Files {
			if l {
				fmt.Println(file)
			} else {
				fmt.Print(file.Name + " ")
			}
		}
	}
	fmt.Println("")
}
func resevre() {
	for i := 0; i < len(result); i++ {
		length := len(result[i].Files)
		for j := 0; j < length/2; j++ {
			result[i].Files[j], result[i].Files[length-1-j] = result[i].Files[length-1-j], result[i].Files[j]
		}
	}
}

func run(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
	}
	var Dir DIR
	Dir.DirName = dir
	for _, file := range files {
		if file.IsDir() && R {
			run(file.Name())
		}
		file, err := FileInfo(file.Name())
		if err != nil {
			fmt.Println(err)
			continue
		}
		Dir.Files = append(Dir.Files, file)
	}
	result = append(result, Dir)
}

func FileInfo(fileName string) (File, error) {
	var file File
	file.Name = fileName

	// Get file info
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		return File{}, err
	}

	file.Time = fileInfo.ModTime()
	file.Mode = fileInfo.Mode().String()
	file.Size = strconv.FormatInt(fileInfo.Size(), 10)

	// Get system-specific data
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		return File{}, fmt.Errorf("failed to get syscall.Stat_t for file: %s", fileName)
	}

	// Get user and group names
	uid := strconv.Itoa(int(stat.Uid))
	gid := strconv.Itoa(int(stat.Gid))

	usr, err := user.LookupId(uid)
	if err != nil {
		return File{}, fmt.Errorf("failed to lookup user id: %s", uid)
	}
	file.UserName = usr.Username

	grp, err := user.LookupGroupId(gid)
	if err != nil {
		return File{}, fmt.Errorf("failed to lookup group id: %s", gid)
	}
	file.GroupName = grp.Name

	return file, nil
}
