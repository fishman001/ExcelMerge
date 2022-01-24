# ExcelMerge

A cli for Excel merging.  
It can merge Excels to one group by sheets.

## How to use

```shell
go get github.com/fishman001/ExcelMerge
#find the file in %GOPATH%/bin
```
```shell
ExcelMerge Options:
  /p, /path:    files directiory path (default: .)
  /d, /deep     whether search inner directiory files
  /s, /sheets:  sheet names(if null then merge all sheets)  -s a -s b -s c
```