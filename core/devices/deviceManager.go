package devices

import (
	"fmt"
	"github.com/rebelit/gome-core/common/config"
	"github.com/rebelit/gome-core/core/devices/roku"
	"log"
	"os"
)

func InitializeDatabases() {
	dbPath := config.App.DbPath
	if !requiredPathExists(dbPath) {
		err := createRequiredPath(dbPath)
		if err != nil {
			log.Fatalf("FATAL: Unable to create required database directory %s", dbPath)
		}
	}

	err := roku.InitializeDb()
	if err != nil {
		log.Fatal("FATAL: unable to initialize required gome-core database for roku")
	}

	return
}

func getLoadedDeviceTypes() (types []string, error error) {
	dbs, err := getDbFileNames(config.App.DbPath)
	if err != nil {
		return nil, err
	}

	return dbs, nil
}

func getAllLoadedDevices(typeFilter string) (devices []Device, error error) {
	switch typeFilter {
	case  "roku", "all":
		rokus, err := roku.GetAllDevicesFromDb()
		if err != nil {
			log.Printf("ERROR: unable to get roku devices")
		}

		for _, r := range rokus {
			d := Device{
				Name: r.Name,
				Type: "roku",
				Addr: r.Address,
				Port: r.Port,
			}
			devices = append(devices, d)
		}
	//ToDo: range over other device types. May have to
	//		split these out into other helper functions later
	//		depending how many device types det added.
	//case "rpiot", "all":
	//
	default:
		return devices, fmt.Errorf("invalid device type")
	}

	return devices, nil
}

//private functions
func requiredPathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func createRequiredPath(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}

	return nil
}

func getDbFileNames(path string) (dbs []string, error error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	files, err := dir.Readdir(-1)
	dir.Close()
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir(){
			dbs = append(dbs, file.Name())
		}
	}

	return dbs, nil
}
