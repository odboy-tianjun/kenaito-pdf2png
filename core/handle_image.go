package core

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"kenaito-pdf2png/vars"
	"path/filepath"
)

func DoHandleImage() {
	total := len(vars.GlobalPdfFilePathList) // 总数
	successCount := 0                        // 成功数
	errorCount := 0                          // 失败数

	for index, pdfFilePath := range vars.GlobalPdfFilePathList {
		//suffix := vars.GlobalFilePath2FileExtMap[pdfFilePath]
		pdfFileFullName := filepath.Base(pdfFilePath)
		pdfFileDir := filepath.Dir(pdfFilePath)
		// 提取图片
		err := api.ExtractImagesFile(pdfFilePath, pdfFileDir, nil, nil)
		if err != nil {
			errorCount = errorCount + 1
		} else {
			successCount = successCount + 1
		}
		fmt.Printf("=== 当前处理进度: %s | 总数：%d | 已完成：%d \n", pdfFileFullName, total, index+1)
	}
}
