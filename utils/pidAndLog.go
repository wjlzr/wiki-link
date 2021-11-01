package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetLogFile(name ...string) *os.File {
	var app string
	if len(name) == 0 {
		exe := strings.Split(os.Args[0], "/")
		app = exe[len(exe)-1]
	} else {
		app = name[0]
	}
	if err := os.Mkdir("logs", 0755); err != nil {
		if !os.IsExist(err) {
			log.Fatalf("create folder error: %v", err)
		}
	}
	file, err := os.OpenFile("logs/"+app+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	return file
}

func SetLog() {
	file := GetLogFile()
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

func SetPid() {
	if err := os.Mkdir("pids", 0755); err != nil {
		if !os.IsExist(err) {
			log.Fatalf("create folder error: %v", err)
		}
	}
	exe := strings.Split(os.Args[0], "/")
	app := exe[len(exe)-1]
	err := ioutil.WriteFile("pids/"+app+".pid", []byte(strconv.Itoa(os.Getpid())), 0644)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func SetLogAndPid() {
	SetLog()
	SetPid()
}
