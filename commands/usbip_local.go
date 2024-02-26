package commands

import (
	"strings"
)

func Get_Local_USB_Devices() map[string][]string {
    deviceMap := map[string][]string{
        "busId": make([]string, 0),
        "name":   make([]string, 0),
        "id":   make([]string, 0),
    }
    output := Run_Command("sudo", "usbip","list","-l")
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

func Bind_Device(busId string) {
	Run_Command("sudo", "usbip", "bind", "-b", busId)
}

func Attach_Device(ipAddress, busId string) {
	Run_Command("sudo", "usbip", "attach", "-r", ipAddress, "-b", busId)
}

func Unbind_Device(busId string) {
    Run_Command("sudo", "usbip", "unbind", "-b", busId)
}
