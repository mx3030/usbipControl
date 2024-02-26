package commands

import (
	"regexp"
	"strings"
)

func Get_Hosts() map[string][]string {
	ipRange := "192.168.178.0/24"
	nmapOutput := Run_Command("nmap", "-p", "22", "--open", ipRange)
	hostMap := Build_Host_Map(nmapOutput)
    return hostMap
}

func (con *SSHConnection) Get_Hosts() map[string][]string {
	ipRange := "192.168.178.0/24"
	nmapOutput := con.Run_Command("nmap -p 22 --open " + ipRange)
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
    var con *SSHConnection
    for _, line := range lines {
        matches := re.FindStringSubmatch(line)
        if len(matches) == 3 {
            host := strings.Split(matches[1],".")[0]
            ip := matches[2]
            var check bool
            if LocalIp == ip {
                check = Is_USBIP_Available()
            } else {
                con  = Establish_SSHConnection_With_PrivateKey(ip, SSHKeyPath)
                defer con.Close()
                check = con.Is_USBIP_Available()
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
	output := Run_Command("sudo", "which", "usbip")
	if strings.TrimSpace(output) == "" {
		return false
	}
	return true
}

func (con *SSHConnection) Is_USBIP_Available() bool {
	command := "sudo which usbip"
	output := con.Run_Command(command)
    if strings.TrimSpace(output) == "" {
		return false
	}
	return true
}

