package main

import (
	"os"
	"path/filepath"
	"strings"
)

func main() {

}

// GetFilesOfFolder 列出文件夹下所有文件，返回map
// length-层级，=1则只返回一层，=2则返回两层
// 支持文件后缀匹配
func GetFilesOfFolder(dir string) ([]string, error) {
	// dirPath, err := os.ReadDir(dir)
	// if err != nil {
	// 	return nil, err
	// }
	var files []string
	sep := string(os.PathSeparator)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			subFiles, err := GetFilesOfFolder(dir + sep + info.Name())
			if err != nil {
				return err
			}
			files = append(files, subFiles...)
		} else {
			// 过滤指定格式的文件
			ok := strings.HasSuffix(info.Name(), ".md")
			if ok {
				files = append(files, dir+sep+info.Name())
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	// for _, fi := range dirPath {
	//
	// }
	return files, nil
}
