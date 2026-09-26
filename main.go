package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Theme = string

const (
	Dark  Theme = "dark"
	Light Theme = "light"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("choose theme: dark or light")
		os.Exit(1)
	}

	theme := os.Args[1]

	switch theme {
	case Dark:
		cmd := exec.Command("gsettings", "set", "org.gnome.desktop.interface", "color-scheme", "prefer-dark")
		if err := cmd.Run(); err != nil {
			fmt.Printf("failed to set %s: %v", theme, err)
		}
	case Light:
		cmd := exec.Command("gsettings", "set", "org.gnome.desktop.interface", "color-scheme", "prefer-light")
		if err := cmd.Run(); err != nil {
			fmt.Printf("failed to set %s: %v", theme, err)
		}
	default:
		fmt.Println("unknown theme, choose: dark or light")
		os.Exit(1)
	}

}
