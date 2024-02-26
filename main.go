package main

import (
    "usbipControl/ui"
    "usbipControl/utils"
    "fmt"
    "os"
	"os/signal"
	"syscall"
)

func main() {
    enableLogging := false

    err := utils.WithLogging(enableLogging, func() error {
        //-----------------------------------------------------------
        // without logging
        myUI := ui.New_UI()

        //-----------------------------------------------------------
        // Handle Ctrl+C interrupt
        sigint := make(chan os.Signal, 1)
        signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)

        go func() {
            <-sigint
            myUI.Exit()
            os.Exit(0)
        }()
        //-----------------------------------------------------------

        if err := myUI.Run(); err != nil {
            panic(err)
        }
        myUI.Exit()
        //-----------------------------------------------------------
        return nil
    })

    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }

}

