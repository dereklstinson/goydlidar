#ifndef __YDDRIVER_FOR_GO__
#define __YDDRIVER_FOR_GO__

#include <stdbool.h>
#include <stdlib.h>
#ifdef __cplusplus
#include "CYdLidar.h"
#include "ydlidar_driver.h"
extern "C" {
#endif
#ifdef __cplusplus

#endif
#ifdef __cplusplus
typedef ydlidar::YDlidarDriver *YDLidarCGODriver;
typedef result_t result_t;
typedef uint32_t uint32_t;
typedef device_health device_health;
typedef device_info device_info;
typedef node_info node_info;
typedef CYdLidar *CYdCGOLidar;
typedef LaserScan *CYdCGOLaserScan;
// typedef scan_frequence
#else
typedef void *CYdCGOLaserScan;
typedef void *YDLidarCGODriver;
typedef void *CYdCGOLidar;
typedef int result_t;
typedef unsigned long int uint32_t;

typedef struct {
  unsigned char status;
  unsigned short int error_code;
} device_health;
typedef struct {
  unsigned char model;
  unsigned short int firmware_version;
  unsigned char hardware_version;
  unsigned char serialnum[16];
} device_info;
typedef struct {
  unsigned char sync_flag;
  unsigned short int sync_quality;
  unsigned short int angle_q6_checkbit;
  unsigned short int distance_q2;
  unsigned long long int stamp;
  unsigned char scan_frequence;
} node_info;
typedef struct {
  unsigned long frequency;
} scan_frequency;
typedef struct {
  unsigned char rate;
} sampling_rate;
typedef struct {
  long int angle;
} offset_angle;
typedef struct {
  //! Start angle for the laser scan [rad].  0 is forward and angles are
  //! measured clockwise when viewing YDLIDAR from the top.
  float min_angle;
  //! Stop angle for the laser scan [rad].   0 is forward and angles are
  //! measured clockwise when viewing YDLIDAR from the top.
  float max_angle;
  //! Scan resolution [rad].
  float ang_increment;
  //! Scan resoltuion [s]
  float time_increment;
  //! Time between scans
  float scan_time;
  //! Minimum range [m]
  float min_range;
  //! Maximum range [m]
  float max_range;
  //! Range Resolution [m]
  float range_res;
} LaserConfig;

// typedef struct {
//  //! Array of ranges
//  float *ranges;
//  //! Array of intensities
//  float *intensities;
//  //! Self reported time stamp in nanoseconds
//  unsigned long long int self_time_stamp;
//  //! System time when first range was measured in nanoseconds
//  unsigned long long int system_time_stamp;
//  //! Configuration of scan
//  LaserConfig config;
//} LaserScan;
#endif
enum {
  DEFAULT_TIMEOUT = 2000,    /**< 默认超时时间. */
  DEFAULT_HEART_BEAT = 1000, /**< 默认检测掉电功能时间. */
  MAX_SCAN_NODES = 2048,     /**< 最大扫描点数. */
  DEFAULT_TIMEOUT_COUNT = 1,
};

typedef enum {
  YDLIDAR_F4 = 1,
  YDLIDAR_T1 = 2,
  YDLIDAR_F2 = 3,
  YDLIDAR_S4 = 4,
  YDLIDAR_G4 = 5,
  YDLIDAR_X4 = 6,
  YDLIDAR_G4PRO = 7,
  YDLIDAR_F4PRO = 8,
  YDLIDAR_G2_SS_1 = 9, // 230400
  YDLIDAR_G10 = 10,    // 256000
  YDLIDAR_S4B = 11,    // 153600
  YDLIDAR_S2 = 12,     // 115200
  YDLIDAR_G25 = 13,    // 512000
  YDLIDAR_Tail,
} YDLidarCGOModel;

typedef enum {
  YDLIDAR_RATE_4K = 0,
  YDLIDAR_RATE_8K = 1,
  YDLIDAR_RATE_9K = 2,
  YDLIDAR_RATE_10K = 3,
} YDLidarCGOScanRate;

// IsOk I guess checks if everything is ok.
void YDLidarCGOInit();
bool YDLidarCGOIsOk();

YDLidarCGODriver YDLidarCGOCreateDriver();

result_t YDLidarCGOconnect(YDLidarCGODriver driver, const char *port_path,
                           uint32_t baudrate);

void YDLidarCGOdisconnect(YDLidarCGODriver driver);

bool YDLidarCGOisscanning(YDLidarCGODriver driver);

bool YDLidarCGOisconnected(YDLidarCGODriver driver);

void YDLidarCGOsetIntensities(YDLidarCGODriver driver,
                              const bool *isintensities);

void YDLidarCGOsetAutoReconnect(YDLidarCGODriver driver, const bool *enable);

result_t YDLidarCGOgetHealth(YDLidarCGODriver driver, device_health *health,
                             uint32_t timeout);

result_t YDLidarCGOgetDeviceInfo(YDLidarCGODriver driver, device_info *info,
                                 uint32_t timeout);

result_t YDLidarCGOstartScan(YDLidarCGODriver driver, bool force,
                             uint32_t timeout);

result_t YDLidarCGOstop(YDLidarCGODriver driver);

result_t YDLidarCGOstartMotor(YDLidarCGODriver driver);

result_t YDLidarCGOstopMotor(YDLidarCGODriver driver);

result_t YDLidarCGOgrabScanData(YDLidarCGODriver driver, node_info *nodebuffer,
                                size_t *count, uint32_t timeout);

result_t YDLidarCGOascendScanData(YDLidarCGODriver driver,
                                  node_info *nodebuffer, size_t count);

result_t YDLidarCGOreset(YDLidarCGODriver driver, uint32_t timeout);

result_t YDLidarCGOgetScanFrequency(YDLidarCGODriver driver,
                                    scan_frequency *frequency,
                                    uint32_t timeout);

result_t YDLidarCGOsetScanFrequencyAdd(YDLidarCGODriver driver,
                                       scan_frequency *frequency,
                                       uint32_t timeout);

result_t YDLidarCGOsetScanFrequencyDis(YDLidarCGODriver driver,
                                       scan_frequency *frequency,
                                       uint32_t timeout);

result_t YDLidarCGOsetScanFrequencyAddMic(YDLidarCGODriver driver,
                                          scan_frequency *frequency,
                                          uint32_t timeout);

result_t YDLidarCGOsetScanFrequencyDisMic(YDLidarCGODriver driver,
                                          scan_frequency *frequency,
                                          uint32_t timeout);

result_t YDLidarCGOgetSamplingRate(YDLidarCGODriver driver, sampling_rate *rate,
                                   uint32_t timeout);

result_t YDLidarCGOsetSamplingRate(YDLidarCGODriver driver, sampling_rate *rate,
                                   uint32_t timeout);

result_t YDLidarCGOgetZeroOffsetAngle(YDLidarCGODriver driver,
                                      offset_angle *angle, uint32_t timeout);
void YDLidarCGOlidarPortList(char **list, int *size);
// static std::string getSDKVersion();
// static std::map<std::string, std::string> lidarPortList();

/*

CYdCGOLidar functions

*/
CYdCGOLidar CYdCGOCreateLidar();
float CYdCGOgetMaxRange(CYdCGOLidar lidar);
float CYdCGOgetMinRange(CYdCGOLidar lidar);
float CYdCGOgetMaxAngle(CYdCGOLidar lidar);
float CYdCGOgetMinAngle(CYdCGOLidar lidar);
bool CYdCGOgetFixedResolution(CYdCGOLidar lidar);
bool CYdCGOgetAutoReconnect(CYdCGOLidar lidar);
int CYdCGOgetSerialBaudrate(CYdCGOLidar lidar);
int CYdCGOgetAbnormalCheckCount(CYdCGOLidar lidar);

void CYdCGOsetMaxRange(CYdCGOLidar lidar, float val);
void CYdCGOsetMinRange(CYdCGOLidar lidar, float val);
void CYdCGOsetMaxAngle(CYdCGOLidar lidar, float val);
void CYdCGOsetMinAngle(CYdCGOLidar lidar, float val);
void CYdCGOsetFixedResolution(CYdCGOLidar lidar, bool val);
void CYdCGOsetAutoReconnect(CYdCGOLidar lidar, bool val);
void CYdCGOsetSerialBaudrate(CYdCGOLidar lidar, int val);
void CYdCGOsetAbnormalCheckCount(CYdCGOLidar lidar, int val);

void CYdCGOgetSerialPort(CYdCGOLidar lidar, char *vals, int nvals, int *actual);
void CYdCGOsetSerialPort(CYdCGOLidar lidar, char *vals, int nvals);
void CYdCGOgetIgnoreArray(CYdCGOLidar lidar, float *vals, int nvals,
                          int *actual);
void CYdCGOsetIgnoreArray(CYdCGOLidar lidar, float *vals, int nvals);

bool CYdCGOinitialize(CYdCGOLidar lidar);
bool CYdCGOdoProcessSimple(CYdCGOLidar lidar, CYdCGOLaserScan outscan,
                           bool *hardwareError);
bool CYdCGOturnOn(CYdCGOLidar lidar);
bool CYdCGOturnOff(CYdCGOLidar lidar);
void CYdCGOdisconnecting(CYdCGOLidar lidar);

/*
CYdCGOLaserScan functions
*/
void CreateCYdCGOLaserScan(CYdCGOLaserScan *outscan);
void DestroyCYdCGOLaserScan(CYdCGOLaserScan outscan);
int CYdCGOLaserScanNranges(CYdCGOLaserScan outscan);
int CYdCGOLaserScannNintensities(CYdCGOLaserScan outscan);
void CYdCGOLaserScanGetRanges(CYdCGOLaserScan outscan, float *ranges,
                              int nranges, int *actual);
void CYdCGOLaserScanGetIntensities(CYdCGOLaserScan outscan, float *intensities,
                                   int nintensities, int *actual);
unsigned long long int CYdCGOLaserScanGetSelfTimeStamp(CYdCGOLaserScan outscan);
unsigned long long int
CYdCGOLaserScanGetSystemTimeStamp(CYdCGOLaserScan outscan);
void CYdCGOLaserScanGetAll(CYdCGOLaserScan outscan, float *ranges,
                           float *intensities, int nvals, int *actual,
                           LaserConfig *config,
                           unsigned long long int *self_time_stamp,
                           unsigned long long int *system_time_stamp);
LaserConfig CYdCGOLaserScanGetLaserConfig(CYdCGOLaserScan outscan);

#ifdef __cplusplus
}
#endif

#endif //__YDDRIVER_FOR_GO__