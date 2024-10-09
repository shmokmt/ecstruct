package ecstruct

type TaskDefinition struct {
	Family          string `json:"family"`
	RuntimePlatform struct {
		OperatingSystemFamily string `json:"operatingSystemFamily"`
	} `json:"runtimePlatform"`
	TaskRoleArn          string                `json:"taskRoleArn"`
	ExecutionRoleArn     string                `json:"executionRoleArn"`
	NetworkMode          string                `json:"networkMode"`
	PlatformFamily       string                `json:"platformFamily"`
	ContainerDefinitions []ContainerDefinition `json:"containerDefinitions"`
	Volumes              []struct {
		Name string `json:"name"`
		Host struct {
			SourcePath string `json:"sourcePath"`
		} `json:"host"`
		ConfiguredAtLaunch        bool `json:"configuredAtLaunch"`
		DockerVolumeConfiguration struct {
			Scope         string `json:"scope"`
			Autoprovision bool   `json:"autoprovision"`
			Driver        string `json:"driver"`
			DriverOpts    struct {
				KeyName string `json:"KeyName"`
			} `json:"driverOpts"`
			Labels struct {
				KeyName string `json:"KeyName"`
			} `json:"labels"`
		} `json:"dockerVolumeConfiguration"`
		EfsVolumeConfiguration struct {
			FileSystemID          string `json:"fileSystemId"`
			RootDirectory         string `json:"rootDirectory"`
			TransitEncryption     string `json:"transitEncryption"`
			TransitEncryptionPort int    `json:"transitEncryptionPort"`
			AuthorizationConfig   struct {
				AccessPointID string `json:"accessPointId"`
				Iam           string `json:"iam"`
			} `json:"authorizationConfig"`
		} `json:"efsVolumeConfiguration"`
	} `json:"volumes"`
	RequiresCompatibilities []string `json:"requiresCompatibilities"`
	CPU                     string   `json:"cpu"`
	Memory                  string   `json:"memory"`
	Tags                    []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"tags"`
	EphemeralStorage struct {
		SizeInGiB int `json:"sizeInGiB"`
	} `json:"ephemeralStorage"`
	PidMode            string `json:"pidMode"`
	IpcMode            string `json:"ipcMode"`
	ProxyConfiguration struct {
		Type          string `json:"type"`
		ContainerName string `json:"containerName"`
		Properties    []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"properties"`
	} `json:"proxyConfiguration"`
	InferenceAccelerators []struct {
		DeviceName string `json:"deviceName"`
		DeviceType string `json:"deviceType"`
	} `json:"inferenceAccelerators"`
}
