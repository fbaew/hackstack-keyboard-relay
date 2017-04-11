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
    KeyboardName, VendorID, ProductID, ManagementPort string
    ManagementHost string
}

func LoadConfig(configFile string) string {
     configData, configLoadError := ioutil.ReadFile(configFile)
    if configLoadError != nil {
        fmt.Println("Erorr loading configuration file.")
        return ""
    }
    return string(configData)
}

func validateConfig(conf Message) (serverConfigValid, clientConfigValid bool) {
    usbConfigValid := conf.ProductID != "" && conf.VendorID != "" && conf.KeyboardName != ""
    managementConfigValid := conf.ManagementPort != ""
    serverConfigValid = usbConfigValid && managementConfigValid
    clientConfigValid = managementConfigValid && conf.ManagementHost != ""
    return
}

func parseConfig(configData string) (Message, error) {

    var m Message
    dec := json.NewDecoder(strings.NewReader(configData))
    for {
        if err := dec.Decode(&m); err == io.EOF {
            break
        } else if err != nil {
            fmt.Println("Problem decoding configuration file. Are you sure it's well-formed?")
            log.Fatal(err)
        }

        serverConfigValid, clientConfigValid := validateConfig(m)

        if  serverConfigValid || clientConfigValid {
            return m, nil
        } else {
            return m, errors.New("Invalid configuration.")
        }
    }
   return m, errors.New("Error decoding JSON")
}

func GetClientConfig(configData string) (Message, error) {
    conf, err := parseConfig(configData)
    if err != nil { log.Fatal("Could not parse client config file") }
    _, clientConfigValid := validateConfig(conf)
    if clientConfigValid {
        return conf, nil
    } else {
        return conf, errors.New("Client configuration invalid. You must specify ManagementHost and ManagementPort.")
    }
}

func GetServerConfig(configData string) (Message, error) {
    conf, _ := parseConfig(configData)
//    if err != nil { log.Fatal("Could not parse server config file") }
    serverConfigValid, _ := validateConfig(conf)
    if serverConfigValid{
        return conf, nil
    } else {
        return conf, errors.New("Server configuration invalid. You must specify ManagementPort" +
            "ProductID, VendorID, and KeyboardName.")
    }
}
