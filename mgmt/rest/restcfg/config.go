package restcfg

import "os"

// default configuration values
var (
	defaultEnable          = true
	DefaultPort            = 8181
	defaultAddress         = ""
	defaultHTTPS           = false
	defaultRestCertificate = ""
	defaultRestKey         = ""
	DefaultRestLoadPath    = os.TempDir()
	defaultAuth            = false
	defaultAuthPassword    = ""
	defaultPortSetByConfig = false
	defaultPprof           = false
)

// holds the configuration passed in through the SNAP config file
//   Note: if this struct is modified, then the switch statement in the
//         UnmarshalJSON method in this same file needs to be modified to
//         match the field mapping that is defined here
type Config struct {
	Enable           bool   `json:"enable"yaml:"enable"`
	Port             int    `json:"port"yaml:"port"`
	Address          string `json:"addr"yaml:"addr"`
	HTTPS            bool   `json:"https"yaml:"https"`
	RestCertificate  string `json:"rest_certificate"yaml:"rest_certificate"`
	RestKey          string `json:"rest_key"yaml:"rest_key"`
	RestLoadPath     string `json:"rest_load_path"yaml:"rest_load_path"`
	RestAuth         bool   `json:"rest_auth"yaml:"rest_auth"`
	RestAuthPassword string `json:"rest_auth_password"yaml:"rest_auth_password"`
	portSetByConfig  bool   ``
	Pprof            bool   `json:"pprof"yaml:"pprof"`
}

const (
	CONFIG_CONSTRAINTS = `
			"restapi" : {
				"type": ["object", "null"],
				"properties" : {
					"enable": {
						"type": "boolean"
					},
					"https" : {
						"type": "boolean"
					},
					"rest_auth": {
						"type": "boolean"
					},
					"rest_auth_password": {
						"type": "string"
					},
					"rest_certificate": {
						"type": "string"
					},
					"rest_key" : {
						"type": "string"
					},
					"rest_load_path" : {
						"type": "string"
					},
					"port" : {
						"type": "integer",
						"minimum": 1,
						"maximum": 65535
					},
					"addr" : {
						"type": "string"
					},
					"pprof": {
						"type": "boolean"
					}
				},
				"additionalProperties": false
			}
	`
)

// GetDefaultConfig gets the default snapteld configuration
func GetDefaultConfig() *Config {
	return &Config{
		Enable:           defaultEnable,
		Port:             DefaultPort,
		Address:          defaultAddress,
		HTTPS:            defaultHTTPS,
		RestCertificate:  defaultRestCertificate,
		RestKey:          defaultRestKey,
		RestLoadPath:     DefaultRestLoadPath,
		RestAuth:         defaultAuth,
		RestAuthPassword: defaultAuthPassword,
		portSetByConfig:  defaultPortSetByConfig,
		Pprof:            defaultPprof,
	}
}

// define a method that can be used to determine if the port the RESTful
// API is listening on was set in the configuration file
func (c *Config) PortSetByConfigFile() bool {
	return c.portSetByConfig
}
