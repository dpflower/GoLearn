package subcomponent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

type FileInfo struct {
	Ext  string // 文件扩展名
	Link string // 文件下载链接
}

type SubInfo struct {
	Desc  string     // 备注信息
	Delay int32      // 字幕相对于视频的延迟时间，单位是毫秒
	Files []FileInfo // 包含文件信息的Array。 注：一个字幕可能会包含多个字幕文件，例如：idx+sub格式
}

func DownloadSub(p string, language string) {
	subInfo, err := fetchSubInfo(p, language)
	if err != nil {
		logger.Error(err)
	}

	for i, _ := range subInfo {
		for _, info := range subInfo[i].Files {
			subfile := generateSubFileName(p, i, info.Ext, language)
			downloadSubData(subfile, info.Link, subInfo[i].Delay)
		}
	}
}

func downloadSubData(subfile string, link string, delay int32) {
	response, req_err := http.Get(link)
	if req_err != nil {
		logger.Error(req_err)
	}
	defer response.Body.Close()

	//_, params, _ := mime.ParseMediaType(response.Header["Content-Disposition"][0])
	//logger.Debugf("%s", response.Body)

	out, os_err := os.Create(subfile)
	if os_err != nil {
		logger.Error(os_err)
	} else {
		_, dl_err := io.Copy(out, response.Body)
		if dl_err != nil {
			logger.Error(dl_err)
		}
	}
	defer out.Close()
	logger.Info("Downloaded " + subfile)

	return
}

func generateSubFileName(p string, index int, ext string, lang string) string {
	basename := p[:strings.LastIndex(p, ".")]
	var subfile string
	if index == 0 {
		subfile = fmt.Sprintf("%s.%s.%s", basename, lang, ext)
	} else {
		subfile = fmt.Sprintf("%s.%s%d.%s", basename, lang, index, ext)
	}
	return subfile
}

func fetchSubInfo(p string, language string) ([]SubInfo, error) {

	logger.Notice("Start searching subtitles for " + path.Base(p))
	Url, err := url.Parse("https://www.shooter.cn/api/subapi.php")

	if err != nil {
		logger.Error(err)
	}

	hash := computeFileHash(p)

	parameters := url.Values{}
	parameters.Add("filehash", hash)
	parameters.Add("pathinfo", p)
	parameters.Add("format", "json")
	parameters.Add("lang", language)
	Url.RawQuery = parameters.Encode()

	req, err := http.NewRequest("POST", Url.String(), bytes.NewBuffer(make([]byte, 0)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	logger.Debug("Response status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)

	var data []SubInfo
	json.Unmarshal(body, &data)
	logger.Debugf("%s subtitles found", len(data))
	return data, err
}
