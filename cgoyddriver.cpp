#include "cgoyddriver.h"
void YDLidarCGOInit() { ydlidar::init(0, NULL); }
bool YDLidarCGOIsOk() { return ydlidar::ok(); }
YDLidarCGODriver YDLidarCGOCreateDriver() { return new ydlidar::YDlidarDriver; }
result_t YDLidarCGOconnect(YDLidarCGODriver driver, const char *port_path,
                           uint32_t baudrate) {
  return driver->connect(port_path, baudrate);
}

void YDLidarCGOdisconnect(YDLidarCGODriver driver) { driver->disconnect(); }
bool YDLidarCGOisscanning(YDLidarCGODriver driver) {
  return driver->isscanning();
}
bool YDLidarCGOisconnected(YDLidarCGODriver driver) {
  return driver->isconnected();
}
void YDLidarCGOsetIntensities(YDLidarCGODriver driver,
                              const bool *isintensities) {
  // void setIntensities(const bool &isintensities);
  driver->setIntensities(isintensities);
}
void YDLidarCGOsetAutoReconnect(YDLidarCGODriver driver, const bool *enable) {
  // void setAutoReconnect(const bool &enable);
  driver->setAutoReconnect(enable);
}
result_t YDLidarCGOgetHealth(YDLidarCGODriver driver, device_health *health,
                             uint32_t timeout) {
  // result_t getHealth(device_health &health, uint32_t timeout =
  // DEFAULT_TIMEOUT);
  if (timeout == 0) {
    return driver->getHealth(*health);
  }
  return driver->getHealth(*health, timeout);
}

result_t YDLidarCGOgetDeviceInfo(YDLidarCGODriver driver, device_info *info,
                                 uint32_t timeout) {
  // result_t getDeviceInfo(device_info &info, uint32_t timeout =
  // DEFAULT_TIMEOUT);
  if (timeout == 0) {
    return driver->getDeviceInfo(*info);
  }
  return driver->getDeviceInfo(*info, timeout);
}

result_t YDLidarCGOstartScan(YDLidarCGODriver driver, bool force,
                             uint32_t timeout) {
  // result_t startScan(bool force = false, uint32_t timeout = DEFAULT_TIMEOUT)
  // ;
  if (timeout == 0) {
    return driver->startScan(force);
  }
  return driver->startScan(force, timeout);
}

result_t YDLidarCGOstop(YDLidarCGODriver driver) {
  // result_t stop()
  return driver->stop();
};
result_t YDLidarCGOstartMotor(YDLidarCGODriver driver) {
  // result_t startMotor();
  return driver->startMotor();
}

result_t YDLidarCGOstopMotor(YDLidarCGODriver driver) {
  // result_t stopMotor();
  return driver->stopMotor();
}

result_t YDLidarCGOgrabScanData(YDLidarCGODriver driver, node_info *nodebuffer,
                                size_t *count, uint32_t timeout) {
  // result_t grabScanData(node_info *nodebuffer, size_t &count,uint32_t timeout
  // = DEFAULT_TIMEOUT) ;
  if (timeout == 0) {
    return driver->grabScanData(nodebuffer, *count);
  }
  return driver->grabScanData(nodebuffer, *count, timeout);
}

result_t YDLidarCGOascendScanData(YDLidarCGODriver driver,
                                  node_info *nodebuffer, size_t count) {
  // result_t ascendScanData(node_info *nodebuffer, size_t count);

  return driver->ascendScanData(nodebuffer, count);
}

result_t YDLidarCGOreset(YDLidarCGODriver driver, uint32_t timeout) {
  // result_t reset(uint32_t timeout = DEFAULT_TIMEOUT);
  if (timeout == 0) {
    return driver->reset();
  }
  return driver->reset(timeout);
}

result_t YDLidarCGOgetScanFrequency(YDLidarCGODriver driver,
                                    scan_frequency *frequency,
                                    uint32_t timeout) {
  if (timeout == 0) {
    return driver->getScanFrequency(*frequency);
  }
  return driver->getScanFrequency(*frequency, timeout);
}

result_t YDLidarCGOsetScanFrequencyAdd(YDLidarCGODriver driver,
                                       scan_frequency *frequency,
                                       uint32_t timeout) {
  if (timeout == 0) {
    return driver->setScanFrequencyAdd(*frequency);
  }
  return driver->setScanFrequencyAdd(*frequency, timeout);
}

result_t YDLidarCGOsetScanFrequencyDis(YDLidarCGODriver driver,
                                       scan_frequency *frequency,
                                       uint32_t timeout) {
  if (timeout == 0) {
    return driver->setScanFrequencyDis(*frequency);
  }
  return driver->setScanFrequencyDis(*frequency, timeout);
}

result_t YDLidarCGOsetScanFrequencyAddMic(YDLidarCGODriver driver,
                                          scan_frequency *frequency,
                                          uint32_t timeout) {

  if (timeout == 0) {
    return driver->setScanFrequencyAddMic(*frequency);
  }
  return driver->setScanFrequencyAddMic(*frequency, timeout);
}

result_t YDLidarCGOsetScanFrequencyDisMic(YDLidarCGODriver driver,
                                          scan_frequency *frequency,
                                          uint32_t timeout) {

  if (timeout == 0) {
    return driver->setScanFrequencyDisMic(*frequency);
  }
  return driver->setScanFrequencyDisMic(*frequency, timeout);
}

result_t YDLidarCGOgetSamplingRate(YDLidarCGODriver driver, sampling_rate *rate,
                                   uint32_t timeout) {

  if (timeout == 0) {
    return driver->getSamplingRate(*rate);
  }
  return driver->getSamplingRate(*rate, timeout);
}

result_t YDLidarCGOsetSamplingRate(YDLidarCGODriver driver, sampling_rate *rate,
                                   uint32_t timeout) {

  if (timeout == 0) {
    return driver->setSamplingRate(*rate);
  }
  return driver->setSamplingRate(*rate, timeout);
}

result_t YDLidarCGOgetZeroOffsetAngle(YDLidarCGODriver driver,
                                      offset_angle *angle, uint32_t timeout) {
  // result_t getZeroOffsetAngle(offset_angle &angle,uint32_t timeout =
  // DEFAULT_TIMEOUT);
  if (timeout == 0) {
    return driver->getZeroOffsetAngle(*angle);
  }
  return driver->getZeroOffsetAngle(*angle, timeout);
}
// void YDLidarCGOlidarPortList(char **list, int *size) {}

/*
CYdCGOLidar
*/
CYdCGOLidar CYdCGOCreateLidar() { return new CYdLidar; }
float CYdCGOgetMaxRange(CYdCGOLidar lidar) { return lidar->getMaxRange(); }
float CYdCGOgetMinRange(CYdCGOLidar lidar) { return lidar->getMinRange(); }
float CYdCGOgetMaxAngle(CYdCGOLidar lidar) { return lidar->getMaxAngle(); }
float CYdCGOgetMinAngle(CYdCGOLidar lidar) { return lidar->getMinAngle(); }
bool CYdCGOgetFixedResolution(CYdCGOLidar lidar) {
  return lidar->getFixedResolution();
}
bool CYdCGOgetAutoReconnect(CYdCGOLidar lidar) {
  return lidar->getAutoReconnect();
}
int CYdCGOgetSerialBaudrate(CYdCGOLidar lidar) {
  return lidar->getSerialBaudrate();
}
int CYdCGOgetAbnormalCheckCount(CYdCGOLidar lidar) {
  return lidar->getAbnormalCheckCount();
}

void CYdCGOgetSerialPort(CYdCGOLidar lidar, char *vals, int nvals,
                         int *actual) {
  std::string s = lidar->getSerialPort();
  *actual = s.size();
  if (nvals < *actual) {
    for (int i = 0; i < nvals; i++) {
      vals[i] = s[i];
    }
    return;
  }
  for (int i = 0; i < *actual; i++) {
    vals[i] = s[i];
  }
  return;
}
void CYdCGOgetIgnoreArray(CYdCGOLidar lidar, float *vals, int nvals,
                          int *actual) {
  std::vector<float> s = lidar->getIgnoreArray();
  *actual = s.size();
  if (nvals < *actual) {
    for (int i = 0; i < nvals; i++) {
      vals[i] = s[i];
    }
    return;
  }
  for (int i = 0; i < *actual; i++) {
    vals[i] = s[i];
  }
  return;
}
void CYdCGOsetIgnoreArray(CYdCGOLidar lidar, float *vals, int nvals) {
  std::vector<float> s;
  if (nvals == 0) {
    s.clear();
    lidar->setIgnoreArray(s);
    return;
  }
  for (int i = 0; i < nvals; i++) {
    s.push_back(vals[i]);
  }
  lidar->setIgnoreArray(s);
  return;
}
void CYdCGOsetSerialPort(CYdCGOLidar lidar, char *vals, int nvals) {
  std::string s;
  for (int i = 0; i < nvals; i++) {
    s += vals[i];
  }
  lidar->setSerialPort(s);

  return;
}

void CYdCGOsetMaxRange(CYdCGOLidar lidar, float val) {
  lidar->setMaxRange(val);
}
void CYdCGOsetMinRange(CYdCGOLidar lidar, float val) {
  lidar->setMinRange(val);
}
void CYdCGOsetMaxAngle(CYdCGOLidar lidar, float val) {
  lidar->setMaxAngle(val);
}
void CYdCGOsetMinAngle(CYdCGOLidar lidar, float val) {
  lidar->setMinAngle(val);
}
void CYdCGOsetFixedResolution(CYdCGOLidar lidar, bool val) {
  lidar->setFixedResolution(val);
}
void CYdCGOsetAutoReconnect(CYdCGOLidar lidar, bool val) {
  lidar->setAutoReconnect(val);
}
void CYdCGOsetSerialBaudrate(CYdCGOLidar lidar, int val) {
  lidar->setSerialBaudrate(val);
}
void CYdCGOsetAbnormalCheckCount(CYdCGOLidar lidar, int val) {
  lidar->setAbnormalCheckCount(val);
}

bool CYdCGOinitialize(CYdCGOLidar lidar) { return lidar->initialize(); }
bool CYdCGOdoProcessSimple(CYdCGOLidar lidar, CYdCGOLaserScan outscan,
                           bool *hardwareError) {
  return lidar->doProcessSimple((*outscan), *hardwareError);
}
bool CYdCGOturnOn(CYdCGOLidar lidar) { return lidar->turnOn(); }
bool CYdCGOturnOff(CYdCGOLidar lidar) { return lidar->turnOff(); }
void CYdCGOdisconnecting(CYdCGOLidar lidar) { lidar->disconnecting(); }

/*

CYdCGOLaserScan

*/
void CreateCYdCGOLaserScan(CYdCGOLaserScan *outscan) {
  *outscan = new LaserScan;
}
void DestroyCYdCGOLaserScan(CYdCGOLaserScan outscan) { delete outscan; }
int CYdCGOLaserScanNranges(CYdCGOLaserScan outscan) {
  return outscan->ranges.size();
}
int CYdCGOLaserScannNintensities(CYdCGOLaserScan outscan) {
  return outscan->intensities.size();
}
void CYdCGOLaserScanGetRanges(CYdCGOLaserScan outscan, float *ranges,
                              int nranges, int *actual) {
  *actual = (outscan->ranges.size());
  if (nranges < *actual) {
    for (int i = 0; i < nranges; i++) {
      ranges[i] = outscan->ranges[i];
    }
    return;
  }
  for (int i = 0; i < *actual; i++) {
    ranges[i] = outscan->ranges[i];
  }
  return;
}
void CYdCGOLaserScanGetIntensities(CYdCGOLaserScan outscan, float *intensities,
                                   int nintensities, int *actual) {
  *actual = (outscan->intensities.size());
  if (nintensities < *actual) {
    for (int i = 0; i < nintensities; i++) {
      intensities[i] = outscan->intensities[i];
    }
    return;
  }

  for (int i = 0; i < *actual; i++) {
    intensities[i] = outscan->intensities[i];
  }
  return;
}
unsigned long long int
CYdCGOLaserScanGetSelfTimeStamp(CYdCGOLaserScan outscan) {
  return outscan->self_time_stamp;
}
unsigned long long int
CYdCGOLaserScanGetSystemTimeStamp(CYdCGOLaserScan outscan) {
  return outscan->system_time_stamp;
}
LaserConfig CYdCGOLaserScanGetLaserConfig(CYdCGOLaserScan outscan) {
  return outscan->config;
}
void CYdCGOLaserScanGetAll(CYdCGOLaserScan outscan, float *ranges,
                           float *intensities, int nvals, int *actual,
                           LaserConfig *config,
                           unsigned long long int *self_time_stamp,
                           unsigned long long int *system_time_stamp) {

  *self_time_stamp = outscan->self_time_stamp;
  *system_time_stamp = outscan->system_time_stamp;
  *actual = (outscan->ranges.size());
  *config = outscan->config;
  if (nvals < *actual) {
    for (int i = 0; i < nvals; i++) {
      ranges[i] = outscan->ranges[i];
      intensities[i] = outscan->intensities[i];
    }
    return;
  }
  for (int i = 0; i < *actual; i++) {
    ranges[i] = outscan->ranges[i];
    intensities[i] = outscan->intensities[i];
  }
  return;
}