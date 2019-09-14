package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

const baseDir = "C:\\Users\\as\\Desktop\\1"

func main() {
	rd, err := ioutil.ReadDir(baseDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	var files []int
	for _, fi := range rd {
		if fi.IsDir() || !strings.Contains(fi.Name(),".ts") {
			continue
		} else {
			temp := strings.TrimSuffix(fi.Name(), ".ts")
			num, err := strconv.Atoi(temp)
			if err != nil {
				fmt.Println(err)
				return
			}
			files = append(files, num)
		}
	}
	sort.Slice(files, func(i, j int) bool { return files[i] < files[j] })

	names := strconv.Itoa(files[0]) + ".ts"
	for _, num := range files[1:]{
		names += "+"
		names += strconv.Itoa(num) + ".ts"
	}

	names = "copy /b " + names + " new.ts"

	err = ioutil.WriteFile(baseDir+"\\1.bat", []byte(names), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
