package goydlidar

import (
	"fmt"
	"testing"
)

func TestDriver(t *testing.T) {
	ldriver := CreateYDdriver()
	err := ldriver.Connect("/dev/ttyUSB1", 128000)
	if err != nil {
		t.Fatal(err)
	}
	/*
		scanfreq, err := ldriver.GetScanFrequency(0)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println("Got scan freq", scanfreq)
		}
		ZeroOffANgle, err := ldriver.GetZeroOffsetAngle(0)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println("Got Zero OFF Angle", ZeroOffANgle)
		}
		samplerate, err := ldriver.GetSamplingRate(0)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println("Got Sampling Rate", samplerate)
		}
	*/
	err = ldriver.StartScan(false, 0)
	if err != nil {
		t.Fatal(err)
	}

	err = ldriver.StartMotor()
	if err != nil {
		t.Fatal(err)
	}
	ldriver.SetAutoReconnect(true)
	if ldriver.IsConnected() {
		fmt.Println("Is Connected")
	} else {
		fmt.Println("Is not connected")
	}
	driverhealth, err := ldriver.GetHealth(0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(driverhealth)
	/*
		scanfreq, err = ldriver.GetScanFrequency(0)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println("Got scan freq", scanfreq)
		}
		ZeroOffANgle, err = ldriver.GetZeroOffsetAngle(0)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println("Got Zero OFF Angle", ZeroOffANgle)
		}
		samplerate, err = ldriver.GetSamplingRate(0)
		if err != nil {
			t.Error(err)
		} else {
			fmt.Println("Got Sampling Rate", samplerate)
		}
	*/
	nodes, err := ldriver.GrabScanData(MaxScanNodes, 0)
	if err != nil {
		t.Error(err)
	} else {
		//	fmt.Println(nodes)
		fmt.Println(len(nodes))
		//	fmt.Println("Got Sampling Rate", samplerate)
		fmt.Println(nodes[0])
		fmt.Println(nodes[1])
		fmt.Println(nodes[2])
		sortednodes, err2 := ldriver.AscendScanData(nodes)
		if err != nil {
			t.Error(err2)
		} else {
			fmt.Println(sortednodes[0])
			fmt.Println(sortednodes[1])
			fmt.Println(sortednodes[2])
		}

	}

	err = ldriver.Stop()
	if err != nil {
		t.Fatal(err)
	}
	err = ldriver.StopMotor()
	if err != nil {
		t.Fatal(err)
	}

}

func TestLidar(t *testing.T) {
	nscans := 10
	lidar := CreateLidar()
	lidar.SetSerialPort("/dev/ttyUSB1")
	lidar.SetFixedResolution(false)
	lidar.SetAutoReconnect(true)
	lidar.SetMaxAngle(180)
	lidar.SetMinAngle(-180)
	lidar.SetMinRange(.001)
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
