
# Supervisor monitor

This project is designed to monitor and manage supervisor processes from the web interface


## Authors

- [@sadoviy_gnom](https://www.github.com/naginnn)


## Usage/Examples

First, to get started, describe the configuration file config.yaml


```java
supervisor:
  url: http://user:123@localhost:9001/RPC2
  config: /your/path/to/supervisord.conf

server:
  host: 0.0.0.0
  port: 2222
```
All static files and templates are included in the assembly In order to assemble simply
```sh
go build main.go
```
Run
```sh
./main
```
Browse host:port specified in the configuration file

0.0.0.0:2222

## Documentation

monitoring
- supervisor and process groups
- read and clear logs
  control
- stop All/One Process/es
- start All/One Process/es
- restart All/One Process/es
- shutdown PID
- kill -9 PID (supervisor)
  configuration
- add new config
- update config
- reload config


