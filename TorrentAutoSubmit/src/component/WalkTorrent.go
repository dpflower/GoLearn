package component

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"github.com/op/go-logging"
)

type Aria2Request struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Id      int           `json:"id"`
	Params  []interface{} `json:"params"`
}

type Aria2Options struct {
	Split                  string `json:"split"`
	MaxConnectionPerServer string `json:"max-connection-per-server"`
	SeedRatio              string `json:"seed-ratio"`
	SeedTime               string `json:"seed-time"`
}

var (
	logger = logging.MustGetLogger("submon")
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
)
var pathSep = string(os.PathSeparator)

func WalkTorrent(path string, url string) {
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		logger.Error(err)
	}

	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		ext := filepath.Ext(info.Name())
		if ext != ".torrent" {
			continue
		}
		err := submitTorrentFile(path, info, url)
		if err != nil {
			logger.Error(err)
			continue
		}
		err = removeTorrentFile(path, info)
		if err != nil {
			logger.Error(err)
		}
	}
}

func submitTorrentFile(path string, info os.FileInfo, url string) error {
	content, err := readTorrentFile(path, info)
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Info(content)
	auth := RequestAuth(url)
	logger.Info(auth)

	return nil
}

func RequestAuth(url string) string {
	if url == "" {
		return ""
	}
	reg, err := regexp.Compile("(?:^http://)([^@]*)(?:@)")
	if err != nil {
		logger.Error(err)
		return ""
	}

	submatch := reg.FindStringSubmatch(url)
	logger.Info(submatch[1])

	return submatch[1]
}

func readTorrentFile(path string, info os.FileInfo) (string, error) {
	logger.Info("SubmitTorrent:" + info.Name())
	fullName := fmt.Sprintf("%s%s%s", path, pathSep, info.Name())
	file, err := os.Open(fullName)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	defer file.Close()
	bytes, err2 := ioutil.ReadAll(file)
	if err2 != nil {
		logger.Error(err)
		return "", err
	}
	content := base64.StdEncoding.EncodeToString(bytes)
	return content, nil
}

func removeTorrentFile(path string, info os.FileInfo) error {
	dir := fmt.Sprintf("%s%s%s", path, pathSep, "oldTorrent")
	_, err := os.Stat(dir)
	if err != nil {
		os.MkdirAll(dir, 0777)
	}
	oldFileName := fmt.Sprintf("%s%s%s", path, pathSep, info.Name())
	newFileName := fmt.Sprintf("%s%s%s", dir, pathSep, info.Name())
	os.Rename(oldFileName, newFileName)
	logger.Infof("文件：%s 已移入 %s 目录", info.Name(), "oldTorrent")
	return err
}
