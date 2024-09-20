package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"syscall"
	"time"
)

func main() {

	getSubDir(".")

	return

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	sum := 0
	for i := 0; i < len(files); i++ {
		blocks := getblocks(files[i].Name())
		sum += blocks

	}
	sum /= 2
	fmt.Println("total ", sum)

	for i := 0; i < len(files); i++ {
		a, b, c, t, s, _ := getInfo(files[i].Name())
		fmt.Printf("%s %s %s %s %s %s\n", c, a, b, s, getTimeFormat(t), files[i].Name())
	}

}
func getSubDir(name string) {
	files, err := os.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		return
	}
	fmt.Printf("%s:\n", name)
	for i := 0; i < len(files); i++ {
		fmt.Print(files[i].Name(), " ")
	}
	fmt.Println("")
	fmt.Println("")

	for i := 0; i < len(files); i++ {
		if files[i].IsDir() {
			fmt.Printf("%s:", files[i].Name())
			getSubDir(files[i].Name())
		}
	}
	fmt.Println("")

}
func getInfo(filename string) (string, string, string, string, string, error) {
	fileInfo, err := os.Stat(filename) // Replace with your file path
	if err != nil {
		return "", "", "", "", "", err
	}
	// Get the file's system information
	stat := fileInfo.Sys().(*syscall.Stat_t)

	// Get the UID and GID
	uid := stat.Uid
	gid := stat.Gid

	userInfo, err := user.LookupId(fmt.Sprint(uid))
	if err != nil {
		return "", "", "", "", "", err
	}

	groupInfo, err := user.LookupGroupId(fmt.Sprint(gid))
	if err != nil {
		return "", "", "", "", "", err
	}

	size := strconv.Itoa(int(fileInfo.Size()))
	for len(size) < 4 {
		size = " " + size
	}
	return userInfo.Username, groupInfo.Name, fileInfo.Mode().String(), fileInfo.ModTime().String(), size, nil
}
func getblocks(namefile string) int {
	// Specify the path to the file

	// Get file information
	fileInfo, err := os.Stat(namefile)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	// Retrieve the system-specific file information
	stat, ok := fileInfo.Sys().(*syscall.Stat_t)
	if !ok {
		fmt.Println("Unable to retrieve system information.")
		return 0
	}

	// Get the number of blocks allocated to the file
	blocks := stat.Blocks
	return int(blocks)
}
func getTimeFormat(t string) string {
	//timeStr := "2024-09-19 23:40:57.665674959 +0100 +01"

	// Parse the time string into a time.Time object
	parsedTime, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", t)
	if err != nil {
		return ""
	}

	// Format the time to "Sep 19 23:40"
	formattedTime := parsedTime.Format("Jan 02 15:04")
	return formattedTime
}
