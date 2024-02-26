package commands

import (
    "testing"
    "reflect"
)

func Test_Is_USBIP_Available(t *testing.T) {
    value := Is_USBIP_Available()
    target := true
    if value != target {
        t.Errorf("target: %t value: %t", target, value)
    }
}

func Test_Is_USBIP_Available_Remote(t *testing.T) {
    ip := "192.168.178.39"
    con := Establish_SSHConnection_With_PrivateKey(ip, SSHKeyPath)
    value := con.Is_USBIP_Available()
    target := true
    if value != target {
        t.Errorf("target: %t value: %t", target, value)
    }
}

func Test_Build_Host_Map(t *testing.T) {
    nmapOutput := `
    Starting Nmap 7.80 ( https://nmap.org ) at 2024-02-24 16:33 CET
    Nmap scan report for raspberrypi.fritz.box (192.168.178.24)
    Host is up (0.00028s latency).

    PORT   STATE SERVICE
    22/tcp open  ssh

    Nmap scan report for raspberrypi.fritz.box (192.168.178.39)
    Host is up (0.00029s latency).

    PORT   STATE SERVICE
    22/tcp open  ssh

    Nmap scan report for maximilian1-P6634.fritz.box (192.168.178.82)
    Host is up (0.011s latency).

    PORT   STATE SERVICE
    22/tcp open  ssh

    Nmap scan report for maximilian-laptop.fritz.box (192.168.178.86)
    Host is up (0.00024s latency).

    PORT   STATE SERVICE
    22/tcp open  ssh

    Nmap done: 256 IP addresses (9 hosts up) scanned in 3.00 seconds
    `
    target := map[string][]string{
        "host": []string{
            "raspberrypi",
            "raspberrypi",
            "maximilian1-P6634",
            "maximilian-laptop",
        },
        "ip": []string{
            "192.168.178.24",
            "192.168.178.39",
            "192.168.178.82",
            "192.168.178.86",
        },
    }

    value := Build_Host_Map(nmapOutput)

    if !reflect.DeepEqual(value, target) {
        t.Errorf("target: %v value: %v", target, value)
    }
}
