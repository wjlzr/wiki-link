package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/oldfritter/sidekiq-go"

	"wiki-link/common"
	"wiki-link/db"
	"wiki-link/initializers"
	"wiki-link/utils"
)

var closeChan = make(chan int)

func main() {
	initializers.InitAllResources()
	utils.SetLogAndPid()
	startAllWorkers()

	defer func() {
		initializers.CloseResources()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	go recycle()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	select {
	case <-closeChan:
		cancel()
	case <-ctx.Done():
		cancel()
	}
}

func startAllWorkers() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("Wiki-Link"),
		newrelic.ConfigLicense("fd4037a09661ecc12378f9da59b161e4a88c9c7e"),
	)
	if err != nil {
		os.Exit(1)
	}
	if err := app.WaitForConnection(5 * time.Second); nil != err {
		fmt.Println(err)
	}

	for _, worker := range common.AllWorkers {
		for i := 0; i < worker.Threads; i++ {
			w := common.AllWorkerIs[worker.Name](&worker)
			initWorker(w)

			go func(w sidekiq.WorkerI) {
				run(w, app)
			}(w)
			common.SWI = append(common.SWI, w)
			log.Println("started: ", w.GetName(), "[", i, "]")
		}
	}
}

func run(w sidekiq.WorkerI, app *newrelic.Application) {
	txn := app.StartTransaction(w.GetName())
	if d, err := sidekiq.Run(w); d && err == nil {
		txn.End()
		time.Sleep(time.Second * 10)
	} else {
		txn.End()
	}
	app.RecordCustomEvent(w.GetName(), map[string]interface{}{
		"color": "blue",
	})
	run(w, app)
}

func initWorker(w sidekiq.WorkerI) {
	w.SetClient(db.RedisClient())
	w.InitLogger()
	w.RegisterQueue()
}

func recycle() {
	for i, w := range common.SWI {
		w.Stop()
		log.Println("stoped: ", w.GetName(), "[", i, "]")
	}
	for _, w := range common.SWI {
		w.Recycle()
	}
	closeChan <- 1
}
