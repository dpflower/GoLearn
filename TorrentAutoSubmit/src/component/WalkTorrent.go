package component

import (
	"io/ioutil"
	"github.com/op/go-logging"
	"os"
	"path/filepath"
	"fmt"
)

var (
	logger = logging.MustGetLogger("submon")
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
)
var pathSep = string(os.PathSeparator)

func WalkTorrent(path string) {
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
		err := submitTorrentFile(path, info)
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

func submitTorrentFile(path string, info os.FileInfo) error {
	logger.Info("SubmitTorrent:" + info.Name())
	return nil
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
