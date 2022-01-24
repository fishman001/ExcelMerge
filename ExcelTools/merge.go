package ExcelTools

import (
	"fmt"
	"github.com/fishman001/ExcelMerge/logger"
	"github.com/xuri/excelize/v2"
	"path/filepath"
	"time"
)

var logging = logger.GetStdLogger()

func MergeExcel(filePathList []string, sheetNameList []string) error {
	var fileList []*excelize.File
	logging.Infoln("开始读取文件...")
	for i, filePath := range filePathList {
		file, err := excelize.OpenFile(filePath)
		if err != nil {
			logging.Warnln(err)
			continue
		}
		fileList = append(fileList, file)
		logging.Debugf("已打开%v个文件: %s", i+1, filePath)
	}
	logging.Infof("共找到%v个文件，成功打开%v个文件\n", len(filePathList), len(fileList))

	var sheetRowsMap = make(map[string][]sheetRows)
	// for _, sheetName := range sheetNameList {
	// 	logging.Infof("开始合并%s表", sheetName)
	// 	for i, file := range fileList {
	// 		logging.Debugf("开始合并第%v个%s表: %s", i+1, sheetName, file.Path)
	// 		rows, err := file.GetRows(sheetName)
	// 		if err != nil {
	// 			logging.Warnf("无法读取文件%s的%s表\n", file.Path, sheetName)
	// 		}
	// 		sheetRowsMap[sheetName] = append(sheetRowsMap[sheetName], rows)
	// 	}
	// }
	logging.Infoln("开始解析文件...")
	for i, file := range fileList {
		logging.Debugf("解析第%v个文件：%s", i+1, file.Path)
		var searchSheetList []string
		if len(sheetNameList) == 0 {
			searchSheetList = file.GetSheetList()
		} else {
			searchSheetList = sheetNameList
		}
		for _, sheetName := range searchSheetList {
			rows, err := file.GetRows(sheetName)
			if err != nil {
				logging.Warnf("无法读取文件 [%s] 的 [%s] 表，可能不存在\n", file.Path, sheetName)
				logging.Debugln(err)
			}
			sheetRowsMap[sheetName] = append(sheetRowsMap[sheetName], rows)
		}
		err := file.Close()
		if err != nil {
			logging.Warnln(err)
		}
	}

	var mergeSheetData = make(map[string]sheetRows)
	for name, rowsList := range sheetRowsMap {
		logging.Infof("开始合并sheet：[%s]", name)
		mergeSheetData[name] = [][]string{}
		for _, rows := range rowsList {
			for _, row := range rows {
				mergeSheetData[name] = append(mergeSheetData[name], row)
			}
		}
	}
	// fmt.Println(mergeSheetData["商家货款"])
	newFile := excelize.NewFile()

	defer newFile.Close()

	newFile.DeleteSheet("Sheet1")
	for name, rows := range mergeSheetData {
		logging.Infof("开始写入sheet：%s", name)
		newFile.NewSheet(name)
		streamWriter, err := newFile.NewStreamWriter(name)
		if err != nil {
			return err
		}
		for i, row := range rows {
			cell, _ := excelize.CoordinatesToCellName(1, i+1)
			err := streamWriter.SetRow(cell, transInterfaceList(row))
			if err != nil {
				return err
			}
		}
		err = streamWriter.Flush()
		if err != nil {
			return err
		}
	}
	resultFileName := fmt.Sprintf("result_%s.xlsx", time.Now().Format("20060102150405"))
	if err := newFile.SaveAs(resultFileName); err != nil {
		return err
	}
	resultPaht, _ := filepath.Abs(resultFileName)
	logging.Infof("合并完成，结果文件存放位置：%s", resultPaht)
	return nil
}

type sheetRows [][]string

func transInterfaceList(row []string) []interface{} {
	var interfaceList []interface{}
	for _, data := range row {
		interfaceList = append(interfaceList, data)
	}
	return interfaceList
}
