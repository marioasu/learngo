package main

import (
	"crypto/rand"
	"fmt"
	"sync"
)

func main() {

}

/*
   session 管理器
*/
type Manager struct {
	cookieName  string     // private cookiename
	lock        sync.Mutex //protects ssion
	provider    Provider
	maxlifetime int64
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error //delete session value
	SessionID() string            // back current sessionID
}
