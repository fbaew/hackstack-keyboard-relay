package main


import (
    "net"
    "fmt"
    "log"
    "flag"
    "hackstack-keyboard-relay/cryptoencoder"
    "hackstack-keyboard-relay/config"
)

func printUsage() {
    fmt.Println("")
    fmt.Println("usage: kbclient <-attach|-detach> [-key=private.key]")
    fmt.Println("")
}

func main() {
    keyfilePointer := flag.String("key", "private.key", "path to private key")
    attachPointer := flag.Bool("attach", false, "attach keyboard to virtual guest")
    detachPointer := flag.Bool("detach",false, "detach keyboard from virtual guest")
    configFilePointer := flag.String("config", "config.json", "path to configuration file")

    flag.Parse()

    if *attachPointer && *detachPointer {
        printUsage()
        return
    } else if !(*attachPointer || *detachPointer) {
        printUsage()
        return
    }

    conf, configurationError := config.GetClientConfig(config.LoadConfig(*configFilePointer))
    if configurationError != nil {
        log.Fatal(configurationError)
    }

    key := cryptoencoder.LoadKey(*keyfilePointer)
    if *attachPointer { connectToMonitorServer("attach", key, &conf)
    } else if *detachPointer {
        connectToMonitorServer("detach", key, &conf)
    } else {}
}

func connectToMonitorServer(command string, key *[32]byte, conf *config.Message) {
    conn, err := net.Dial("tcp",conf.ManagementHost + ":" + conf.ManagementPort)
    if err != nil {
        fmt.Println("There was a problem connecting to the monitor server")
        fmt.Println(err)
        fmt.Println("-----------------------------\n")
        log.Fatal("Terminating")
    }

//    key := cryptoencoder.LoadKey("private.key")
    encryptedCommand,encodingError := cryptoencoder.Encode(command, key)
    if encodingError != nil { log.Fatal("Problem encrypting command." )}

    fmt.Fprintf(conn,"%s",encryptedCommand)
    //commandResponse,commandError := bufio.NewReader(conn).ReadString('\n')
    conn.Close()
}
