# gome-core
GoLang Home Automation - Core Services

## WHY?
The original gome (Golang Home) was too bloated and became unmanageable.  This is splitting it out in to 
micro-services.  Start by installing the core service, add the gome-ui, gome-schedule, more, etc... as services
are developed. 

Start it up, add devices (kept in local database), hit the api's to perform actions against loaded devices.

## Application Configuration
### Defaults
* name: `"gome-core"`
* statsd: `""`
* dbPath: `"badgerDatabase"`
* slackWebhook: `""`
* authToken: `"changeMePlease"`
* port: `"6660"`

### Flags
* Override any of the defaults with args
    * `-name "myApp"  -statsd "127.0.0.1:8125" -dbPath "./myDatabases" -slackWebhook "https://hooks.slack.com/services/bunchOfChars/moreKeyChars" -authToken "ThisIsMyNewP@ssw0Rd" -port "8080"`

## Installation
### Docker
* **NOTE** if you change the default port you need to change it in the command below and pass in the `-port` flag
1. `docker build -t gome-core:latest .`
2. `docker run -it --rm -p 6660:6660 -v $PWD:/gome-core gome-core` - Run with defaults
1. `docker run -it --rm -p 6660:6660 -v $PWD:/gome-core gome-core -authToken "abc"` - Override defaults
1. `docker run -it --rm -p 8080:8080 -v $PWD:/gome-core gome-core -port "8080"` - Run on different port

### Manual
1. build it yourself with `go build` and any args you need 
2. execute it

## Supported Devices
[x] roku
[] rpIoT (my [raspberryPi IoT](https://github.com/RebelIT/rpIoT) manager webservice)
[] tuya devices
[] ecoVacs robot vacuum 
[] plex

## Usage: 
* `/api/status`
    * **Method**: GET
    * **Purpose**: Checks the status of the web api
    * **Returns**: Status code

* `/api/deviceTypes`
    * **Method**: GET
    * **Purpose**: Gets what device types are available (loaded in the database)
    * **Returns**: Array of device types loaded
    
* `/api/device/{type}`
    * **Method**: GET
    * **Purpose**: Gets the devices in the database by type
    * **Returns**: Array of devices
    
* `/api/device/{type}`
    * **Method**: POST
    * **Purpose**: Adds a new device by type (DISCLAIMER, currently hard coded for `roku` type)
    * **Returns**: Status code
    
* `/api/roku/{name}/info`
    * **Method**: GET
    * **Purpose**: Gets roku device information
    * **Returns**: Device information
    
* `/api/roku/{name}/online`
    * **Method**: GET
    * **Purpose**: Gets the device online status (network connected)
    * **Returns**: Online state
    
* `/api/roku/{name}/power`
    * **Method**: GET
    * **Purpose**: Gets the device power status (powerOn | PowerOff)
    * **Returns**: Power state
    
* `/api/roku/{name}/power/{state}`
    * **Method**: PUT
    * **Purpose**: Sets a new power state for the device (on | off)
    * **Returns**: Power state
    
* `/api/roku/{name}/app`
    * **Method**: GET
    * **Purpose**: Gets the installed roku apps
    * **Returns**: Array of apps name & id
    
* `/api/roku/{name}/app/active`
    * **Method**: GET
    * **Purpose**: Gets the active app
    * **Returns**: App that is currently launched.  Empty json if in sleep mode.
    
* `/api/roku/{name}/app/launch/{id}`
    * **Method**: POST
    * **Purpose**: Launches an app by id
    * **Returns**: Status code
    
* `/api/roku/{name}/key/{key}`
    * **Method**: GET
    * **Purpose**: Generic keypress (same keys as the remote)
    * **Returns**: Status code
    
## QuickStart Playground
1. git clone `git@github.com:RebelIT/gome-core.git`
2. `docker build -t gome-core:latest .`
1. `docker run -it --rm -p 6660:6660 -v $PWD:/gome-core gome-core`
1. `curl -i http://localhost:6660/api/status`
1. `curl -i -H "Content-Type: application/json" -H "Authorization: Bearer changeMePlease" -X POST -d '{"name": "basement","address": "192.168.1.10","port": "8060"}' http://localhost:6660/api/device/roku`
1. `curl -i -H "Content-Type: application/json" -H "Authorization: Bearer changeMePlease" http://localhost:6660/api/device/roku`
1. `curl -i -H "Content-Type: application/json" -H "Authorization: Bearer changeMePlease" -X POST http://localhost:6660/api/roku/basement/key/home`
