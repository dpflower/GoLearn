package subcomponent

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/deckarep/golang-set"
	"github.com/op/go-logging"
)


var (
	logger = logging.MustGetLogger("submon")
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
)

var videoFormats = []interface{}{
	".webm",
	".mkv",
	".flv",
	".flv",
	".vob",
	".ogv",
	".ogg",
	".drc",
	".gifv",
	".mng",
	".avi",
	".mov",
	".qt",
	".wmv",
	".yuv",
	".rm",
	".rmvb",
	".asf",
	".amv",
	".mp4",
	".m4p",
	".m4v",
	".mpg",
	".mp2",
	".mpeg",
	".mpe",
	".mpv",
	".mpg",
	".mpeg",
	".m2v",
	".m4v",
	".svi",
	".3gp",
	".3g2",
	".mxf",
	".roq",
	".nsv",
	".flv",
	".f4v",
	".f4p",
	".f4a",
	".f4b",
}

var videoFormatsSet = mapset.NewSetFromSlice(videoFormats)

var subFormats = []interface{}{
	".srt",
	".ass",
	".sub",
	".idx",
	".ssa",
	".smi",
}

var pathSep = string(os.PathSeparator)

func isVideoFile(p string) bool {
	ext := filepath.Ext(p)
	return videoFormatsSet.Contains(ext)
}

func isExistsSub(p string, lang string) bool {
	filedir := filepath.Dir(p)
	filenameOnly := strings.TrimSuffix(filepath.Base(p), filepath.Ext(p))

	for _, subExt := range subFormats{
		subfile := fmt.Sprintf("%s%s%s.%s%s", filedir, pathSep, filenameOnly, lang, subExt)
		//fmt.Println("subfile:"+subfile)
		_, err :=os.Stat(subfile)
		if err == nil{
			logger.Info("存在字幕文件：" + p)
			return true
		}
	}


	subfile := fmt.Sprintf("%s%s%s.%s%s", filedir, pathSep, filenameOnly, lang, ".notfind")
	fi, err :=os.Stat(subfile)
	if err == nil{
		modTime := fi.ModTime()
		logger.Info(fmt.Sprintf("存在字幕不存在文件：%s 创建时间：", subfile, modTime.Format("2006-01-02 15:04:05")))
		return true
	}


	return false
}

func WalkDir(filePath string, lang string) {
	err := filepath.Walk(filePath, func(path string, f os.FileInfo, err error) error {
		//logger.Infof("%s", filePath)
		if !isVideoFile(path) {
			return nil
		}
		if isExistsSub(path, lang) {
			return nil
		}

		logger.Info("VideoFile:", path)
		DownloadSub(path, lang)
		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return
}
