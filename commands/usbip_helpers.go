package commands

import (
    "strings"
)

func Get_BusId(line string) string {
	temp := strings.Split(line, " ")
	return temp[3]
}

func Get_Device_Info(line string) (string, string) {
	temp := strings.Split(line, " ")
	name := strings.Join(temp[:len(temp)-1], " ")
    id := strings.Trim(temp[len(temp)-1], "()")
    return name, id
}

