package ecstruct

type ContainerDefinition struct {
	Name                  string `json:"name"`
	Image                 string `json:"image"`
	RepositoryCredentials struct {
		CredentialsParameter string `json:"credentialsParameter"`
	} `json:"repositoryCredentials"`
	CPU               int      `json:"cpu"`
	Memory            int      `json:"memory"`
	MemoryReservation int      `json:"memoryReservation"`
	Links             []string `json:"links"`
	PortMappings      []struct {
		ContainerPort int    `json:"containerPort"`
		HostPort      int    `json:"hostPort"`
		Protocol      string `json:"protocol"`
	} `json:"portMappings"`
	Essential   bool     `json:"essential"`
	EntryPoint  []string `json:"entryPoint"`
	Command     []string `json:"command"`
	Environment []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"environment"`
	EnvironmentFiles []struct {
		Value string `json:"value"`
		Type  string `json:"type"`
	} `json:"environmentFiles"`
	MountPoints []struct {
		SourceVolume  string `json:"sourceVolume"`
		ContainerPath string `json:"containerPath"`
		ReadOnly      bool   `json:"readOnly"`
	} `json:"mountPoints"`
	VolumesFrom []struct {
		SourceContainer string `json:"sourceContainer"`
		ReadOnly        bool   `json:"readOnly"`
	} `json:"volumesFrom"`
	LinuxParameters struct {
		Capabilities struct {
			Add  []string `json:"add"`
			Drop []string `json:"drop"`
		} `json:"capabilities"`
		Devices []struct {
			HostPath      string   `json:"hostPath"`
			ContainerPath string   `json:"containerPath"`
			Permissions   []string `json:"permissions"`
		} `json:"devices"`
		InitProcessEnabled bool `json:"initProcessEnabled"`
		SharedMemorySize   int  `json:"sharedMemorySize"`
		Tmpfs              []struct {
			ContainerPath string   `json:"containerPath"`
			Size          int      `json:"size"`
			MountOptions  []string `json:"mountOptions"`
		} `json:"tmpfs"`
		MaxSwap    int `json:"maxSwap"`
		Swappiness int `json:"swappiness"`
	} `json:"linuxParameters"`
	Secrets []struct {
		Name      string `json:"name"`
		ValueFrom string `json:"valueFrom"`
	} `json:"secrets"`
	DependsOn []struct {
		ContainerName string `json:"containerName"`
		Condition     string `json:"condition"`
	} `json:"dependsOn"`
	StartTimeout           int      `json:"startTimeout"`
	StopTimeout            int      `json:"stopTimeout"`
	Hostname               string   `json:"hostname"`
	User                   string   `json:"user"`
	WorkingDirectory       string   `json:"workingDirectory"`
	DisableNetworking      bool     `json:"disableNetworking"`
	Privileged             bool     `json:"privileged"`
	ReadonlyRootFilesystem bool     `json:"readonlyRootFilesystem"`
	DNSServers             []string `json:"dnsServers"`
	DNSSearchDomains       []string `json:"dnsSearchDomains"`
	ExtraHosts             []struct {
		Hostname  string `json:"hostname"`
		IPAddress string `json:"ipAddress"`
	} `json:"extraHosts"`
	DockerSecurityOptions []string `json:"dockerSecurityOptions"`
	Interactive           bool     `json:"interactive"`
	PseudoTerminal        bool     `json:"pseudoTerminal"`
	DockerLabels          struct {
		KeyName string `json:"KeyName"`
	} `json:"dockerLabels"`
	Ulimits []struct {
		Name      string `json:"name"`
		SoftLimit int    `json:"softLimit"`
		HardLimit int    `json:"hardLimit"`
	} `json:"ulimits"`
	LogConfiguration struct {
		LogDriver string `json:"logDriver"`
		Options   struct {
			KeyName string `json:"KeyName"`
		} `json:"options"`
		SecretOptions []struct {
			Name      string `json:"name"`
			ValueFrom string `json:"valueFrom"`
		} `json:"secretOptions"`
	} `json:"logConfiguration"`
	HealthCheck struct {
		Command     []string `json:"command"`
		Interval    int      `json:"interval"`
		Timeout     int      `json:"timeout"`
		Retries     int      `json:"retries"`
		StartPeriod int      `json:"startPeriod"`
	} `json:"healthCheck"`
	SystemControls []struct {
		Namespace string `json:"namespace"`
		Value     string `json:"value"`
	} `json:"systemControls"`
	ResourceRequirements []struct {
		Value string `json:"value"`
		Type  string `json:"type"`
	} `json:"resourceRequirements"`
	FirelensConfiguration struct {
		Type    string `json:"type"`
		Options struct {
			KeyName string `json:"KeyName"`
		} `json:"options"`
	} `json:"firelensConfiguration"`
}
