package subcomponent

import (
	"path/filepath"
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"encoding/json"
	"io"
	"github.com/mholt/archiver"
	"os"
)

type SubUrlInfo struct {
	Success bool   `json:"success"`
	Url     string `json:"url"`
	Msg     string `json:"msg"`
}

var subhdSearchUrl = "http://subhd.com/search0/"
var subhdAjaxUrl = "http://subhd.com/ajax/down_ajax"

var filter = ".row .col-md-9 .box .lb_r"

func SubHDDownload(p string, language string) {
	subid, err := getSubId(p, language)
	if err != nil {
		logger.Warning(err)
		return
	}
	logger.Info("SubId:" + subid)

	subUrl, err := getSubUrl(subid)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(subUrl)

	//subUrl := "http://dl.subhd.com/sub/2017/10/150735700219240.rar"
	tempFileName, err := downfileToTempPath(subUrl)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(tempFileName)

	//tempFileName := "150735700219240.rar280521659"

	ext := string(subUrl[strings.LastIndex(subUrl, ".")+1:])
	logger.Info(ext)
	decompressionFile(tempFileName, p, ext)

}

func getSubId(p string, language string) (string, error) {
	fileName := filepath.Base(p)
	fmt.Println("fileName:" + fileName)

	document, err := goquery.NewDocument(subhdSearchUrl + fileName)
	if err != nil {
		logger.Error(err)
	}

	s := *document.Find(filter).First()
	a := s.Find("a")
	band := a.Text()
	val, _ := a.Attr("href")
	logger.Info(band)
	logger.Info(val)
	subid := val[strings.LastIndex(val, "/")+1:]

	return subid, nil
}

func downfileToTempPath(subUrl string) (string, error) {
	Url, err := url.Parse(subUrl)
	if err != nil {
		logger.Error(err)
	}

	tempFileName := subUrl[strings.LastIndex(subUrl, "/")+1:]

	request, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		logger.Error(err)
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logger.Error(err)
	}
	defer response.Body.Close()

	tempFile, err := ioutil.TempFile(".", tempFileName)
	if err != nil {
		logger.Error(err)
	}
	defer tempFile.Close()

	_, errc := io.Copy(tempFile, response.Body)
	if errc != nil {
		logger.Error(errc)
	}
	return tempFile.Name(), err
}

func getSubUrl(subid string) (string, error) {
	Url, err := url.Parse(subhdAjaxUrl)
	if err != nil {
		logger.Error(err)
	}

	s := "sub_id=" + subid
	request, err := http.NewRequest("POST", Url.String(), strings.NewReader(s))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		logger.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logger.Error(err)
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	logger.Info(string(body))
	var data SubUrlInfo
	json.Unmarshal(body, &data)
	logger.Info(data)
	if (data.Success) {
		return data.Url, nil
	}
	return "http://dl.subhd.com/sub/2017/10/150735700219240.rar", nil
}

func decompressionFile(tempFileName string, p string, ext string) {
	fileDir := filepath.Dir(p)
	logger.Info(fileDir)

	var err error = nil

	switch strings.ToLower(ext) {
	case "rar":
		err = archiver.Rar.Open(tempFileName, fileDir)
	case "zip":
		err = archiver.Zip.Open(tempFileName, fileDir)
	//case "7z":
	//	err = archiver..Open(tempFileName, fileDir)
	}

	if err != nil {
		logger.Error(err)
	}

	os.Remove(tempFileName)

}

func getSubId2(p string, language string) (string, error) {
	fileName := filepath.Base(p)
	fmt.Println("fileName:" + fileName)
	Url, err := url.Parse(subhdSearchUrl + fileName)
	if err != nil {
		logger.Error(err)
	}

	request, err := http.NewRequest("GET", Url.String(), nil)
	if err != nil {
		logger.Error(err)
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logger.Error(err)
	}
	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error(err)
	}
	body := string(bytes)
	logger.Info(body)

	return "", nil
}
