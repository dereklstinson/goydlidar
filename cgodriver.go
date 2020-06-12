//Package goydlidar is bindings to the ydlidar driver sdk.  The documentation is in Chinese. So, I have no idea what it says.
package goydlidar

/*
#include "cgoyddriver.h"
#include <stdlib.h>

#cgo CFLAGS: -I/usr/include
#cgo CXXFLAGS:   --std=c++11
//#cgo CXXFLAGS: -I/usr/include/c++/7
#cgo LDFLAGS: -L${SRCDIR} -lydlidar_driver
*/
import "C"
import (
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

//YDdriver is the driver for YDlidar
type YDdriver struct {
	y C.YDLidarCGODriver
}

//DefaultFlags that are used with some functions
const (
	DefaultTimeout      = uint32(C.DEFAULT_TIMEOUT)
	DefaultHeartBeat    = uint32(C.DEFAULT_HEART_BEAT)
	DefaultTimeoutCount = uint32(C.DEFAULT_TIMEOUT_COUNT)
	MaxScanNodes        = uint(C.MAX_SCAN_NODES)
	TypicalX4Baudrate   = uint32(128000)
)

//DeviceHealth is the device health
type DeviceHealth C.device_health

//DeviceInfo is a wrapper for device_info
type DeviceInfo C.device_info

//IsOk is some func that that is used in the sample.
func IsOk() bool {
	return (bool)(C.YDLidarCGOIsOk())
}
func init() {
	//this is the init given
	C.YDLidarCGOInit()
}

//Model returns the model of device
func (d *DeviceInfo) Model() uint8 {
	return (uint8)(d.model)
}

//Firmware returns the firmware version of device
func (d *DeviceInfo) Firmware() uint16 {
	return (uint16)(d.firmware_version)
}

//Hardware returns the hardware version of device
func (d *DeviceInfo) Hardware() uint8 {
	return (uint8)(d.hardware_version)
}

//SerialNumber returns the serial number of device
func (d *DeviceInfo) SerialNumber() []byte {
	serial := make([]byte, len(d.serialnum))
	for i := range serial {
		serial[i] = (byte)(d.serialnum[i])
	}
	return serial
}

//String implements the Stringer interface
func (d *DeviceInfo) String() string {
	return fmt.Sprintf("Model: %v\nFirmware: %v\nHardware: %v\nSerialNumber: %v\n", d.Model(), d.Firmware(), d.Hardware(), d.SerialNumber())
}

//CreateYDdriver creates YDdriver
func CreateYDdriver() *YDdriver {
	y := new(YDdriver)
	y.y = C.YDLidarCGOCreateDriver()
	return y
}

//Connect Connects to device
func (y *YDdriver) Connect(portPath string, baudrate uint32) error {
	cpp := C.CString(portPath)

	return result(C.YDLidarCGOconnect(y.y, cpp, (C.uint32_t)(baudrate))).error()
}

//Disconnect disconnects the driver
func (y *YDdriver) Disconnect() {
	C.YDLidarCGOdisconnect(y.y)
}

//IsScanning checks to see if lidar is scanning
func (y *YDdriver) IsScanning() bool {
	return (bool)(C.YDLidarCGOisscanning(y.y))
}

//IsConnected checks to see if lidar is connected
func (y *YDdriver) IsConnected() bool {
	return (bool)(C.YDLidarCGOisconnected(y.y))
}

//SetIntensities sets intensities
func (y *YDdriver) SetIntensities(val bool) {
	C.YDLidarCGOsetIntensities(y.y, (*C.bool)(&val))
}

//SetAutoReconnect sets the auto reconnect
func (y *YDdriver) SetAutoReconnect(val bool) {
	C.YDLidarCGOsetAutoReconnect(y.y, (*C.bool)(&val))
}

//GetHealth gets the device health.  If timeout is 0 then it does the default
func (y *YDdriver) GetHealth(timeout uint32) (d DeviceHealth, err error) {

	to := (C.uint32_t)(timeout)
	err = result(C.YDLidarCGOgetHealth(y.y, (*C.device_health)(&d), to)).error()
	return d, err
}

//GetDeviceInfo gets the device info.  If timeout is 0 then it does the default
func (y *YDdriver) GetDeviceInfo(timeout uint32) (d DeviceInfo, err error) {

	to := (C.uint32_t)(timeout)
	err = result(C.YDLidarCGOgetDeviceInfo(y.y, (*C.device_info)(&d), to)).error()
	return d, err
}

//StartScan starts scanning device.
func (y *YDdriver) StartScan(force bool, timeout uint32) (err error) {

	to := (C.uint32_t)(timeout)
	err = result(C.YDLidarCGOstartScan(y.y, (C.bool)(force), to)).error()
	return err
}

//Stop stops?
func (y *YDdriver) Stop() error {
	return result(C.YDLidarCGOstop(y.y)).error()
}

//Reset resets?
func (y *YDdriver) Reset(timeout uint32) error {
	return result(C.YDLidarCGOreset(y.y, (C.uint32_t)(timeout))).error()
}

//StartMotor starts the motor
func (y *YDdriver) StartMotor() error {
	return result(C.YDLidarCGOstartMotor(y.y)).error()
}

//StopMotor stops the motor
func (y *YDdriver) StopMotor() error {
	return result(C.YDLidarCGOstopMotor(y.y)).error()
}

//NodeInfo is the node info
type NodeInfo C.node_info

func (n NodeInfo) String() string {
	return fmt.Sprintf("NodeInfo{\n SyncFlag: %d\n SyncQuality: %d\n AngleQ6CheckBit: %d\n DistanceQ2: %d\n TimeStamp: %d\n ScanFreqency: %d\n",
		n.sync_flag, n.sync_quality, n.angle_q6_checkbit, n.distance_q2, n.stamp, n.scan_frequence)
}

//SyncFlag is a value of in the C.node_info struct
func (n NodeInfo) SyncFlag() byte { return (byte)(n.sync_flag) }

//SyncQuality is a value of in the C.node_info struct
func (n NodeInfo) SyncQuality() uint16 { return (uint16)(n.sync_quality) }

//AngleQ6CheckBit is a value of in the C.node_info struct
func (n NodeInfo) AngleQ6CheckBit() uint16 { return (uint16)(n.angle_q6_checkbit) }

//DistanceQ2 is a value of in the C.node_info struct
func (n NodeInfo) DistanceQ2() uint16 { return (uint16)(n.distance_q2) }

//TimeStamp is a value of in the C.node_info struct
func (n NodeInfo) TimeStamp() uint64 { return (uint64)(n.stamp) }

//ScanFreqency is a value of in the C.node_info struct
func (n NodeInfo) ScanFreqency() byte { return (byte)(n.scan_frequence) }

func cnodeinfotoNodeInfo(array []C.node_info) []NodeInfo {
	n := make([]NodeInfo, len(array))
	for i := range n {
		n[i] = (NodeInfo)(array[i])
	}
	return n
}
func nodeInfotoC(array []NodeInfo) []C.node_info {
	n := make([]C.node_info, len(array))
	for i := range n {
		n[i] = (C.node_info)(array[i])
	}
	return n
}

//GrabScanData grabs scan data the size of count.  I don't know if it will grab that exact amount, but original function passes
//count by reference.  If the size of count changes it will change the size of n.
func (y *YDdriver) GrabScanData(count uint, timeout uint32) (n []NodeInfo, err error) {
	ccount := (C.size_t)(count)
	cn := make([]C.node_info, count)
	err = result(C.YDLidarCGOgrabScanData(y.y, &cn[0], &ccount, (C.uint32_t)(timeout))).error()
	n = cnodeinfotoNodeInfo(cn[:ccount])
	return n, err
}

//AscendScanData returns a copy of the data ascended.
func (y *YDdriver) AscendScanData(n []NodeInfo) ([]NodeInfo, error) {
	ccount := (C.size_t)(len(n))
	cn := nodeInfotoC(n)
	err := result(C.YDLidarCGOascendScanData(y.y, &cn[0], ccount)).error()
	return cnodeinfotoNodeInfo(cn), err
}

//Frequency is the frequency
type Frequency C.scan_frequency

func (f *Frequency) cptr() *C.scan_frequency {
	return (*C.scan_frequency)(f)
}

//GetScanFrequency gets the scan frequency
func (y *YDdriver) GetScanFrequency(timeout uint32) (f Frequency, err error) {
	err = result(C.YDLidarCGOgetScanFrequency(y.y, f.cptr(), (C.uint32_t)(timeout))).error()
	return f, err
}

//SetScanFrequencyAdd - Not sure what this does will have to test. But I think it adds 1Hz to the frequency
func (y *YDdriver) SetScanFrequencyAdd(timeout uint32) (f Frequency, err error) {
	err = result(C.YDLidarCGOgetScanFrequency(y.y, f.cptr(), (C.uint32_t)(timeout))).error()
	return f, err
}

//SetScanFrequencyMinus - Not sure what this does will have to test. But I think it minuses 1Hz to the frequency
func (y *YDdriver) SetScanFrequencyMinus(timeout uint32) (f Frequency, err error) {
	err = result(C.YDLidarCGOsetScanFrequencyDis(y.y, f.cptr(), (C.uint32_t)(timeout))).error()
	return f, err
}

//SetScanFrequencyAddMic - Not sure what this does will have to test. But I think it adds .1Hz to the frequency
func (y *YDdriver) SetScanFrequencyAddMic(timeout uint32) (f Frequency, err error) {
	err = result(C.YDLidarCGOsetScanFrequencyAddMic(y.y, f.cptr(), (C.uint32_t)(timeout))).error()
	return f, err
}

//SetScanFrequencyMinusMic - Not sure what this does will have to test. But I think it minuses .1Hz to the frequency
func (y *YDdriver) SetScanFrequencyMinusMic(timeout uint32) (f Frequency, err error) {
	err = result(C.YDLidarCGOsetScanFrequencyDis(y.y, f.cptr(), (C.uint32_t)(timeout))).error()
	return f, err
}

//GetSamplingRate - gets the sampling rate
func (y *YDdriver) GetSamplingRate(timeout uint32) (rate byte, err error) {
	srate := new(C.sampling_rate)
	err = result(C.YDLidarCGOgetSamplingRate(y.y, srate, (C.uint32_t)(timeout))).error()
	rate = (byte)(srate.rate)
	return rate, err
}

//SetSamplingRate sets the sampling rate
func (y *YDdriver) SetSamplingRate(rate byte, timeout uint32) (err error) {
	srate := new(C.sampling_rate)
	srate.rate = (C.uchar)(rate)
	return result(C.YDLidarCGOsetSamplingRate(y.y, srate, (C.uint32_t)(timeout))).error()
}

//GetZeroOffsetAngle gets the zero offset angle
func (y *YDdriver) GetZeroOffsetAngle(timeout uint32) (angle int32, err error) {
	cangle := new(C.offset_angle)
	err = result(C.YDLidarCGOgetZeroOffsetAngle(y.y, cangle, (C.uint32_t)(timeout))).error()
	angle = (int32)(cangle.angle)
	return angle, err

}

/*

Lidar

functions that need to be added are

void CYdCGOgetSerialPort(CYdCGOLidar lidar, char *vals, int nvals, int *actual);

void CYdCGOgetIgnoreArray(CYdCGOLidar lidar, float *vals, int nvals,
                          int *actual);



*/

//Lidar is the lidar
type Lidar struct {
	l C.CYdCGOLidar
}

//CreateLidar creates the lidar
func CreateLidar() *Lidar {
	l := new(Lidar)
	l.l = C.CYdCGOCreateLidar()
	return l
}

//SetSerialPort sets the serial port for the lidar
func (l *Lidar) SetSerialPort(port string) {
	length := (C.int)(len(port))

	cs := C.CString(port)
	C.CYdCGOsetSerialPort(l.l, cs, length)
	C.free(unsafe.Pointer(cs))
}

//SetIgnoreArray sets an array of ignore values?
func (l *Lidar) SetIgnoreArray(ignore []float32) {
	if ignore == nil {
		C.CYdCGOsetIgnoreArray(l.l, nil, C.int(0))
		return
	}
	nvals := (C.int)(len(ignore))
	C.CYdCGOsetIgnoreArray(l.l, (*C.float)(&ignore[0]), nvals)
	return
}

//GetMaxRange returns the max range of the lidar
func (l *Lidar) GetMaxRange() float32 {
	return float32(C.CYdCGOgetMaxRange(l.l))
}

//GetMinRange gets that property for the lidar
func (l *Lidar) GetMinRange() float32 { return (float32)(C.CYdCGOgetMinRange(l.l)) }

//GetMaxAngle gets that property for the lidar
func (l *Lidar) GetMaxAngle() float32 { return (float32)(C.CYdCGOgetMaxAngle(l.l)) }

//GetMinAngle gets that property for the lidar
func (l *Lidar) GetMinAngle() float32 { return (float32)(C.CYdCGOgetMinAngle(l.l)) }

//GetFixedResolution gets that property for the lidar
func (l *Lidar) GetFixedResolution() bool { return (bool)(C.CYdCGOgetFixedResolution(l.l)) }

//GetAutoReconnect gets that property for the lidar
func (l *Lidar) GetAutoReconnect() bool { return (bool)(C.CYdCGOgetAutoReconnect(l.l)) }

//GetSerialBaudrate gets that property for the lidar
func (l *Lidar) GetSerialBaudrate() int32 { return (int32)(C.CYdCGOgetSerialBaudrate(l.l)) }

//GetAbnormalCheckCount gets that property for the lidar
func (l *Lidar) GetAbnormalCheckCount() int32 { return (int32)(C.CYdCGOgetAbnormalCheckCount(l.l)) }

//SetMaxRange sets the max range.
func (l *Lidar) SetMaxRange(val float32) {
	C.CYdCGOsetMaxRange(l.l, (C.float)(val))
}

//SetMinRange sets that property for the lidar
func (l *Lidar) SetMinRange(val float32) { C.CYdCGOsetMinRange(l.l, (C.float)(val)) }

//SetMaxAngle sets that property for the lidar
func (l *Lidar) SetMaxAngle(val float32) { C.CYdCGOsetMaxAngle(l.l, (C.float)(val)) }

//SetMinAngle sets that property for the lidar
func (l *Lidar) SetMinAngle(val float32) { C.CYdCGOsetMinAngle(l.l, (C.float)(val)) }

//SetFixedResolution sets that property for the lidar
func (l *Lidar) SetFixedResolution(val bool) { C.CYdCGOsetFixedResolution(l.l, (C.bool)(val)) }

//SetAutoReconnect sets that property for the lidar
func (l *Lidar) SetAutoReconnect(val bool) { C.CYdCGOsetAutoReconnect(l.l, (C.bool)(val)) }

//SetSerialBaudrate sets that property for the lidar
func (l *Lidar) SetSerialBaudrate(val int32) { C.CYdCGOsetSerialBaudrate(l.l, (C.int)(val)) }

//SetAbnormalCheckCount sets that property for the lidar
func (l *Lidar) SetAbnormalCheckCount(val int32) {
	C.CYdCGOsetAbnormalCheckCount(l.l, (C.int)(val))
}

//Initialize initializes the lidar.
func (l *Lidar) Initialize() bool { return (bool)(C.CYdCGOinitialize(l.l)) }

//DoProcessSimple does a simple process scan.
//Check TestLidar example in yddriver_test.go.
func (l *Lidar) DoProcessSimple(outscan *LidarScan) (funcbool, hardwareerror bool) {
	funcbool = (bool)(C.CYdCGOdoProcessSimple(l.l, outscan.l, (*C.bool)(&hardwareerror)))
	return
}

//TurnOn turns on the lidar
func (l *Lidar) TurnOn() bool { return (bool)(C.CYdCGOturnOn(l.l)) }

//TurnOff turns off the lidar
func (l *Lidar) TurnOff() bool { return (bool)(C.CYdCGOturnOff(l.l)) }

//Disconnecting I guess disconnects
func (l *Lidar) Disconnecting() { C.CYdCGOdisconnecting(l.l) }

/*

Type LidarScan


*/

//LidarScan is the scan from the lidar
type LidarScan struct {
	l           C.CYdCGOLaserScan
	ranges      []float32
	intensities []float32
	config      *LaserConfig
}

//CreateLidarScan creates a lidar scan
func CreateLidarScan() *LidarScan {
	l := new(LidarScan)
	C.CreateCYdCGOLaserScan(&l.l)
	runtime.SetFinalizer(l, destroyLidarScan)
	return l
}
func destroyLidarScan(l *LidarScan) {
	C.DestroyCYdCGOLaserScan(l.l)
}

//LaserConfig is the laserconfig
type LaserConfig C.LaserConfig

func (l *LaserConfig) String() string {
	return fmt.Sprintf("LaserConfig{\nMinAngle: %v,\nMaxAngle: %v,\n"+
		"AngIncrement: %v,\nTimeIncrement: %v,\n"+
		"ScanTime: %v,\nMinRange: %v,\n"+
		"MaxRange: %v,\nRangeRes: %v}\n", l.min_angle, l.max_angle, l.ang_increment, l.time_increment, l.scan_time, l.min_range, l.max_range, l.range_res)
}
func (l *LaserConfig) cptr() *C.LaserConfig {
	return (*C.LaserConfig)(l)
}

//MinAngle returns the minangle
func (l *LaserConfig) MinAngle() float32 {
	return (float32)(l.min_angle)
}

//MaxAngle returns the max_angle
func (l *LaserConfig) MaxAngle() float32 {
	return (float32)(l.max_angle)
}

//AngIncrement returns the ang_increment
func (l *LaserConfig) AngIncrement() float32 {
	return (float32)(l.ang_increment)
}

//TimeIncrement returns the time_increment
func (l *LaserConfig) TimeIncrement() float32 {
	return (float32)(l.time_increment)
}

//ScanTime returns the scan_time
func (l *LaserConfig) ScanTime() float32 {
	return (float32)(l.scan_time)
}

//MinRange returns the min_range
func (l *LaserConfig) MinRange() float32 {
	return (float32)(l.min_range)
}

//MaxRange returns the max_range
func (l *LaserConfig) MaxRange() float32 {
	return (float32)(l.max_range)
}

//RangeRes returns the range_res
func (l *LaserConfig) RangeRes() float32 {
	return (float32)(l.range_res)
}

//NewLidarOutScan is a helper func to create a new lidar scan.
func NewLidarOutScan() *LidarScan {
	return new(LidarScan)
}

//GetNRanges gets the size of the ranges
func (l *LidarScan) getNRanges() int32 {
	return (int32)(C.CYdCGOLaserScanNranges(l.l))
}

//GetNIntensities gets the NIntensities
func (l *LidarScan) getNIntensities() int32 {
	return (int32)(C.CYdCGOLaserScannNintensities(l.l))
}

//GetAllValues gets all the hidden values.
func (l *LidarScan) GetAllValues() (ranges, intensities []float32, config *LaserConfig, selfstamp, systemstamp uint64) {
	if len(l.ranges) < 1 {
		nvals := l.getNRanges()
		if nvals < 1 {
			return nil, nil, nil, 0, 0
		}
		l.ranges = make([]float32, nvals)
		l.intensities = make([]float32, nvals)
	}
	if l.config == nil {
		l.config = new(LaserConfig)
	}
	nvals := (C.int)(len(l.ranges))
	var actual C.int
	C.CYdCGOLaserScanGetAll(l.l,
		(*C.float)(&l.ranges[0]),
		(*C.float)(&l.intensities[0]),
		nvals, &actual,
		l.config.cptr(),
		(*C.ulonglong)(&selfstamp),
		(*C.ulonglong)(&systemstamp))
	if nvals != actual {
		l.ranges = make([]float32, actual)
		l.intensities = make([]float32, actual)
		return l.GetAllValues()
	}
	ranges = l.ranges
	intensities = l.intensities
	config = l.config
	return ranges, intensities, config, selfstamp, systemstamp
}

//GetRanges gets the ranges
func (l *LidarScan) GetRanges() (ranges []float32) {
	if len(l.ranges) < 1 {
		nranges := l.getNRanges()
		if nranges < 1 {
			return nil
		}
		l.ranges = make([]float32, nranges)
	}
	nranges := (C.int)(len(l.ranges))
	var actual C.int
	C.CYdCGOLaserScanGetRanges(l.l, (*C.float)(&l.ranges[0]), nranges, &actual)
	if nranges != actual {
		l.ranges = make([]float32, actual)
		return l.GetRanges()
	}
	ranges = l.ranges
	return ranges
}

/*
//GetIntensities gets the intensities
func (l *LidarScan) GetIntensities() (intensities []float32) {
	if len(l.intensities) < 1 {
		nintensities := l.getNIntensities()
		if nintensities < 1 {
			return nil
		}
		l.intensities = make([]float32, nintensities)
	}
	nintensities := (C.int)(len(l.intensities))
	var actual C.int
	C.CYdCGOLaserScanGetIntensities(l.l, (*C.float)(&l.intensities[0]), nintensities, &actual)
	if nintensities != actual {
		l.intensities = make([]float32, actual)
		return l.GetIntensities()
	}
	intensities = l.intensities
	return intensities
}

*/

type result C.result_t

func (r result) error() error {
	switch r {
	case 0:
		return nil
	case -1:
		return errors.New("ydlidar - TimeOut")
	case -2:
		return errors.New("ydlidar - Fail")
	}
	return errors.New("ydlidar - Fail Uknown Error FLag")
}
