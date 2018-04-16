package config

type item struct {
    Path string
}

type cate struct {
    Index item
    Pic item
    Js item
}

type serverConfig struct {
    Port int
    Path string
    Category cate
}
