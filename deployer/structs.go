package deployer

type azureCdnDeployer struct {
	HostName     string
	ProfileName  string
	EndpointName string
	Location     string
}

type azureDeployer struct {
	Location    string
	Count       int
	VMSize      string
	Environment string
	PrivateKey  string
	PublicKey   string
}

type cloudFrontDeployer struct {
	Origin string
	Region string
}

type digitalOceanDeployer struct {
	Image        string
	Fingerprint  string
	PrivateKey   string
	PublicKey    string
	Size         string
	Count        int
	Region       string
	DefaultUser  string
	Name         string
	FirewallName string
}

type googleCloudDeployer struct {
	Region            string
	Project           string
	Count             int
	SSHUser           string
	SSHPubKeyFile     string
	SSHPrivateKeyFile string
	MachineType       string
	Image             string
}

type apiGatewayDeployer struct {
	TargetURI string
	StageName string
}

//Output Parsing Structs

type TerraformOutput struct {
	Master OuterLevel `json:"providers"`
}

type OuterLevel struct {
	ProviderValues Providers `json:"value"`
}

type Providers struct {
	AWSProvider    AWSProvider    `json:"AWS"`
	DOProvider     DOProvider     `json:"DO"`
	GoogleProvider GoogleProvider `json:"GOOGLE"`
	AzureProvider  AzureProvider  `json:"AZURE"`
}

type AWSProvider struct {
	Instances      []AWSInstance      `json:"instances"`
	API            []AWSApi           `json:"api"`
	DomainFront    []AWSDomainFront   `json:"domain_front"`
	SecurityGroups []AWSSecurityGroup `json:"security_group"`
}

type AWSConfigWrapper struct {
	Config    AWSRegionConfig
	RegionMap map[string]int
}

type AWSInstance struct {
	Config  AWSRegionConfig   `json:"config"`
	IPIDMap map[string]string `json:"ip_id"`
}

type AWSRegionConfig struct {
	ModuleName      string
	SecurityGroup   string `json:"hidensneak"`
	SecurityGroupID string `json:"aws_sg_id"`
	Count           string `json:"region_count"`
	CustomAmi       string `json:"custom_ami"`
	InstanceType    string `json:"aws_instance_type"`
	DefaultUser     string `json:"ec2_default_user"`
	Region          string `json:"region"`
	PrivateKeyFile  string `json:"private_key_file"`
	PublicKeyFile   string `json:"public_key_file"`
}

type AWSApi struct {
}

type AWSDomainFront struct{}

type AWSSecurityGroup struct{}

type DOProvider struct {
	Instances []DOInstance `json:"instances"`
}

type DOInstance struct {
	Config  DORegionConfig    `json:"config"`
	IPIDMap map[string]string `json:"ip_id"`
}

type DORegionConfig struct {
}

type GoogleProvider struct {
	Instances []GoogleInstance `json:"instances"`
}

type GoogleInstance struct {
	Config  GoogleRegionConfig `json:"config"`
	IPIDMap map[string]string  `json:"ip_id"`
}

type GoogleRegionConfig struct {
}

type AzureProvider struct {
	Instances   []AzureInstance    `json:"instances"`
	DomainFront []AzureDomainFront `json:"domain_front"`
}

type AzureDomainFront struct{}

type AzureInstance struct {
	Config  AzureRegionConfig `json:"config"`
	IPIDMap map[string]string `json:"ip_id"`
}

type AzureRegionConfig struct{}

//ReadList contains a list of all of the resources
//across different providers per region
type ReadList struct {
	azureCdnDeployerList     []azureCdnDeployer
	azureDeployerList        []azureDeployer
	cloudFrontDeployerList   []cloudFrontDeployer
	digitalOceanDeployerList []digitalOceanDeployer
	googleCloudDeployerList  []googleCloudDeployer
	apiGatewayDeployerList   []apiGatewayDeployer
}
