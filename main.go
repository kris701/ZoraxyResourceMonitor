package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	mem "github.com/pbnjay/memory"
	cpu "github.com/shirou/gopsutil/cpu"

	plugin "github.com/kris701/zoraxyresourcemonitor/mod/zoraxy_plugin"
)

const (
	PLUGIN_ID = "zoraxyresourcemonitor"
	UI_PATH   = "/"
	WEB_ROOT  = "/www"
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
		VersionMajor:  0,
		VersionMinor:  0,
		VersionPatch:  1,
		UIPath:        UI_PATH,
	})
	if err != nil {
		panic(err)
	}

	embedWebRouter := plugin.NewPluginEmbedUIRouter(PLUGIN_ID, &content, WEB_ROOT, UI_PATH)
	embedWebRouter.RegisterTerminateHandler(func() {
		fmt.Println("Top Monitor Plugin Exited")
	}, nil)

	embedWebRouter.HandleFunc("/api/data", getData, nil)

	http.Handle(UI_PATH, embedWebRouter.Handler())
	fmt.Println("Top Monitor started at http://127.0.0.1:" + strconv.Itoa(runtimeCfg.Port))
	err = http.ListenAndServe("127.0.0.1:"+strconv.Itoa(runtimeCfg.Port), nil)
	if err != nil {
		panic(err)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]any{}

	response["freeMemory"] = mem.FreeMemory()
	response["totalMemory"] = mem.TotalMemory()

	response["cpuUsage"], _ = cpu.Percent(0, false)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
