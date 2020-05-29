package goydlidar

import (
	"fmt"
	"runtime"
	"testing"
)

func TestLidar(t *testing.T) {
	nscans := 10
	runtime.LockOSThread()
	lidar := CreateLidar()
	lidar.SetSerialPort(TypicalDevicePort)
	lidar.SetFixedResolution(false)
	lidar.SetAutoReconnect(true)
	lidar.SetMaxAngle(180)
	lidar.SetMinAngle(-180)
	lidar.SetMinRange(.1)
	lidar.SetMaxRange(12)
	lidar.SetIgnoreArray(nil)
	if !lidar.Initialize() {
		t.Fatal("Lidar didn't Initialize")
	} else {
		fmt.Println("Lidar Lnitialized")
	}
	if !lidar.TurnOn() {
		t.Fatal("Lidar didn't turn on")
	} else {
		fmt.Println("Lidar Is On")
	}
	scans := make([]*LidarScan, nscans)
	for i := 0; i < nscans; i++ {
		scans[i] = CreateLidarScan()

		scanned, hardwareerror := lidar.DoProcessSimple(scans[i])
		if scanned {
			fmt.Println("Scan was successful")
			if hardwareerror {
				fmt.Println("Hardware Error")
			}
			go func(i int) {
				ranges, intense, config, selfstamp, systemstamp := scans[i].GetAllValues()
				maxrange := float32(0)
				for i := range ranges {
					if ranges[i] > maxrange {
						maxrange = ranges[i]
					}

				}
				fmt.Println("MaxRange is ", maxrange)
				fmt.Println("Ranges")
				fmt.Println(ranges)
				fmt.Println("intense")
				fmt.Println(intense)
				fmt.Println("config")
				fmt.Println(config)
				fmt.Println("selfstamp")
				fmt.Println(selfstamp)
				fmt.Println("systemstamp")
				fmt.Println(systemstamp)
			}(i)

		} else {
			t.Errorf("Didn't Get Lidar Data")
		}
	}
	lidar.TurnOff()
	lidar.Disconnecting()

}
