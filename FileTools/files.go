package FileTools

import (
	"github.com/fishman001/ExcelMerge/logger"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var logging = logger.GetStdLogger()

func WalkDir(dirPath, suffix string, deep bool) (files []string, err error) {
	files = make([]string, 0, 0)
	suffix = strings.ToUpper(suffix) // 忽略后缀匹配的大小写
	if deep {
		err = filepath.Walk(dirPath, func(filename string, file os.FileInfo, err error) error { // 遍历目录
			if err != nil { // 忽略错误
				logging.Warnln(err)
			}
			if file.IsDir() { // 忽略目录
				return nil
			}
			if strings.HasSuffix(strings.ToUpper(file.Name()), suffix) {
				files = append(files, filename)
			}
			return nil
		})
	} else {
		fileList, err := ioutil.ReadDir(dirPath)
		if err != nil {
			return nil, err
		}
		for _, file := range fileList {
			if file.IsDir() {
				continue
			}
			if strings.HasSuffix(strings.ToUpper(file.Name()), suffix) {
				files = append(files, path.Join(dirPath, file.Name()))
			}
		}
	}
	return files, err
}

func GetFilePathList(dirPath string, suffix string, isDeep bool) (files []string, err error) {
	absSourcePath, err := filepath.Abs(dirPath)
	if err != nil {
		return nil, err
	}

	dirs, err := WalkDir(absSourcePath, suffix, isDeep)
	if err != nil {
		return nil, err
	}
	return dirs, err
}
