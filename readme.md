# Zoraxy Resource Monitor Plugin

Tmp


## Plugin

You can now install the Zoraxy plugin itself, by doing the following:

```bash
mkdir -p /opt/zoraxy/plugins/zoraxyresourcemonitor
cd /opt/zoraxy/plugins/zoraxyresourcemonitor
# wget <LINK_TO_LATEST_BINARY>
wget https://github.com/kris701/zoraxyresourcemonitor/releases/download/v1.0.0/zoraxyresourcemonitor
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
