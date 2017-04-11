package config

import (
    "fmt"
    "io/ioutil"
    "encoding/json"
    "strings"
    "io"
    "log"
    "errors"
)
type Message struct {
    KeyboardName, VendorID, ProductID string
}
//func main() {
//    config := LoadConfig("config.json")
//    msg, err := parseConfig(config)
//    if err != nil { log.Fatal(err) }
//    fmt.Printf("msg: %s\n",msg)
//}

func LoadConfig(configFile string) string {
     configData, configLoadError := ioutil.ReadFile(configFile)
    if configLoadError != nil {
        fmt.Println("Erorr loading configuration file.")
        return ""
    }
    return string(configData)
}

func ParseConfig(configData string) (Message, error) {

    var m Message
    dec := json.NewDecoder(strings.NewReader(configData))
    for {
        if err := dec.Decode(&m); err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("'%s' - %s:%s\n",
            m.KeyboardName,
            m.VendorID,
            m.ProductID)

        if m.ProductID != "" && m.VendorID != "" && m.KeyboardName != "" {
            return m, nil
        } else {
            fmt.Println("Incomplete config. You must specify ProductID, VendorID, and KeyboardName!")
            return m, errors.New("Invalid configuration.")
        }
    }
   return m, errors.New("Error decoding JSON")
}
