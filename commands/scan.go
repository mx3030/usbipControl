package commands

import (
	"regexp"
	"strings"
)

func Get_Hosts() map[string][]string {
	ipRange := "192.168.178.0/24"
	nmapOutput, err := Run_Command("nmap", "-p", "22", "--open", ipRange)
    if err != nil {
        return make(map[string][]string)
    }
    hostMap := Build_Host_Map(nmapOutput)
    return hostMap
}

func (con *SSHConnection) Get_Hosts() map[string][]string {
	ipRange := "192.168.178.0/24"
	nmapOutput, err := con.Run_Command("nmap -p 22 --open " + ipRange)
    if err != nil {
        return make(map[string][]string)
    }
	hostMap := Build_Host_Map(nmapOutput)
    return hostMap
}

func Build_Host_Map(nmapOutput string) map[string][]string {
    hostMap := map[string][]string{
        "host": make([]string, 0),
        "ip":   make([]string, 0),
    }
    re, _ := regexp.Compile((`Nmap scan report for ([^\s]+) \(([\d\.]+)\)`))
    lines := strings.Split(nmapOutput, "\n")
    for _, line := range lines {
        matches := re.FindStringSubmatch(line)
        if len(matches) == 3 {
            host := strings.Split(matches[1],".")[0]
            ip := matches[2]
            var check bool
            if LocalIp == ip {
                check = Is_USBIP_Available()
            } else {
                con, err := Establish_SSHConnection_With_PrivateKey(ip, SSHKeyPath)
                defer con.Close()
                if err != nil {
                    check = false
                } else {
                    check = con.Is_USBIP_Available()
                }
            }
            if check == true {
                hostMap["host"] = append(hostMap["host"], host)
                hostMap["ip"] = append(hostMap["ip"], ip)
            }
        }
    }
    return hostMap
}

func Is_USBIP_Available() bool {
	output, err := Run_Command("sudo", "which", "usbip")
    if err != nil {
        return false
    }
	if strings.TrimSpace(output) == "" {
		return false
	}
	return true
}

func (con *SSHConnection) Is_USBIP_Available() bool {
	command := "sudo which usbip"
    output, err := con.Run_Command(command)
    if err != nil {
        return false
    }
    if strings.TrimSpace(output) == "" {
		return false
	}
	return true
}

