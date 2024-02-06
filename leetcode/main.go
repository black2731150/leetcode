package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	err := renameDirectories("./")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func renameDirectories(parentDir string) error {
	err := filepath.Walk(parentDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理目录
		if info.IsDir() {
			if strings.Contains(info.Name(), "面试题") {
				name := []rune(info.Name())
				// 新目录名
				newName := "MST" + string(name[3:])

				// 构造新路径
				newPath := filepath.Join(filepath.Dir(path), newName)

				// 更改目录名
				err := os.Rename(path, newPath)
				if err != nil {
					return err
				}

				fmt.Printf("Renamed: %s -> %s\n", path, newPath)
				// 返回 SkipDir 以跳过已更名的目录
				return filepath.SkipDir
			}
		}

		return nil
	})

	return err
}
