package collect

const (
	npuID       = "id"
	modelName   = "model_name"
	npuUUID     = "v_dev_id"
	npuPCIEInfo = "pcie_bus_info"
	namespace   = "namespace"
	podName     = "pod_name"
	cntrName    = "container_name"
)

var cardLabel = []string{npuID, modelName, npuUUID, npuPCIEInfo, namespace, podName, cntrName}
