# cmd-prometheus-exporter

Used to create prometheus guages from any bash command. Most commands that return a single number/float should work. If elevated privileges are required, the process should probably be run as a unique user with limited `sudo` access to only be able to run the desired commands

# Usage

 * create a config.yml in the same directory as the binary file, there is an example in the repository.
 * run the binary
 
# Output

```
 curl http://localhost:8088/metrics                   
# HELP cmd_output Generates gauges from arbitary linux cmds
# TYPE cmd_output gauge
cmd_output{name="line count of rsnapshot log"} 49
cmd_output{name="test"} 1
cmd_output{name="test_2"} 20
```
