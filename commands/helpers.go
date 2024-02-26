package commands

import (
    "os/user"
    "strings"
    "os/exec"
)

var LocalIp = Get_Local_Ip()
var HomeDir = Get_Home_Dir()
var SSHKeyPath = HomeDir + "/.ssh/id_rsa"

func Get_Local_Ip() string {
    output := Run_Command("hostname", "-I")
    ip := strings.Split(output, " ")[0]
    return ip
}

func Get_Home_Dir() string {
    currentUser, _ := user.Current()
    return currentUser.HomeDir
}

func Run_Command(command string, args ...string) string {
    cmd := exec.Command(command, args...)
    output, _ := cmd.CombinedOutput()
    return string(output)
}

func (con *SSHConnection) Run_Command(command string) string {
	session, _ := con.Client.NewSession()
	defer session.Close()
	output, _ := session.CombinedOutput(command)
	return string(output)
}





