package database

import (
    "time"
    "github.com/go-redis/redis/v8"
    "errors"
)

type RedisConfig struct {
    Enable       bool     `json:"enable" yaml:"enable"`
    Addresses    []string `json:"addresses" yaml:"addresses"`
    Password     string   `json:"password" yaml:"password"`           // 密码
    DialTimeout  uint64   `json:"dial_timeout" yaml:"dial_timeout"`   // 单位毫秒
    ReadTimeout  uint64   `json:"read_timeout" yaml:"read_timeout"`   // 单位毫秒
    WriteTimeout uint64   `json:"write_timeout" yaml:"write_timeout"` // 单位毫秒
    ClusterMode  bool     `json:"cluster_mode" yaml:"cluster_mode"`   // 是否集群模式
}

func CreateRedisClient(conf *RedisConfig) (redis.Cmdable, error) {

    if len(conf.Addresses) == 0 {
        return nil, errors.New("addresses is empty")
    }
    if conf.ClusterMode {
        var redisCli *redis.ClusterClient

        redisCli = redis.NewClusterClient(&redis.ClusterOptions{
            Addrs:        conf.Addresses,
            Password:     conf.Password,
            DialTimeout:  time.Duration(conf.DialTimeout) * time.Millisecond,
            ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Millisecond,
            WriteTimeout: time.Duration(conf.WriteTimeout) * time.Millisecond,
        })

        return redisCli, nil
    }

    var redisCli *redis.Client

    redisCli = redis.NewClient(&redis.Options{
        Addr:         conf.Addresses[0],
        Password:     conf.Password,
        DialTimeout:  time.Duration(conf.DialTimeout) * time.Millisecond,
        ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Millisecond,
        WriteTimeout: time.Duration(conf.WriteTimeout) * time.Millisecond,
    })

    return redisCli, nil

}
