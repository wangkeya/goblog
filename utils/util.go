package utils

import (
	"os"
	"log"
	"strings"
	"time"
	"fmt"
	"path/filepath"
)


func CreateDir(Path string, IsSubDate bool) string {

	var folderPath string
	if IsSubDate {
		folderName := time.Now().Format("20060102")
		folderPath = filepath.Join(Path, folderName)
	}else{
		folderPath = Path
	}

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(folderPath, os.ModePerm) //0777也可以os.ModePerm
		os.Chmod(folderPath, os.ModePerm)
	}
	return folderPath
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}


func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}