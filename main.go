package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/fishman001/ExcelMerge/ExcelTools"
	"github.com/fishman001/ExcelMerge/FileTools"
	"github.com/fishman001/ExcelMerge/logger"
	"github.com/flopp/go-findfont"
	"os"
	"strings"
)

var logging = logger.GetStdLogger()

type options struct {
	SourcePath   string   `short:"p" long:"path" description:"文件夹路径" default:"."`
	IsDeepSearch bool     `short:"d" long:"deep" description:"是否遍历文件夹中所有内容"`
	SheetNames   []string `short:"s" long:"sheets" description:"sheet名"`
}

var msg string
var help string

func main() {
	help = `Path: 文件夹路径
Sheets: 要合并的sheet，多个sheet间用英文逗号(,)隔开，留空则合并所有sheet
Deep: 是否查找子文件夹中文件`
	var opts options
	// _, err := flags.Parse(&opts)
	// if err != nil {
	// 	logging.Fatalln(err)
	// 	os.Exit(-1)
	// }
	myApp := app.New()
	myWindow := myApp.NewWindow("ExcelMerge")
	var isDeep bool
	isDeepCheck := widget.NewCheckWithData("", binding.BindBool(&isDeep))
	var sheetsStr string
	sheetsEntry := widget.NewEntryWithData(binding.BindString(&sheetsStr))
	var sourcePath string
	sourcePath = "."
	pathEntry := widget.NewEntryWithData(binding.BindString(&sourcePath))
	msgLabel := widget.NewLabel(msg)
	helpLabel := widget.NewLabel(help)
	form := widget.NewForm(
		widget.NewFormItem("Help:", helpLabel),
		widget.NewFormItem("Path:", pathEntry),
		widget.NewFormItem("Sheets:", sheetsEntry),
		widget.NewFormItem("Deep:", isDeepCheck),
		widget.NewFormItem("Start:", widget.NewButton("Merge", func() {
			widget.NewFormItem("Help", widget.NewLabel("Sheet用英文逗号(,)隔开，留空则合并所有sheet"))
			opts.IsDeepSearch = isDeep
			opts.SourcePath = sourcePath
			opts.SheetNames = strings.Split(sheetsStr, ",")
			if len(opts.SheetNames) == 1 && opts.SheetNames[0] == "" {
				opts.SheetNames = opts.SheetNames[1:]
			}
			msg = fmt.Sprintf("input:\nPath:%s\nSheets:%s\nDeep:%v\nStart merging...", opts.SourcePath, strings.Join(opts.SheetNames, " | "), opts.IsDeepSearch)
			msgLabel.SetText(msg)
			err := merge(opts)
			if err != nil {
				logging.Errorln(err)
				msg = fmt.Sprintf("err: %s", err)
				msgLabel.SetText(msg)
				return
			}
			msgLabel.SetText("Merge Complete!")
		})),
		widget.NewFormItem("", msgLabel),
	)

	myWindow.Resize(fyne.Size{
		Width:  800,
		Height: 420,
	})
	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}

func merge(opts options) error {
	filePathList, err := FileTools.GetFilePathList(opts.SourcePath, ".xlsx", opts.IsDeepSearch)
	if err != nil {
		return err
	}
	err = ExcelTools.MergeExcel(filePathList, opts.SheetNames)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		// 楷体:simkai.ttf
		// 黑体:simhei.ttf
		if strings.Contains(path, "simkai.ttf") {
			// fmt.Println(path)
			os.Setenv("FYNE_FONT", path) // 设置环境变量  // 取消环境变量 os.Unsetenv("FYNE_FONT")
			break
		}
	}
}
