package ui

import (
    "github.com/rivo/tview"
    "github.com/gdamore/tcell/v2"
)

//-----------------------------------------------------------
type UI struct {
    app                     *tview.Application
    flex                    *tview.Flex
    conA                    *ConnectionsArea
    cA                      *ControlArea
    focusElements           []Focusable
    focusIndex              int
}

type Focusable interface {
    Get_Primitive() tview.Primitive
    Handle_Arrow_Keys(ui *UI, event *tcell.EventKey)
}
//-----------------------------------------------------------

func New_UI() *UI {
    app := tview.NewApplication()

    flex := tview.NewFlex().SetDirection(tview.FlexRow)
    conA := New_Connections_Area()
    cA := New_Control_Area()
    title := tview.NewTextView().SetText("USBIP-CONTROL").SetTextAlign(tview.AlignCenter)
    gap := tview.NewTextView()

    flex.AddItem(gap, 1, 1, false)
    flex.AddItem(title, 1, 1, false)
    flex.AddItem(gap, 1, 1, false)
    flex.AddItem(cA.flex, 1, 1, false)
    flex.AddItem(gap, 1, 1, false)
    flex.AddItem(conA.flex, 0, 1, false)
    flex.AddItem(gap, 1, 1, false)

    ui := &UI{
        app:                    app,
        flex:                   flex,
        conA:                   conA,
        cA:                     cA,
        focusElements:          nil,
        focusIndex:             0,
    }

    cA.parent = ui

    return ui
}

func (ui *UI) Get_Focus_Elements() {
    ui.focusElements = []Focusable{ui.cA}
    for _, con := range ui.conA.connections {
        ui.focusElements = append(ui.focusElements, con)
    }
}

func (ui *UI) Handle_Arrow_Keys(event *tcell.EventKey) *tcell.EventKey {
    switch event.Key() {
    case tcell.KeyTab:
        ui.focusIndex = (ui.focusIndex + 1) % len(ui.focusElements)
    }
    ui.app.SetFocus(ui.focusElements[ui.focusIndex].Get_Primitive())
    ui.focusElements[ui.focusIndex].Handle_Arrow_Keys(ui, event)
    return event
}

func (ui *UI) Run() error {
    ui.Get_Focus_Elements()
    ui.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        return ui.Handle_Arrow_Keys(event)
    })
    err := ui.app.SetRoot(ui.flex, true).SetFocus(ui.cA.addButton).Run()
    if err != nil {
        return err
    }
    return nil
}

func (ui *UI) Exit() {
    ui.cA.Handle_Close_Button()
}

