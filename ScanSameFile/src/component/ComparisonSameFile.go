package component

import (
	"path/filepath"
	"os"
	"github.com/op/go-logging"
	"fmt"
	"time"
)

var (
	logger = logging.MustGetLogger("scanFile")
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
)

func Comparison(fullPath string) {

	fileHashCodeMap := make(map[string][]FileHashDetailInfo)

	err := filepath.Walk(fullPath, func(path string, f os.FileInfo, err error) error {

		if f.IsDir() {
			return nil
		}
		begin := time.Now()
		detailInfo, err := GetFileHashDetailInfo(path)
		duration := time.Now().Sub(begin)
		if err != nil {
			logger.Error(err)
			return nil
		}

		logger.Infof("FileFullName:%s\r\n Sha512Hash:%s\r\n FileSize:%d\r\n Duration:%v\r\n\r\n", detailInfo.FileFullName, detailInfo.Sha512Hash, detailInfo.FileSize, duration)

		key := fmt.Sprintf("%s-%v", detailInfo.Sha512Hash, detailInfo.FileSize)

		if detailInfos, ok := fileHashCodeMap[key]; ok {
			detailInfos = append(detailInfos, detailInfo)
			fileHashCodeMap[key] = detailInfos
		} else {
			detailInfos := []FileHashDetailInfo{detailInfo}
			fileHashCodeMap[key] = detailInfos
		}

		return nil
	})


	logger.Info("=========================结果===========================")
	for key, detailInfos := range fileHashCodeMap {
		if len(detailInfos) > 1 {
			logger.Infof("========= %v =========", key)
			for _, value := range detailInfos {
				logger.Infof("FileFullName:%v  FileSize:%v ", value.FileFullName, value.FileSize)
			}
			logger.Info("\r\n")
		}
	}

	if err != nil {
		logger.Error(err)
	}

}
