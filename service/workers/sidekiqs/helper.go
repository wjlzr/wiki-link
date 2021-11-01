package sidekiqs

import (
	"errors"
	"time"
)

var (
	ErrMaxRequest    = errors.New("max request")
	BtcChannelChange Lock
)

type Lock struct {
	Timestamp int64
	Locked    bool
}

func (lock *Lock) isLocked() (locked bool) {
	if lock.Locked && lock.Timestamp > time.Now().Add(-time.Second*10).Unix() {
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
