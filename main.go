package main

import (
	_ "embed"
	"fmt"
	"github.com/fatih/color"
	"kenaito-pdf2png/core"
	"kenaito-pdf2png/vars"
	"os"
	"time"
)

func main() {
	color.Yellow("===========================\n")
	color.Yellow("=== 应用名：PDF批处理 V1.0 ===\n")
	color.Yellow("=== 作者：odboy")
	color.Yellow("=== 开源地址：https://github.com/odboy-tianjun/kenaito-pdf2png")
	color.Yellow("===========================\n\n")
	fmt.Println("=== 警告：该程序将在5秒后开始执行")
	time.Sleep(time.Second * 5)
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("=== 获取当前路径异常", err)
		return
	}
	scanner := core.FileScanner{}
	scanner.DoScan(rootDir)
	scanner.DoFilter()
	if len(vars.GlobalPdfFilePathList) == 0 {
		fmt.Println("=== 未扫描到PDF文件, 该程序将在5秒后退出, 或者直接关闭窗口也是可以的", err)
		time.Sleep(time.Second * 5)
		return
	}
	core.DoHandleImage()
	fmt.Println("=== 批处理完毕, 该程序将在5秒后退出, 或者直接关闭窗口也是可以的")
	time.Sleep(time.Second * 5)
}
