package ui

import (
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"
    "usbipControl/commands"
    "strings"
    "fmt"
)

//-----------------------------------------------------------
type Connection struct {
    flex                *tview.Flex
    sourceDropdown      *tview.DropDown
    deviceDropdown      *tview.DropDown
    targetDropdown      *tview.DropDown
    conButton           *tview.Button
    disconButton        *tview.Button
    reloadButton        *tview.Button
    deleteButton        *tview.Button
    infoBox             *tview.TextArea
    infoBoxWidth        int
    focusElements       []tview.Primitive
    focusIndex          int
    hostMap             map[string][]string
    deviceMap           map[string][]string
    ucon                *commands.USBIPConnection
    parent              *ConnectionsArea
    id                  int64
}

//-----------------------------------------------------------

func New_Connection() *Connection {
    flex := tview.NewFlex()

    sourceDropdown := tview.NewDropDown().
    SetLabel("Source: ")
    deviceDropdown := tview.NewDropDown().
    SetLabel("Device: ")
    targetDropdown := tview.NewDropDown().
    SetLabel("Target: ")

    reloadSymbol := "\u27F3"
    reloadButton := tview.NewButton(reloadSymbol)
    conSymbol := "\u2713"
    conButton := tview.NewButton(conSymbol)
    disconSymbol := "\u2717"
    disconButton := tview.NewButton(disconSymbol).SetDisabled(true)
    deleteButton := tview.NewButton("DEL")

    infoBoxWidth := 20
    infoBox := tview.NewTextArea()
    infoBox.SetFormAttributes(0, tcell.ColorRed, tcell.ColorRed, tcell.ColorWhite, tcell.ColorRed)

    gap := tview.NewTextView()

    flex.AddItem(gap, 1, 1, false)
    flex.AddItem(sourceDropdown, 0, 1, false)
    flex.AddItem(gap, 5, 1, false)
    flex.AddItem(deviceDropdown, 0, 1, false)
    flex.AddItem(gap, 5, 1, false)
    flex.AddItem(targetDropdown, 0, 1, false)
    flex.AddItem(gap, 5, 1, false)
    flex.AddItem(conButton, 5, 1, false)
    flex.AddItem(gap, 1, 1, false)
    flex.AddItem(disconButton, 5, 1, false)
    flex.AddItem(gap, 3, 1, false)
    flex.AddItem(reloadButton, 5, 1, false)
    flex.AddItem(gap, 1, 1, false)
    flex.AddItem(deleteButton, 5, 1, false)
    flex.AddItem(gap, 3, 1, false)
    flex.AddItem(infoBox, infoBoxWidth, 1, false)
    flex.AddItem(gap, 1, 1, false)

    focusElements := []tview.Primitive{
        sourceDropdown,
        deviceDropdown,
        targetDropdown,
        conButton,
        disconButton,
        reloadButton,
        deleteButton,
    }

    c := &Connection{
        flex:               flex,
        sourceDropdown:     sourceDropdown,
        deviceDropdown:     deviceDropdown,
        targetDropdown:     targetDropdown,
        conButton:          conButton,
        disconButton:       disconButton,
        reloadButton:       reloadButton,
        deleteButton:       deleteButton,
        infoBox:            infoBox,
        infoBoxWidth:       infoBoxWidth,
        focusElements:      focusElements,
        hostMap:            nil,
        deviceMap:          nil,
        ucon:               &commands.USBIPConnection{},
        parent:             nil,
        id:                 0,
    }

    c.Set_Info_Box_Text("disconnected")
    c.Handle_Dropdowns()
    c.Handle_Connection_Buttons()

    return c
}

func (c *Connection) Set_Info_Box_Text (text string) {
    width := c.infoBoxWidth
    textLen := len(text)
	var centeredText string
    if textLen >= width {
		centeredText = text
	}
	spaces := (width - textLen) / 2
	padding := strings.Repeat(" ", spaces)
	centeredText = padding + text
    c.infoBox.SetText(centeredText, false)
}

func (c *Connection) Get_Primitive() tview.Primitive {
    return c.flex
}

func (c *Connection) Handle_Arrow_Keys(ui *UI, event *tcell.EventKey) {
    switch event.Key() {
        case tcell.KeyRight:
            c.focusIndex = (c.focusIndex + 1) % len(c.focusElements)
        case tcell.KeyLeft:
            c.focusIndex = (c.focusIndex - 1 + len(c.focusElements)) % len(c.focusElements)
    }
    ui.app.SetFocus(c.focusElements[c.focusIndex])
}

func (c *Connection) Handle_Dropdowns() {
    c.hostMap = commands.Get_Hosts()
    fmt.Println(c.hostMap)
    c.sourceDropdown.SetOptions(c.hostMap["host"], nil)
    c.targetDropdown.SetOptions(c.hostMap["host"], nil)

    c.sourceDropdown.SetSelectedFunc(func(host string, index int) {
        ip := c.hostMap["ip"][index]
        if ip == commands.LocalIp {
            c.deviceMap = commands.Get_Local_USB_Devices()
        } else {
            c.deviceMap = commands.Get_Remote_USB_Devices(ip)
        }
        c.deviceDropdown.SetOptions(c.deviceMap["name"], nil)
    })
}

