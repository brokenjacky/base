package database

import (
    "github.com/go-xorm/xorm"
    "fmt"
    "xorm.io/core"
    "context"
    "time"
)

type XormEngin struct {
    *xorm.Engine
    c      context.Context
    cancel context.CancelFunc
}

func NewXormEngine(config Config) (*XormEngin, error) {
    sourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", config.User, config.Pass, config.Addr, config.Name)
    engine, err := xorm.NewEngine("mysql", sourceName)
    if err != nil {
        return nil, err
    }
    err = engine.Ping()
    if err != nil {
        return nil, err
    }

    engine.SetLogLevel(core.LOG_DEBUG)
    // debug
    engine.ShowSQL(false)
    c, cancel := context.WithCancel(context.Background())
    xormEngin := &XormEngin{
        Engine: engine,
        c:      c,
        cancel: cancel,
    }
    go xormEngin.keepAlive()
    return xormEngin, nil
}

func (e *XormEngin) keepAlive() {
    tic := time.NewTicker(time.Second * 60)
    for true {
        select {
        case <-tic.C:
            e.Ping()
        case <-e.c.Done():
            e.Close()
            return
        }
    }
}

func (e *XormEngin) Close() {
    e.cancel()
}
