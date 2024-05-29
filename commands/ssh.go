package commands

import (
	"golang.org/x/crypto/ssh"
	"os"
    "fmt"
)

//-----------------------------------------------------------
type SSHConnection struct {
	Client *ssh.Client
	Config *ssh.ClientConfig
    ServerIp string 
}
//-----------------------------------------------------------

func Establish_SSHConnection_With_Password(ip, username, password string) (*SSHConnection, error) {
    config := &ssh.ClientConfig{
        User: username,
        Auth: []ssh.AuthMethod{
            ssh.Password(password),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    client, err := ssh.Dial("tcp", ip+":22", config)
    if err != nil {
        return nil, fmt.Errorf("failed to establish SSH connection for IP %s with username %s and password: %w", ip, username, err)
    }
    return &SSHConnection{
        Client: client,
        Config: config,
        ServerIp: ip,
    }, nil
}

func Establish_SSHConnection_With_PrivateKey(ip, privateKeyPath string) (*SSHConnection, error) {
    auth, err := Get_SSH_Auth(privateKeyPath)
    if err != nil {
        return nil, fmt.Errorf("failed to establish SSH connection with private key for IP %s: %w", ip, err)
    }
    username := SSHConfig[ip]
    config := &ssh.ClientConfig{
        User: username,
        Auth: []ssh.AuthMethod{auth},
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    client, err := ssh.Dial("tcp", ip+":22", config)
    if err != nil {
        return nil, fmt.Errorf("failed to establish SSH connection for IP %s: %w", ip, err)
    }
    return &SSHConnection{
        Client: client,
        Config: config,
        ServerIp: ip,
    }, nil
}

func Get_SSH_Auth(file string) (ssh.AuthMethod, error) {
    buffer, err := os.ReadFile(file)
    if err != nil {
        return nil, fmt.Errorf("failed to read private key file %s: %w", file, err)
    }
    key, err := ssh.ParsePrivateKey(buffer)
    if err != nil {
        return nil, fmt.Errorf("failed to parse SSH private key: %w", err)
    }
    return ssh.PublicKeys(key), nil
}

func (con *SSHConnection) Close() error {
    if con == nil {
        return fmt.Errorf("SSHConnection is nil")
    }
    if con.Client == nil {
        return fmt.Errorf("SSHConnection.Client is nil")
    }
    return con.Client.Close()
}

