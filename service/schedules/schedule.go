package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/robfig/cron/v3"

	"wiki-link/initializers"
	"wiki-link/service/schedules/address"
	"wiki-link/service/schedules/backup/tasks"
	"wiki-link/service/schedules/blockHeight"
	"wiki-link/utils"
)

func main() {
	initializers.InitAllResources()
	defer initializers.CloseResources()

	address.InitWorker()
	blockHeight.InitWorker()

	initSchedule()

	utils.SetLogAndPid()
	err := ioutil.WriteFile("pids/schedule.pid", []byte(strconv.Itoa(os.Getpid())), 0644)
	if err != nil {
		log.Println(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	closeResource()
}

func closeResource() {
}

func initSchedule() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("*/30 * * * * *", address.CheckAddress)
	c.AddFunc("30 */10 * * * *", address.ReRunErrors)

	c.AddFunc("*/10 * * * * *", blockHeight.GetBtcLastestBlockHeight)
	c.AddFunc("*/10 * * * * *", blockHeight.CheckBtcBlockHeight)
	c.AddFunc("*/10 * * * * *", blockHeight.GetEthLastestBlockHeight)
	c.AddFunc("*/10 * * * * *", blockHeight.CheckEthBlockHeight)
	c.AddFunc("*/30 * * * * *", blockHeight.ReRunErrors)

	// 日志备份
	c.AddFunc("0 55 23 * * *", tasks.BackupLogFiles)
	c.AddFunc("0 56 23 * * *", tasks.UploadLogFileToQiniu)
	c.AddFunc("0 59 23 * * *", tasks.CleanLogs)

	c.Start()
}
