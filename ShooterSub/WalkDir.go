package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/deckarep/golang-set"
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

func isVideoFile(p string) bool {
	ext := filepath.Ext(p)
	return videoFormatsSet.Contains(ext)
}

func isExistsSub(p string, lang string) bool {
	filenameOnly := strings.TrimSuffix(filepath.Base(p), filepath.Ext(p))
	subfile := fmt.Sprintf("%s.%s*", filenameOnly, lang)
	matchs, _ := filepath.Glob(filepath.Join(filepath.Dir(p), subfile))
	if len(matchs) > 0 {
		return true
	}

	// logger.Info("SubFile:", subfile)
	// dir := filepath.Dir(p)
	// fileinfos, err := ioutil.ReadDir(dir)
	// if err != nil {
	// 	logger.Error(err)
	// }
	// for _, fileinfo := range fileinfos {
	// 	if fileinfo.IsDir() {
	// 		continue
	// 	}
	// 	//logger.Info(fileinfo.Name())
	// }
	// //basename := p[:strings.LastIndex(p, ".")]

	return false
}

func walkDir(filePath string, lang string) {
	err := filepath.Walk(filePath, func(path string, f os.FileInfo, err error) error {
		//logger.Infof("%s", filePath)
		if !isVideoFile(path) {
			return nil
		}
		if isExistsSub(path, lang) {
			return nil
		}

		downloadSub(path, lang)
		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return
}
