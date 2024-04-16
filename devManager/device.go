package devManager

// 对 Nvidia GPU 通过 nvml 获取设备信息
// 对 Ascend NPU 通过 dcmi 获取设备信息
// 此文件意在取代Nvidia的nvml/device.go和Ascend的devmanager.go

//import (
//	"huawei.com/npu-exporter/v6/devmanager/dcmi"
//)

type Device interface {
	DeviceGetCount() int
	DeviceGetName() string
	DeviceGetId() int
	DeviceGetType() string
	DeviceGetIndex() int
	DeviceGetMemory() uint64
	DeviceGetPath() string
	DeviceIsHealthy() (bool, error) // exporter 所需
}

//type DeviceManager struct {
//	// DcMgr for common dev manager
//	DcMgr dcmi.DcDriverInterface
//	// DevType the value is the same as the device type corresponding to the DcMgr variable.
//	// Options: common.Ascend310,common.Ascend310P,common.Ascend910
//	DevType string
//	// ProductTypes product type in server, multi type will be in 310P mix scene
//	ProductTypes []string
//	// isTrainingCard whether the device is used for training
//	isTrainingCard bool
//	dcmiVersion    string
//}
