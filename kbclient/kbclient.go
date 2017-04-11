package main


import (
    "net"
    "fmt"
    "cryptoencoder"
    "log"
    "flag"
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

    flag.Parse()

    if *attachPointer && *detachPointer {
        printUsage()
        return
    } else if !(*attachPointer || *detachPointer) {
        printUsage()
        return
    }

    key := cryptoencoder.LoadKey(*keyfilePointer)
    if *attachPointer { connectToMonitorServer("attach", key)
    } else if *detachPointer {
        connectToMonitorServer("detach", key)
    } else {}
}

func connectToMonitorServer(command string, key *[32]byte) {
    conn, err := net.Dial("tcp","192.168.0.13:7357")
    if err != nil {
        fmt.Println("There was a problem connecting to the monitor server")
        fmt.Println(err)
        fmt.Println("-----------------------------\n")
    }

//    key := cryptoencoder.LoadKey("private.key")
    encryptedCommand,encodingError := cryptoencoder.Encode(command, key)
    if encodingError != nil { log.Fatal("Problem encrypting command." )}

    fmt.Fprintf(conn,"%s",encryptedCommand)
    //commandResponse,commandError := bufio.NewReader(conn).ReadString('\n')
    conn.Close()
}
