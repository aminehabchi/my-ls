package funcs

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func FileInfo(file File) (File, int) {

	var total int
	Info, err := os.Stat(file.Path + "/" + file.Name)
	if err != nil {
		//fmt.Println(file.Path + "/" + file.Name)
		file.Err = err
		return file, total
	}
	file.Name = Info.Name()
	file.Time = Info.ModTime()
	file.Mode = Info.Mode().String()
	file.Size = strconv.FormatInt(Info.Size(), 10)

	stat, ok := Info.Sys().(*syscall.Stat_t)
	if !ok {
		file.Err = fmt.Errorf("failed to get syscall.Stat_t for file: %s", file.Name)
		return file, total
	}

	if Flag_a || (!Flag_a && file.Name[0] != '.') {
		total = total + int(stat.Blocks)

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
