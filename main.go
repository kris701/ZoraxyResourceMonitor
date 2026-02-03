package main

import (
	"bufio"
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	plugin "github.com/kris701/zoraxyresourcemonitor/mod/zoraxy_plugin"
)

const (
	PLUGIN_ID = "zoraxyresourcemonitor"
	UI_PATH   = "/"
	WEB_ROOT  = "/www"
	LOG_PATH  = "log.txt"
	// Once a minute
	//LOG_DELAY = 60000
	LOG_DELAY = 1000
	// 1440 log entries
	//LOG_LENGTH = 1440
	LOG_LENGTH = 10
)

//go:embed www/*
var content embed.FS

func main() {
	runtimeCfg, err := plugin.ServeAndRecvSpec(&plugin.IntroSpect{
		ID:            "zoraxyresourcemonitor",
		Name:          "Resource Monitor",
		Author:        "Kristian Skov Johansen",
		AuthorContact: "kris701kj@gmail.com",
		Description:   "A plugin to get simple resource usage of the Zoraxy host",
		URL:           "https://github.com/kris701/zoraxyresourcemonitor",
		Type:          plugin.PluginType_Utilities,
		VersionMajor:  1,
		VersionMinor:  0,
		VersionPatch:  1,
		UIPath:        UI_PATH,
	})
	if err != nil {
		panic(err)
	}

	embedWebRouter := plugin.NewPluginEmbedUIRouter(PLUGIN_ID, &content, WEB_ROOT, UI_PATH)
	embedWebRouter.RegisterTerminateHandler(func() {
		fmt.Println("Resource Monitor Plugin Exited")
	}, nil)

	embedWebRouter.HandleFunc("/api/data", getData, nil)

	// Background saving of resource usage
	go func() {
		for {
			logData()
			time.Sleep(LOG_DELAY * time.Millisecond)
		}
	}()

	http.Handle(UI_PATH, embedWebRouter.Handler())
	fmt.Println("Resource Monitor started at http://127.0.0.1:" + strconv.Itoa(runtimeCfg.Port))
	err = http.ListenAndServe("127.0.0.1:"+strconv.Itoa(runtimeCfg.Port), nil)
	if err != nil {
		panic(err)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	f, err := os.OpenFile(LOG_PATH, os.O_RDONLY, 0777)
	if err != nil {
		http.Error(w, "Could not open log file: "+err.Error(), http.StatusInternalServerError)
	}

	response := map[string](map[string]string){}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ";")
		response[lineSplit[0]] = make(map[string]string)
		response[lineSplit[0]]["usedMemory"] = lineSplit[1]
		response[lineSplit[0]]["totalMemory"] = lineSplit[2]
		response[lineSplit[0]]["cpu"] = lineSplit[3]
	}

	f.Close()

	if err := scanner.Err(); err != nil {
		http.Error(w, "Error reading log file: "+err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func logData() {
	f, err := os.OpenFile(LOG_PATH, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return
	}
	fData, err := os.ReadFile(LOG_PATH)
	if err != nil {
		return
	}
	fStr := string(fData)
	for countRune(fStr, '\n') > LOG_LENGTH {
		index := strings.Index(fStr, "\n") + 1
		fStr = fStr[index:]
	}

	data := GetResourceData()

	fStr += time.Now().Format(time.RFC3339) + ";" + strconv.FormatUint(data.usedMemory, 10) + ";" + strconv.FormatUint(data.totalMemory, 10) + ";" + strconv.FormatFloat(data.cpu, 'f', -1, 64) + "\n"

	f.Truncate(0)
	f.WriteString(fStr)
	f.Close()
}

// https://stackoverflow.com/a/47240301
func countRune(s string, r rune) int {
	count := 0
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count
}
