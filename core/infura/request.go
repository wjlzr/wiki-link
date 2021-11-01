package infura

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"wiki-link/common"
)

var (
	ErrVisitLimit = errors.New("time out")
	ErrMaxRequest = errors.New("max request")
)

type Lock struct {
	Timestamp int64
	Locked    bool
}

func (lock *Lock) isLocked() (locked bool) {
	if lock.Locked && lock.Timestamp > time.Now().Add(-time.Second*60).Unix() {
		locked = true
	}
	return
}

func (lock *Lock) doLock() {
	lock.Timestamp = time.Now().Unix()
	lock.Locked = true
}
func (lock *Lock) unLock() {
	lock.Locked = false
}

var (
	changeKeyLock = Lock{time.Now().Unix(), false}
	index         = 0
	keys          = []string{
		"90dd1b5f220b4252bdaaa99fa62f81b1",
		"982aae15e5f64531abf33d107fc59f25",
	}
)

func InfuraPost(method string, r result, req *AddressBalanceReq) (err error) {
	url := "https://mainnet.infura.io/v3/" + keys[index]
	bs, err := postFromInfura(url, req)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bs, &r); err != nil {
		return
	} else if r.GetErrorCode() == -32005 {
		if !changeKeyLock.isLocked() {
			changeKeyLock.doLock()
			index += 1
			if index >= len(keys) {
				index = 0
			}
		}
		return fmt.Errorf("Infura %w", ErrMaxRequest)
	} else {
		changeKeyLock.unLock()
	}
	return
}

func postFromInfura(url string, abr *AddressBalanceReq) ([]byte, error) {
	q, _ := json.Marshal(abr)
	params := bytes.NewReader(q)
	req, err := http.NewRequest(common.POST, url, params)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	info, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return info, nil
}
