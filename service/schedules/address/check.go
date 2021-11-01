package address

import (
	"errors"
	"strconv"
	"time"

	"github.com/oldfritter/sidekiq-go"
	"gorm.io/gorm"

	"wiki-link/common"
	"wiki-link/db"
	"wiki-link/model"
)

var (
	currentId = 0
	addressB  = "address:b"
)
var worker sidekiq.WorkerI

func ReRunErrors() {
	worker.ReRunErrors()
}

func InitWorker() {
	client := db.RedisClient()
	for _, w := range common.AllWorkers {
		if w.GetName() == "AddressCheck" {
			worker = common.AllWorkerIs[w.Name](&w)
			worker.SetClient(client)
		}
	}
}

func CheckAddress() {
	if worker.GetMaxQuery() < worker.GetQuerySize() {
		return
	}
	client := db.RedisClient()
	var currentId int
	if client.Get(addressB).Val() == "" {
		client.Set(addressB, 0, 0)
		currentId = 0
	} else {
		currentId, _ = strconv.Atoi(client.Get(addressB).Val())
	}
	var addresses []model.Address
	if err := db.MainDb.Where("id > ?", currentId).
		Order("id ASC").Limit(200).Find(&addresses).Error; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	if len(addresses) < 200 {
		currentId = 0
		return
	}
	for _, address := range addresses {
		if currentId < int(address.ID) {
			currentId = int(address.ID)
			client.Set(addressB, currentId, 0)
		}
		if address.UpdatedAt.Before(time.Now().Add(-time.Hour * 24 * 3)) {
			message := map[string]string{"id": strconv.Itoa(int(address.ID)), "address": address.Address, "chain": address.Coin}
			worker.Perform(message)
		}
	}
}
