/*
    kbserver
    Author: Gregg Lewis (code@gregglewis.net)

    This is a sloppy little server to relay commands to the qemu monitor for the purpose of attaching
    and detaching my keyboard from a guest VM.

    attachKeyboard() and findTargetDevice() both have values hard-coded for my
    personal setup at present. Really, we should read them from a config file
    written by a helper tool.

*/

package main


import (
    "net"
    "bufio"
    "fmt"
    "strings"
    "hackstack-keyboard-relay/cryptoencoder"
    "flag"
    "hackstack-keyboard-relay/config"
    "log"
    "time"
)

func qemuCommand(command string) string{
    conn, err := net.Dial("tcp", "localhost:4445")
    if err != nil {
        fmt.Printf("There was a problem connecting to the monitor:\n")
        fmt.Println(err)
    }

    status,err := bufio.NewReader(conn).ReadString(')')
    fmt.Println(status)
    if err != nil { fmt.Println(err) }
    fmt.Fprintf(conn, command + "\r\n")
    commandResponse,commandErr := bufio.NewReader(conn).ReadString(')')
    if commandErr != nil {
        fmt.Printf("There was a problem with command %s", command)
        fmt.Println(err)
    }

    fmt.Printf("Issued command '%s'\n",command)
    fmt.Printf("Got response:\n%s",commandResponse)
    conn.Close()
    return commandResponse
}

func removeUSBDevice(deviceID string) {
    qemuCommand(fmt.Sprintf("usb_del %s", deviceID))
}

func findTargetDevice(devices []string, conf *config.Message) string {
    targetDevice := conf.KeyboardName //"Corsair K65 Gaming Keyboard"

    deviceID := ""
    fmt.Printf("Searching for target device '%s'---------\n",targetDevice)
    for x:=0; x < len(devices); x++ {

        if strings.Contains(devices[x], targetDevice) {
            fmt.Printf("[%d] %s\n", x, devices[x])
            deviceIDWithPrefix := strings.Split(devices[x],",")[0]
            deviceID = strings.Split(deviceIDWithPrefix," ")[3]
        }
    }
    return deviceID
}

func findKeyboard(conf *config.Message) string {

    guestUSBQueryStatus := qemuCommand("info usb")
    deviceID := ""

    raw_usb_info := strings.Split(guestUSBQueryStatus, "\n")
    fmt.Printf("%d entries in guest usb list\n",len(raw_usb_info))

    deviceID = findTargetDevice(raw_usb_info, conf)
    return deviceID

}

func removeKeyboard(deviceID string) {
    if deviceID == "" {
        fmt.Println("Could not find the device to remove...")
    } else {
        fmt.Printf("Removing device %s\n", deviceID)
        removeUSBDevice(deviceID)
    }
}

func attachKeyboard(conf *config.Message) {
    if findKeyboard(conf) != "" {
        fmt.Println("Device appears to be attached to guest already. Doing nothing.")
    } else {
        fmt.Println("Attaching keyboard")
        qemuCommand(fmt.Sprintf("usb_add host:%s:%s",conf.VendorID, conf.ProductID))
    }
}

func handleCommand(encryptedCommand string, conf *config.Message) {

    key := cryptoencoder.LoadKey("private.key")
    command, decodingError := cryptoencoder.Decode([]byte(encryptedCommand),key)
    if decodingError != nil {
        fmt.Println("Got invalid encrypted data; skipping.")
        return
    }
    fmt.Printf("Decrypted command [%s]\n", command)
    fmt.Printf("Cached device ID: %s\n", conf.QEMUDeviceID)
    switch command {
    case "detach":
        if conf.QEMUDeviceID != "" {
            removeKeyboard(conf.QEMUDeviceID)
        } else {
            removeKeyboard(findKeyboard(conf))
        }
    case "attach":
        attachKeyboard(conf)
        go updateCachedDeviceID(conf)
    default:
        fmt.Println("Got unrecognized command...")
    }
}

func handleConnection(conn net.Conn, conf *config.Message) {
    fmt.Println("Got a new connection:")
    fmt.Println(conn)

    for {
        command, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println(err)
            break
        }
        go handleCommand(command, conf)
    }
    fmt.Println("Closed connection:")
    fmt.Println(conn)
}

func updateCachedDeviceID(conf *config.Message) {
    fmt.Println("Updating cached device id... waiting a sec before doing so.")
    time.Sleep(3 * time.Second)
    conf.QEMUDeviceID = findKeyboard(conf)
}

func main() {
    configFilePointer := flag.String("config", "config.json", "Configuration file")
    flag.Parse()

    conf, configurationError := config.GetServerConfig(config.LoadConfig(*configFilePointer))
    if configurationError != nil {
        log.Fatal(configurationError)
    }

    ln, err := net.Listen("tcp", ":" + conf.ManagementPort)
    if err != nil {
        fmt.Println("Error starting server")
        fmt.Println(err)
        return
    }
    fmt.Printf("Started listening on port %s\n",conf.ManagementPort)

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting new connection")
            fmt.Println(err)
        }
        go handleConnection(conn, &conf)
    }
}
