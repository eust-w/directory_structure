package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//go:generate goversioninfo -icon=./icon.ico -manifest=resource/goversioninfo.exe.manifest
	var dirname string
	fmt.Print("you can enter dirname like this:\nAbsolute path: C:\\workspace\nRelative path: ./path\nEmpty path: No input will be generated in the running directory\nPlease enter your dirname: ")
	fmt.Scanln(&dirname)
	nowPath, _ := os.Getwd()
	if dirname == "" {
		dirname = nowPath
	}
	if dirname[0:1] == "." {
		dirname = nowPath + dirname[1:]
		fmt.Println(dirname)
	}
	var fileList []string
	listFiles(dirname, &fileList)
	fmt.Println(fileList)
	writeLines(fileList, dirname+"./overview.md", dirname)

}

func listFiles(dirname string, fileList *[]string) {
	fileInfos, _ := ioutil.ReadDir(dirname)
	for _, fi := range fileInfos {
		filename := dirname + "\\" + fi.Name()
		*fileList = append(*fileList, filename)
		if fi.IsDir() {

			listFiles(filename, fileList)
		}
	}
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
func writeLines(lines []string, path string, dirname string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	nilLen := len(dirname)
	w := bufio.NewWriter(file)
	fmt.Fprintln(w, "[toc]")
	for _, line := range lines {
		layel := strings.Count(line[nilLen+1:], "\\")
		isdir := IsDir(line)
		the_name := line[strings.LastIndex(line, "\\")+1:]
		if isdir {
			tem := strings.Repeat("#", layel+1) + " " + the_name
			fmt.Fprintln(w, tem)
		} else {
			tem := "[" + the_name + "]" + "(." + line[nilLen:] + ")"
			fmt.Fprintln(w, tem)
		}
	}
	return w.Flush()
}
