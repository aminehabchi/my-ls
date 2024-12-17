package funcs

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func FileInfo(fileName string) (File, int) {
	var file File
	var total int
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		file.Err = err
		return file, total
	}
	file.Name = fileInfo.Name()
	file.Time = fileInfo.ModTime()
	file.Mode = fileInfo.Mode().String()
	file.Size = strconv.FormatInt(fileInfo.Size(), 10)
	total += int(fileInfo.Size())
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		file.Err = fmt.Errorf("failed to get syscall.Stat_t for file: %s", file.Name)
		return file, total
	}

	// Get user and group names
	uid := strconv.Itoa(int(stat.Uid))
	gid := strconv.Itoa(int(stat.Gid))

	usr, err := user.LookupId(uid)
	if err != nil {
		file.Err = fmt.Errorf("failed to lookup user id: %s", uid)
		return file, total
	}
	file.UserName = usr.Username

	grp, err := user.LookupGroupId(gid)
	if err != nil {
		file.Err = fmt.Errorf("failed to lookup group id: %s", gid)
		return file, total
	}
	file.GroupName = grp.Name
	file.Hlink = int(stat.Nlink)

	return file, total
}
