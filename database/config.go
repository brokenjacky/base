package database

type Config struct {
    Enable bool   `json:"enable"`
    Addr string `json:"addr"`
    User string `json:"user"`
    Pass string `json:"pass"`
    Name string `json:"name"`
}
