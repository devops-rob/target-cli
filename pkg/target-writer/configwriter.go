package targetwriter

type Config struct {
	Vault   map[string]*Vault   `hcl:"vault,omitempty" mapstructure:"vault"`
	Consul  map[string]*Consul  `hcl:"consul,omitempty" mapstructure:"consul"`
	Nomad   map[string]*Nomad   `hcl:"nomad,omitempty" mapstructure:"nomad"`
	Serf    map[string]*Serf    `hcl:"serf,omitempty" mapstructure:"serf"`
	Default map[string]*Default `hcl:"default,omitempty" mapstructure:"default"`
}

type Vault struct {
	VaultEndopoint string `hcl:"endpoint" mapstructure:"endpoint"`
	VaultToken     string `hcl:"token,omitempty" mastructure:"token"`
	VaultCaPath    string `hcl:"ca_path,omiempty" mapstructure:"ca_path"`
	VaultCaCert    string `hcl:"ca_cert,omitempty" mapstructure:"ca_cert"`
	VaultCert      string `hcl:"cert,omitempty" mapstructure:"cert"`
	VaultKey       string `hcl:"key,omitempty" mapstructure:"key"`
	VaultFormat    string `hcl:"format,omitempty" mapstructure:"format"`
	VaultNamespace string `hcl:"namespace,omitempty" mapstructure:"namespace"`
}

type Consul struct {
	ConsulEndopoint string `hcl:"endpoint" mapstructure:"endpoint"`
	ConsulToken     string `hcl:"token,omitempty" mastructure:"token"`
	ConsulCaPath    string `hcl:"ca_path,omiempty" mapstructure:"ca_path"`
	ConsulCaCert    string `hcl:"ca_cert,omitempty" mapstructure:"ca_cert"`
	ConsulCert      string `hcl:"cert,omitempty" mapstructure:"cert"`
	ConsulKey       string `hcl:"key,omitempty" mapstructure:"key"`
	ConsulTokenFile string `hcl:"token_file,omitempty" mapstructure:"token_file"`
}

type Nomad struct {
	NomadEndopoint string `hcl:"endpoint" mapstructure:"endpoint"`
	NomadToken     string `hcl:"token,omitempty" mastructure:"token"`
	NomadCaPath    string `hcl:"ca_path,omiempty" mapstructure:"ca_path"`
	NomadCaCert    string `hcl:"ca_cert,omitempty" mapstructure:"ca_cert"`
	NomadCert      string `hcl:"cert,omitempty" mapstructure:"cert"`
	NomadKey       string `hcl:"key,omitempty" mapstructure:"key"`
	NomadRegion    string `hcl:"region,omitempty" mapstructure:"region"`
	NomadNamespace string `hcl:"namespace,omitempty" mapstructure:"namespace"`
}

type Serf struct {
	SerfEndopoint string `hcl:"endpoint" mapstructure:"endpoint"`
	SerfToken     string `hcl:"token,omitempty" mastructure:"token"`
}

type Default struct {
	VaultProfile  string `hcl:"vault_profile,omitempty" mapstracture:"vault_profile"`
	NomadProfile  string `hcl:"vault_profile,omitempty" mapstracture:"vault_profile"`
	ConsulProfile string `hcl:"vault_profile,omitempty" mapstracture:"vault_profile"`
	SerfProfile   string `hcl:"vault_profile,omitempty" mapstracture:"vault_profile"`
}
