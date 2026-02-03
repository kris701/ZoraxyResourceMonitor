//go:build linux

package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func GetResourceData() ResourceResult {
	response := ResourceResult{
		usedMemory:  0,
		totalMemory: 0,
		cpu:         0,
	}

	cmd := exec.Command("top", "-b", "-n", "1")
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return response
	}
	stdoutstr := string(stdout)

	totalMemoryRegex := regexp.MustCompile(`(?m)MiB Mem :([^.]*)`)
	usedMemoryRegex := regexp.MustCompile(`(?m)MiB Mem :.*free,([^.]*)`)
	cpuRegex := regexp.MustCompile(`(?m)%Cpu\(s\):([^u]*)`)

	totalMemory := totalMemoryRegex.FindStringSubmatch(stdoutstr)
	freeMemory := usedMemoryRegex.FindStringSubmatch(stdoutstr)
	cpu := cpuRegex.FindStringSubmatch(stdoutstr)

	if len(totalMemory) > 0 {
		targetStr := strings.TrimSpace(totalMemory[1])
		value, _ := strconv.ParseUint(targetStr, 10, 64)
		response.totalMemory = value
	}
	if len(freeMemory) > 0 {
		targetStr := strings.TrimSpace(freeMemory[1])
		value, _ := strconv.ParseUint(targetStr, 10, 64)
		response.usedMemory = value
	}
	if len(cpu) > 0 {
		targetStr := strings.TrimSpace(cpu[1])
		value, _ := strconv.ParseFloat(targetStr, 64)
		response.cpu = value
	}

	return response
}
