package config

import (
    "log"
    "os"
    "encoding/json"
)

func GetConfig(path string) *serverConfig {
    file, err := os.Open(path)
    if(err != nil) {
        log.Println(err)
        return nil
    }
    defer file.Close()

    stat, err := file.Stat()

    if(err != nil) {
        log.Println(err)
        return nil
    }

    data := make([]byte, stat.Size())
    _, err = file.Read(data)

    var c serverConfig
    err = json.Unmarshal(data, &c)
    if(err != nil) {
        log.Println(err)
        return nil
    }

    log.Println(c);

    return &c
}

func GetByteArray(config *serverConfig) []byte {
    b, err := json.Marshal(config);
    if(err != nil) {
        log.Println("GetByteArray convert failed");
        return nil;
    }
    return b;
}
