package funcs

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func FileInfo(dir DIR, file File) DIR {

	Info, err := os.Stat(file.Path + "/" + file.Name)
	if err != nil {
		file.Err = err
		return dir
	}
	file.Name = Info.Name()

	file.Time = Info.ModTime()
	file.Mode = Info.Mode().String()
	file.Size = strconv.FormatInt(Info.Size(), 10)
	if len(file.Size) > dir.PInfo.MaxSize {
		dir.PInfo.MaxSize = len(file.Size)
	}
	stat, ok := Info.Sys().(*syscall.Stat_t)
	if !ok {
		file.Err = fmt.Errorf("failed to get syscall.Stat_t for file: %s", file.Name)
		return dir
	}

	if Flag_a || (!Flag_a && file.Name[0] != '.') {
		dir.Total += int(stat.Blocks)

	}

	// Get user and group names
	uid := strconv.Itoa(int(stat.Uid))
	gid := strconv.Itoa(int(stat.Gid))

	usr, err := user.LookupId(uid)
	if err != nil {
		file.Err = fmt.Errorf("failed to lookup user id: %s", uid)
		return dir
	}
	file.UserName = usr.Username
	if len(file.UserName) > dir.PInfo.MaxUName {
		dir.PInfo.MaxUName = len(file.UserName)
	}

	grp, err := user.LookupGroupId(gid)
	if err != nil {
		file.Err = fmt.Errorf("failed to lookup group id: %s", gid)
		return dir
	}
	file.GroupName = grp.Name
	if len(file.GroupName) > dir.PInfo.MaxGrName {
		dir.PInfo.MaxGrName = len(file.GroupName)
	}
	file.Hlink = strconv.Itoa(int(stat.Nlink))
	if dir.PInfo.MaxHlink < len(file.Hlink) {
		dir.PInfo.MaxHlink = len(file.Hlink)
	}
	dir.Files = append(dir.Files, file)
	return dir
}
