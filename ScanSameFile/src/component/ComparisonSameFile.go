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

var (
	GOROUTINE_COUNT = 20
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


func ConcurrentComparison(fullPath string){

	fileHashCodeMap := make(map[string][]FileHashDetailInfo)
	chFullPath := make(chan string, 100)
	chFileHashInfo := make(chan FileHashDetailInfo, 100)
	chHashDetailGoRoutine := make(chan int, GOROUTINE_COUNT)
	chResult := make(chan int, 1)

	
	go func() error{
		err := filepath.Walk(fullPath, func(path string, f os.FileInfo, err error) error {
			if f.IsDir() {
				return nil
			}
			chFullPath <- path	
			return nil
		})
		close(chFullPath)
		return err
	}()

	
	for i := 0; i < GOROUTINE_COUNT; i++ {
		go func() {
            for {
				path, ok := <-chFullPath
				if !ok {
					chHashDetailGoRoutine <- 1
					break;
				}
				begin := time.Now()
				detailInfo, err := GetFileHashDetailInfo(path)
				duration := time.Now().Sub(begin)
				if err != nil {
					logger.Error(err)
					continue
				}
				logger.Infof("FileFullName:%s\r\n Sha512Hash:%s\r\n FileSize:%d\r\n Duration:%v\r\n\r\n", detailInfo.FileFullName, detailInfo.Sha512Hash, detailInfo.FileSize, duration)
                chFileHashInfo <- detailInfo
            }
        }()
	}

	go func(){
		for i := 0; i < GOROUTINE_COUNT; i++ {
			<-chHashDetailGoRoutine
		}
		close(chFileHashInfo)
	}()

	
		
	go func() {
		for {
			detailInfo, ok := <-chFileHashInfo
			if !ok {
				chResult <- 1
				break;
			}
			key := fmt.Sprintf("%s-%v", detailInfo.Sha512Hash, detailInfo.FileSize)

			if detailInfos, ok := fileHashCodeMap[key]; ok {
				detailInfos = append(detailInfos, detailInfo)
				fileHashCodeMap[key] = detailInfos
			} else {
				detailInfos := []FileHashDetailInfo{detailInfo}
				fileHashCodeMap[key] = detailInfos
			}
		}
	}()

	


	<- chResult

	logger.Info("=========================结果===========================")
	logger.Info(len(fileHashCodeMap))
	for key, detailInfos := range fileHashCodeMap {
		if len(detailInfos) > 1 {
			logger.Infof("========= %v =========", key)
			for _, value := range detailInfos {
				logger.Infof("FileFullName:%v  FileSize:%v ", value.FileFullName, value.FileSize)
			}
			logger.Info("\r\n")
		}
	}

}