package util

import (
	"github.com/gabriel-vasile/mimetype"
	"log"
)

// ReadFileMimeInfo 获取文件mime信息
func ReadFileMimeInfo(filepath string) *mimetype.MIME {
	mt, err := mimetype.DetectFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return mt
}
