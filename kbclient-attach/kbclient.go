package main


import (
    "net"
    "fmt"
//    "bufio"
)

func main() {
    connectToMonitorServer()
}

func connectToMonitorServer() {
    conn, err := net.Dial("tcp","192.168.0.13:7357")
    if err != nil {
        fmt.Println("There was a problem connecting to the monitor server")
        fmt.Println(err)
        fmt.Println("-----------------------------\n")
    }

    fmt.Fprintf(conn,"attach\n")
    //commandResponse,commandError := bufio.NewReader(conn).ReadString('\n')
    
    conn.Close()
}
