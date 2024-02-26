package commands

import (
	"golang.org/x/crypto/ssh"
	"os"
)

//-----------------------------------------------------------
type SSHConnection struct {
	Client *ssh.Client
	Config *ssh.ClientConfig
}
//-----------------------------------------------------------

func Establish_SSHConnection_With_Password(ip, username, password string) *SSHConnection {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, _ := ssh.Dial("tcp", ip + ":22", config)
	return &SSHConnection{
		Client: client,
		Config: config,
	}
}

func Establish_SSHConnection_With_PrivateKey(ip, privateKeyPath string) *SSHConnection {
    auth := Get_SSH_Auth(privateKeyPath)
    username := SSHConfig[ip]
    config := &ssh.ClientConfig{
        User: username,
        Auth: []ssh.AuthMethod{auth},
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }
    client, _ := ssh.Dial("tcp", ip + ":22", config)
    return &SSHConnection{
        Client: client,
        Config: config,
    }
}

func Get_SSH_Auth(file string) ssh.AuthMethod {
    buffer, _ := os.ReadFile(file)
    key, _ := ssh.ParsePrivateKey(buffer)
    return ssh.PublicKeys(key)
}

func (con *SSHConnection) Close() {
	con.Client.Close()
}
