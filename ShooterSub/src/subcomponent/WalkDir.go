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
		logger.Info("存在字幕文件：" + p)
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
		//downloadSub(path, lang)
		return nil
	})

	if err != nil {
		logger.Error(err)
	}

	return
}
