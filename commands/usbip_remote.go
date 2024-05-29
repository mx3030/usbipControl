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
	output, err := con.Run_Command(command)
    if err != nil {
        return make(map[string][]string)
    }
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

func (con *SSHConnection) Bind_Device(busId string) bool {
    command := "sudo usbip bind -b" + busId
    _, err := con.Run_Command(command)
    if err != nil {
        return false
    }
    return true

}

func (con *SSHConnection) Attach_Device(ipAddress, busId string) bool {
	command := "sudo usbip attach -r " + ipAddress + " -b " + busId
    _, err := con.Run_Command(command)
    if err != nil {
        return false
    }
    return true
}

func (con *SSHConnection) Unbind_Device(busId string) bool {
    command := "sudo usbip unbind -b" + busId
    _, err := con.Run_Command(command)
    if err != nil {
        return false
    }
    return true
}
