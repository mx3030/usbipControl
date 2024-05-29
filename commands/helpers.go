package commands

import (
    "os/user"
    "strings"
    "os/exec"
    "fmt"
)

var LocalIp, _ = Get_Local_Ip()
var HomeDir, _ = Get_Home_Dir()
var SSHKeyPath = HomeDir + "/.ssh/id_rsa"

func Get_Local_Ip() (string, error) {
    output, err := Run_Command("hostname", "-I")
    if err != nil {
        return "", fmt.Errorf("failed geting local ip: %w", err)
    }
    ip := strings.Split(output, " ")[0]
    return ip, nil
}

func Get_Home_Dir() (string, error) {
    currentUser, err := user.Current()
    if err != nil {
        return "", fmt.Errorf("failed getting home directory: %w", err)
    }
    return currentUser.HomeDir, nil
}

func Run_Command(command string, args ...string) (string, error) {
    cmd := exec.Command(command, args...)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", fmt.Errorf("failed execute command %s: %w", command, err)
    }
    return string(output), nil
}

func (con *SSHConnection) Run_Command(command string) (string, error) {
     if con == nil {
        return "", fmt.Errorf("SSHConnection is nil")
    }
    if con.Client == nil {
        return "", fmt.Errorf("SSHConnection.Client is nil")
    }
    session, err := con.Client.NewSession()
    if err != nil {
        return "", fmt.Errorf("failed to ssh into %s: %w", con.ServerIp, err)
    }
    defer session.Close()
    output, err := session.CombinedOutput(command)
    if err != nil {
        return "", fmt.Errorf("failed executing command %s on %s: %w", command, con.ServerIp, err)
    }
    return string(output), nil
}
