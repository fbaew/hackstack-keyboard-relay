package main


import (
    "net"
    "fmt"
    "os"
    "cryptoencoder"
    "log"
)

func printUsage() {
    fmt.Println("")
    fmt.Println("usage: kbclient [attach|detach]")
    fmt.Println("")
}

func main() {
    if len(os.Args) != 2 {
        printUsage()
    } else {
        args := os.Args[1:]

        switch args[0] {
            case "attach":
                connectToMonitorServer("attach")
            case "detach":
                connectToMonitorServer("detach")
            default:
                printUsage()
        }
    }
}

func connectToMonitorServer(command string) {
    conn, err := net.Dial("tcp","192.168.0.13:7357")
    if err != nil {
        fmt.Println("There was a problem connecting to the monitor server")
        fmt.Println(err)
        fmt.Println("-----------------------------\n")
    }

    key := cryptoencoder.LoadKey("private.key")
    encryptedCommand,encodingError := cryptoencoder.Encode(command, key)
    if encodingError != nil { log.Fatal("Problem encrypting command." )}

    fmt.Fprintf(conn,"%s",encryptedCommand)
    //commandResponse,commandError := bufio.NewReader(conn).ReadString('\n')
    conn.Close()
}
