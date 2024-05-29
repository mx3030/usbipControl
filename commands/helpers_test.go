package commands

import (
    "testing"
)

func Test_Get_Local_Ip(t *testing.T) {
    value, _ := Get_Local_Ip()
    target := "192.168.178.86"
    if value != target {
        t.Errorf("target: %s value: %s", target, value)
    }
}

func Test_Get_Home_Dir(t *testing.T) {
    value, _ := Get_Home_Dir()
    target := "/home/maximilian"
    if value != target {
        t.Errorf("target: %s value: %s", target, value)
    }
}

