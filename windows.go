//go:build windows

package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func GetResourceData() ResourceResult {
	response := ResourceResult{
		usedMemory:  0,
		totalMemory: 0,
		cpu:         0,
	}

	cmd := exec.Command("Get-CimInstance", "Win32_OperatingSystem | % { '{0};{1}' -f $_.TotalVisibleMemorySize, $_.FreePhysicalMemory }")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return response
	}
	stdoutstr := string(stdout)
	split := strings.Split(stdoutstr, ";")
	if len(split) > 0 {
		totalMemory, _ := strconv.ParseUint(split[0], 10, 64)
		response.totalMemory = totalMemory / 100
		freeMemory, _ := strconv.ParseUint(split[1], 10, 64)
		response.usedMemory = (totalMemory - freeMemory) / 100
	}

	cmd2 := exec.Command("Get-CimInstance", "Win32_Processor | % { '{0}' -f $_.LoadPercentage }")
	stdout2, err2 := cmd2.Output()
	if err2 != nil {
		fmt.Println(err2.Error())
		return response
	}
	stdoutstr2 := string(stdout2)
	value, _ := strconv.ParseFloat(stdoutstr2, 64)
	response.cpu = value

	return response
}
