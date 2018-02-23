package component

import (
	"os"
	"crypto/sha512"
	"encoding/hex"
	"io"
)

type FileHashDetailInfo struct {
	FileFullName string
	Sha512Hash   string
	FileSize     int64
	FileDetailInfo interface{}
}

func GetFileHashDetailInfo(fullFileName string) (FileHashDetailInfo, error) {

	info := FileHashDetailInfo{FileFullName: fullFileName}
	fileInfo, err := os.Stat(fullFileName)
	if err != nil {
		return info, err
	}
	info.FileSize = fileInfo.Size()
	info.FileDetailInfo = fileInfo

	file, err := os.Open(fullFileName)
	if err != nil {
		return info, err
	}
	defer file.Close()

	sha512Hash := sha512.New()
	io.Copy(sha512Hash, file)
	info.Sha512Hash = hex.EncodeToString(sha512Hash.Sum(nil))

	return info, nil
}
