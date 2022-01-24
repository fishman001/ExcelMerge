package main

import (
	"github.com/fishman001/ExcelMerge/ExcelTools"
	"github.com/fishman001/ExcelMerge/FileTools"
	"github.com/fishman001/ExcelMerge/logger"
	"github.com/jessevdk/go-flags"
	"os"
)

var logging = logger.GetStdLogger()

type options struct {
	SourcePath   string   `short:"p" long:"path" description:"文件夹路径" default:"."`
	IsDeepSearch bool     `short:"d" long:"deep" description:"是否遍历文件夹中所有内容"`
	SheetNames   []string `short:"s" long:"sheets" description:"sheet名"`
}

func main() {
	var opts options
	_, err := flags.Parse(&opts)
	if err != nil {
		logging.Fatalln(err)
		os.Exit(-1)
	}

	filePathList, err := FileTools.GetFilePathList(opts.SourcePath, ".xlsx", opts.IsDeepSearch)
	if err != nil {
		logging.Fatalln(err)
	}
	err = ExcelTools.MergeExcel(filePathList, opts.SheetNames)
	if err != nil {
		logging.Error(err)
	}
}
