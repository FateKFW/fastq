package main

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var logger = new(Fastlog)

type Fastlog struct {
	logLevel int
	logger *log.Logger
}

//level 记录日志级别 4：error 3：warring及以上 2：info及以上 1：debug及以上
func (fl *Fastlog) initFastlog(level int){
	logFile, _ := os.Create("."+ string(filepath.Separator) + time.Now().Format("20060102_150405") + ".txt")
	fl.logger = log.New(logFile, PROJECT_NAME+" ", log.Ldate | log.Ltime)
	fl.logLevel = level
}

func (fl *Fastlog) debug(content string){
	if fl.logLevel <= 1 {
		fl.logger.Printf(" DEBUG: %v", content)
	}
}

func (fl *Fastlog) info(content string){
	if fl.logLevel <= 2 {
		fl.logger.Printf(" INFO: %v", content)
	}
}

func (fl *Fastlog) warring(content string){
	if fl.logLevel <= 3 {
		fl.logger.Printf(" WARRING: %v", content)
	}
}


func (fl *Fastlog) error(err interface{}){
	if fl.logLevel <= 4 {
		fl.logger.Panicf(" ERROR: %v", err)
	}
}

func (fl *Fastlog) result(title string,content string){
	fl.logger.Printf(" RESULT>>%v\n%v", title, content)
}