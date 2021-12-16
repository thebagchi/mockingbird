package common

import (
	"fmt"
	"os/exec"
	"strings"
)

func Execute(command, arguments string) bool {
	if true {
		status := exec.Command(command, strings.Fields(arguments)...).Run()
		if nil != status {
			fmt.Println("Failed executing command:", command, arguments)
			return false
		}
		return true
	} else {
		output, err := exec.Command(command, strings.Fields(arguments)...).Output()
		if nil != err {
			fmt.Println("Failed executing command:", command, arguments)
			fmt.Println("Output:", string(output))
			return false
		}
	}
	return true
}

func ExecuteEX(command, arguments string) bool {
	if false {
		status := exec.Command(command, strings.Fields(arguments)...).Run()
		if nil != status {
			fmt.Println("Failed executing command:", command, arguments)
			return false
		}
		return true
	} else {
		output, err := exec.Command(command, strings.Fields(arguments)...).CombinedOutput()
		if nil != err {
			fmt.Println("Failed executing command:", command, arguments)
			fmt.Println("Error: ", err)
			fmt.Println("Output:", string(output))
			return false
		}
	}
	return true
}
