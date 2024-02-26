package commands

import (
	"strings"
)

func (con *SSHConnection) Get_Local_USB_Devices() map[string][]string {
	deviceMap := map[string][]string{
        "busId": make([]string, 0),
        "name":   make([]string, 0),
        "id":   make([]string, 0),
    }
    command := "sudo usbip list -l"
	output := con.Run_Command(command)
    lines := strings.Split(output, "\n")
	for i := 0; i < len(lines)-1; i += 3 {
		busId := Get_BusId(lines[i])
        deviceMap["busId"] = append(deviceMap["busId"], busId)
        name, id := Get_Device_Info(lines[i+1])
        deviceMap["name"] = append(deviceMap["name"], name)
        deviceMap["id"] = append(deviceMap["id"], id)
	}
	return deviceMap
}

func (con *SSHConnection) Bind_Device(busId string) {
	command := "sudo usbip bind -b" + busId
	con.Run_Command(command)
}

func (con *SSHConnection) Attach_Device(ipAddress, busId string) {
	command := "sudo usbip attach -r " + ipAddress + " -b " + busId
	con.Run_Command(command)
}

func (con *SSHConnection) Unbind_Device(busId string) {
	command := "sudo usbip unbind -b" + busId
	con.Run_Command(command)
}

