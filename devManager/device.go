package devManager

// 对 Nvidia GPU 通过 nvml 获取设备信息
// 对 Ascend NPU 通过 dcmi 获取设备信息
// 此文件意在取代Nvidia的nvml/device.go和Ascend的devmanager.go

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
