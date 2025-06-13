package core

import (
	"fmt"
	"github.com/redmask-hb/GoSimplePrint/goPrint"
	"kenaito-pdf2png/util"
	"kenaito-pdf2png/vars"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type FileScanner struct {
}

func (FileScanner) DoScan(rootDir string) {
	fmt.Printf("=== 开始扫描文件 \n")
	if err := filepath.Walk(rootDir, visit); err != nil {
		log.Fatal(err)
	}
	doReadPdfFile()
}

func doReadPdfFile() {
	total := len(vars.GlobalFilePathList)
	fmt.Printf("=== 文件总数: %d \n", total)
	// 扫描文件mime信息
	var count = 0
	bar := goPrint.NewBar(100)
	bar.SetNotice("=== 扫描PDF文件：")
	bar.SetGraph(">")
	bar.SetNoticeColor(goPrint.FontColor.Red)
	for _, currentPath := range vars.GlobalFilePathList {
		vars.GlobalFilePath2MimeInfoMap[currentPath] = util.ReadFileMimeInfo(currentPath).String()
		count = count + 1
		bar.PrintBar(util.CalcPercentage(count, total))
	}
	bar.PrintEnd("=== 扫描完毕")
}

func (FileScanner) DoFilter() {
	total := len(vars.GlobalFilePathList)
	var count = 0
	var pdfCnt = 0
	bar := goPrint.NewBar(100)
	bar.SetNotice("=== 过滤已支持的媒体：")
	bar.SetGraph(">")
	for _, globalFilePath := range vars.GlobalFilePathList {
		fileMime := vars.GlobalFilePath2MimeInfoMap[globalFilePath]
		if strings.Contains(fileMime, "pdf") {
			vars.GlobalPdfFilePathList = append(vars.GlobalPdfFilePathList, globalFilePath)
			pdfCnt = pdfCnt + 1
		}
		count = count + 1
		bar.PrintBar(util.CalcPercentage(count, total))
	}
	bar.PrintEnd("=== 过滤完毕, 文件总数：" + strconv.Itoa(pdfCnt))
}

// 定义walkFn回调函数visit
func visit(currentPath string, info os.FileInfo, err error) error {
	if err != nil {
		return err // 如果有错误，直接返回
	}
	if !info.IsDir() {
		vars.GlobalFilePathList = append(vars.GlobalFilePathList, currentPath)
		// filename, include ext
		vars.GlobalFilePath2FileNameMap[currentPath] = filepath.Base(currentPath)
		// file ext
		vars.GlobalFilePath2FileExtMap[currentPath] = path.Ext(currentPath)
	}
	return nil
}
