# Zoraxy Resource Monitor Plugin

This is a simple plugin for [Zoraxy](https://github.com/tobychui/zoraxy) that adds a page where you can see the resource usage (CPU and Memory) over the last 24 hours.
<img width="1899" height="937" alt="image" src="https://github.com/user-attachments/assets/34b4b13b-b591-4240-bea8-b4fc2b987ff6" />

## Plugin

You can now install the Zoraxy plugin itself, by doing the following:

```bash
mkdir -p /opt/zoraxy/plugins/zoraxyresourcemonitor
cd /opt/zoraxy/plugins/zoraxyresourcemonitor
# wget <LINK_TO_LATEST_BINARY>
wget https://github.com/kris701/zoraxyresourcemonitor/releases/download/v1.0.3/zoraxyresourcemonitor
chmod +x zoraxyresourcemonitor
```

Then you can restart your Zoraxy server or service and you should be able to see the new plugin in the sidebar.

# Development

Execute run script `devRun.ps1`.
Then run the server with `./zoraxy -dev=true -noauth=true -port=:8564`
You can then run `devRun.ps1` whenever you want to update the binary.
The script needs WSL installed, and it launches a wsl process for the Zoraxy server.
You can use `devKill.ps1` to kill the server again.

This is a rather rudimentary dev system, if anyone can figure out to set up a propper Docker environment with Zoraxy working, i would greatly appreciate the help :)
