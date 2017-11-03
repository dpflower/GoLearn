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
)

type SubUrlInfo struct {
	Success  bool `json:"success"`
	Url string `json:"url"`
	Msg string `json:"msg"`
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

	tempFileName, err := downfileToTempPath(subid)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(tempFileName)

	decompressionFile(tempFileName, p)

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

func downfileToTempPath(subid string) (string, error) {

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
	if(data.Success){
		return data.Url, nil
	}
	return "http://dl.subhd.com/sub/2017/10/150735700219240.rar", nil
}

func decompressionFile(tempFileName string, p string) {
	fileDir := filepath.Dir(p)
	logger.Info(fileDir)

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
