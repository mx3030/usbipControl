package commands

//-----------------------------------------------------------
type USBIPConnection struct {
    SourceIp    string
    TargetIp    string
    DeviceName  string
    DeviceId    string
    BusId       string
}
//-----------------------------------------------------------

func Get_Remote_USB_Devices(ip string) map[string][]string {
    con, err := Establish_SSHConnection_With_PrivateKey(ip, SSHKeyPath)
    defer con.Close()
    if err != nil {
        return make(map[string][]string)
    }
    deviceMap := con.Get_Local_USB_Devices() 
    return deviceMap
}

func (ucon *USBIPConnection) Create_Connection() bool {
    if ucon.SourceIp == ucon.TargetIp {
        return false
    }
    if ucon.SourceIp == LocalIp {
        Bind_Device(ucon.BusId) 
    } else {
        scon, err := Establish_SSHConnection_With_PrivateKey(ucon.SourceIp, SSHKeyPath)
        defer scon.Close()
        if err != nil {
            return false
        }
        scon.Bind_Device(ucon.BusId)
    }
    if ucon.TargetIp == LocalIp {
        Attach_Device(ucon.SourceIp, ucon.BusId)
    } else {
        tcon, err := Establish_SSHConnection_With_PrivateKey(ucon.TargetIp, SSHKeyPath)
        defer tcon.Close()
        if err != nil {
            return false
        }
        tcon.Attach_Device(ucon.SourceIp, ucon.BusId)
    }
    return true
}

func (ucon *USBIPConnection) Close_Connection() bool {
    if ucon.SourceIp == LocalIp {
        Unbind_Device(ucon.BusId)
    } else {
        scon, err := Establish_SSHConnection_With_PrivateKey(ucon.SourceIp, SSHKeyPath)
        defer scon.Close()
        if err != nil {
            return false
        }
        scon.Unbind_Device(ucon.BusId)
    }
    return true
}

