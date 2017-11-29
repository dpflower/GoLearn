package main

import (
	"./component"
	"github.com/op/go-logging"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"fmt"
	"path/filepath"
	"time"
)

var (
	logger = logging.MustGetLogger("submon")
	format = logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)


)


var (
	app = kingpin.New("TorrentSubmit", "").Version("0.1")

	submit = app.Command("submit", "Download subtitle for specific file")
	submitPath = submit.Arg("path", "target file path").Required().String()
	logfile = submit.Flag("logfile", "LogFile Path defult ./log/TorrentSubmit.log.yyyyMMdd.txt").Default("").String()

)

func main()  {

	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	logFilePath  := *logfile

	if len(logFilePath) == 0{
		logFilePath = fmt.Sprintf("%s/log/TorrentSubmit.log.%s.txt",filepath.Dir(os.Args[0]), time.Now().Format("20060102"))
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
	case submit.FullCommand():

		logger.Info(fmt.Sprintf("开始遍历目录：%s，开始时间：%s", *submitPath, time.Now().Format("2006-01-02 15:04:05")))
		component.WalkTorrent(*submitPath)
		logger.Info("遍历完成！")
	default:
		kingpin.Usage()

	}


}

