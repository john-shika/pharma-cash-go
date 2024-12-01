package nokocore

import (
	"net/url"
	"strconv"
	"time"
)

type NetworkConfigImpl interface {
	GetScheme() string
	GetProtocol() string
	GetAddress() string
	GetPort() uint16
	GetHost() string
	GetURL() *url.URL
	String() string
}

type NetworkConfig struct {
	Scheme   string `mapstructure:"scheme" json:"scheme" yaml:"scheme"`
	Protocol string `mapstructure:"protocol" json:"protocol" yaml:"protocol"`
	Address  string `mapstructure:"address" json:"address" yaml:"address"`
	Port     uint16 `mapstructure:"port" json:"port" yaml:"port"`
}

func NewNetworkConfig() NetworkConfigImpl {
	return &NetworkConfig{
		Scheme:   "http",
		Protocol: "tcp",
		Address:  "localhost",
		Port:     80,
	}
}

func (NetworkConfig) GetNameType() string {
	return "network"
}

func (n *NetworkConfig) GetScheme() string {
	return n.Scheme
}

func (n *NetworkConfig) GetProtocol() string {
	return n.Protocol
}

func (n *NetworkConfig) GetAddress() string {
	return n.Address
}

func (n *NetworkConfig) GetPort() uint16 {
	return n.Port
}

func (n *NetworkConfig) GetHost() string {
	return n.Address + ":" + strconv.Itoa(int(n.Port))
}

func (n *NetworkConfig) GetURL() *url.URL {
	return &url.URL{
		Scheme: n.Scheme,
		Host:   n.GetHost(),
	}
}

func (n *NetworkConfig) String() string {
	return n.GetURL().String()
}

func (n *NetworkConfig) WaitForHttpAlive(iterations int, duration time.Duration) {
	TryFetchUrlWaitForAlive(n.GetURL(), iterations, duration)
}
