package targetwriter

type Config struct {
	Vault   map[string]*Vault   `json:"vault,omitempty" mapstructure:"vault"`
	Consul  map[string]*Consul  `json:"consul,omitempty" mapstructure:"consul"`
	Nomad   map[string]*Nomad   `json:"nomad,omitempty" mapstructure:"nomad"`
	Default map[string]*Default `json:"default,omitempty" mapstructure:"default"`
}

type Vault struct {
	VaultEndopoint string `json:"endpoint" mapstructure:"endpoint"`
	VaultToken     string `json:"token,omitempty" mastructure:"token"`
	VaultCaPath    string `json:"ca_path,omiempty" mapstructure:"ca_path"`
	VaultCaCert    string `json:"ca_cert,omitempty" mapstructure:"ca_cert"`
	VaultCert      string `json:"cert,omitempty" mapstructure:"cert"`
	VaultKey       string `json:"key,omitempty" mapstructure:"key"`
	VaultFormat    string `json:"format,omitempty" mapstructure:"format"`
	VaultNamespace string `json:"namespace,omitempty" mapstructure:"namespace"`
}

type Consul struct {
	ConsulEndopoint string `json:"endpoint" mapstructure:"endpoint"`
	ConsulToken     string `json:"token,omitempty" mastructure:"token"`
	ConsulCaPath    string `json:"ca_path,omiempty" mapstructure:"ca_path"`
	ConsulCaCert    string `json:"ca_cert,omitempty" mapstructure:"ca_cert"`
	ConsulCert      string `json:"cert,omitempty" mapstructure:"cert"`
	ConsulKey       string `json:"key,omitempty" mapstructure:"key"`
	ConsulTokenFile string `json:"token_file,omitempty" mapstructure:"token_file"`
}

type Nomad struct {
	NomadEndopoint string `json:"endpoint" mapstructure:"endpoint"`
	NomadToken     string `json:"token,omitempty" mastructure:"token"`
	NomadCaPath    string `json:"ca_path,omiempty" mapstructure:"ca_path"`
	NomadCaCert    string `json:"ca_cert,omitempty" mapstructure:"ca_cert"`
	NomadCert      string `json:"cert,omitempty" mapstructure:"cert"`
	NomadKey       string `json:"key,omitempty" mapstructure:"key"`
	NomadRegion    string `json:"region,omitempty" mapstructure:"region"`
	NomadNamespace string `json:"namespace,omitempty" mapstructure:"namespace"`
}

type Default struct {
	VaultProfile    string `json:"vault_profile,omitempty" mapstracture:"vault_profile"`
	NomadProfile    string `json:"nomad_profile,omitempty" mapstracture:"nomad_profile"`
	ConsulProfile   string `json:"consul_profile,omitempty" mapstracture:"consul_profile"`
	BoundaryProfile string `json:"boundary_profile,omitempty" mapstracture:"boundary_profile"`
}
