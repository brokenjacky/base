package database

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "time"
    "log"
)

type MysqlClient struct {
    gorm.DB
    closeCh chan struct{}
}

func CreateMysqlClient(addr, dbName, user, pass string) (*MysqlClient, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, pass, addr, dbName)
    db, err := gorm.Open(mysql.Open(dsn))
    if err != nil {
        log.Println("careate mysql client err...", err.Error())
        return nil, err
    }
    client := MysqlClient{
        *db,
        make(chan struct{}),
    }

    go keepAlive(db, client.closeCh)
    return &client, nil
}

func (client *MysqlClient) Close() {
    close(client.closeCh)
}

func keepAlive(db *gorm.DB, ch chan struct{}) {
    tic := time.NewTicker(time.Second * 60)
    for true {
        select {
        case <-tic.C:
            d, _ := db.DB()
            d.Ping()
        case <-ch:
            log.Println("exit db...")
            return
        }
    }
}
