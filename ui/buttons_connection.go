package ui

import (
    "github.com/gdamore/tcell/v2"
)

func (c *Connection) Handle_Connection_Buttons() {
    c.reloadButton.SetSelectedFunc(func() {
        c.Handle_Dropdowns()
    })

    c.conButton.SetSelectedFunc(func() {
        sourceIndex, _ := c.sourceDropdown.GetCurrentOption()
        deviceIndex, _ := c.deviceDropdown.GetCurrentOption()
        targetIndex, _ := c.targetDropdown.GetCurrentOption()

        sourceIp := c.hostMap["ip"][sourceIndex]
        deviceName := c.deviceMap["name"][deviceIndex]
        deviceId := c.deviceMap["id"][deviceIndex]
        busId := c.deviceMap["busId"][deviceIndex]
        targetIp := c.hostMap["ip"][targetIndex]

        c.ucon.SourceIp = sourceIp
        c.ucon.TargetIp = targetIp
        c.ucon.DeviceName = deviceName
        c.ucon.DeviceId = deviceId
        c.ucon.BusId = busId

        status := c.ucon.Create_Connection()

        if status == true {
            c.disconButton.SetDisabled(false)
            c.reloadButton.SetDisabled(true)
            c.conButton.SetDisabled(true)
            c.sourceDropdown.SetDisabled(true)
            c.deviceDropdown.SetDisabled(true)
            c.targetDropdown.SetDisabled(true)
            c.infoBox.SetFormAttributes(0, tcell.ColorGreen, tcell.ColorGreen, tcell.ColorWhite, tcell.ColorGreen)
            c.Set_Info_Box_Text("connected")
        } else {
            c.infoBox.SetFormAttributes(0, tcell.ColorRed, tcell.ColorRed, tcell.ColorWhite, tcell.ColorRed)
            c.Set_Info_Box_Text("disconnected")
        }

    })

    c.disconButton.SetSelectedFunc(func() {
        status := c.ucon.Close_Connection()
        if status == true {
            c.reloadButton.SetDisabled(false)
            c.conButton.SetDisabled(false)
            c.sourceDropdown.SetDisabled(false)
            c.deviceDropdown.SetDisabled(false)
            c.targetDropdown.SetDisabled(false)
            c.infoBox.SetFormAttributes(0, tcell.ColorRed, tcell.ColorRed, tcell.ColorWhite, tcell.ColorRed)
            c.Set_Info_Box_Text("disconnected")
        }
    })

    c.deleteButton.SetSelectedFunc(func() {
        c.ucon.Close_Connection()
        c.parent.Remove_Connection(c)
    })
}


