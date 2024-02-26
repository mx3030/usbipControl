package ui

func (ca *ControlArea) Handle_Control_Buttons() {
    ca.addButton.SetSelectedFunc(ca.Handle_Add_Button)
    ca.disconAllButton.SetSelectedFunc(ca.Handle_DisconAll_Button)
    ca.closeButton.SetSelectedFunc(ca.Handle_Close_Button)
}

func (cA *ControlArea) Handle_Add_Button(){
    cA.parent.conA.Add_Connection()
    cA.parent.Get_Focus_Elements()
}

func (cA *ControlArea) Handle_DisconAll_Button() {
    for _, c := range cA.parent.conA.connections {
        c.ucon.Close_Connection()
    }
}

func (cA *ControlArea) Handle_Close_Button() {
    cA.Handle_DisconAll_Button()
    cA.parent.app.Stop()
}
