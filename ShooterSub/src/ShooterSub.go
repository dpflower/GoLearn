package main

import (
	"./subcomponent"
	"fmt"
	logging "github.com/op/go-logging"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"time"
	"path/filepath"
)

var (
	logger = logging.MustGetLogger("submon")
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)


)



var (
	app = kingpin.New("shootersub", "").Version("0.1")

	download = app.Command("download", "Download subtitle for specific file")
	downloadPath = download.Arg("path", "target file path").Required().String()
	lang = download.Flag("lang", "language, choice: chn, eng").Default("chn").String()
	logfile = download.Flag("logfile", "LogFile Path defult ./log/ShooterSub.log.yyyyMMdd.txt").Default("").String()

)

func main() {

	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	logFilePath  := *logfile

	if len(logFilePath) == 0{
		logFilePath = fmt.Sprintf("%s/log/ShooterSub.log.%s.txt",filepath.Dir(os.Args[0]), time.Now().Format("20060102"))
	}

	fileDir := filepath.Dir(logFilePath)
	_, err :=os.Stat(fileDir)
	if err != nil{
		os.MkdirAll(fileDir,  0777)
	}

	_, err =os.Stat(logFilePath)
	if err != nil{
		os.Create(logFilePath)
	}

	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY,0666)
	if err != nil{
		fmt.Println(err)
	}
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)


	backend1Formatter := logging.NewBackendFormatter(backend1, format)
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend1Leveled := logging.AddModuleLevel(backend1Formatter)
	backend1Leveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backend1Leveled, backend2Formatter)



	switch command {
	case download.FullCommand():

		logger.Info(fmt.Sprintf("开始遍历目录：%s，字幕语言：%s, 开始时间：%s", *downloadPath, *lang, time.Now().Format("2006-01-02 15:04:05")))
		subcomponent.WalkDir(*downloadPath, *lang)
		logger.Info("遍历完成！")
	default:
		kingpin.Usage()

	}

}
