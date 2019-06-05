package metal

import (
	"fmt"
	"strings"
)

// A MacAddress is the type for mac adresses. When using a
// custom type, we cannot use strings directly.
type MacAddress string

// Nic information.
type Nic struct {
	MacAddress MacAddress `rethinkdb:"macAddress"`
	Name       string     `rethinkdb:"name"`
	Vrf        string     `rethinkdb:"vrf"`
	Neighbors  Nics       `rethinkdb:"neighbors"`
}

// Nics is a list of nics.
type Nics []Nic

// Prefix is a ip with mask, either ipv4/ipv6
type Prefix struct {
	IP     string `rethinkdb:"ip"`
	Length string `rethinkdb:"length"`
}

// Prefixes is an array of prefixes
type Prefixes []Prefix

// NewPrefixFromCIDR returns a new prefix from a given cidr.
func NewPrefixFromCIDR(cidr string) (*Prefix, error) {
	parts := strings.Split(cidr, "/")
	if len(parts) != 2 {
		return nil, fmt.Errorf("cannot split cidr into pieces: %v", cidr)
	}
	ip := strings.TrimSpace(parts[0])
	length := strings.TrimSpace(parts[1])
	return &Prefix{
		IP:     ip,
		Length: length,
	}, nil
}

// String implements the Stringer interface
func (p *Prefix) String() string {
	return p.IP + "/" + p.Length
}

func (p Prefixes) String() []string {
	result := []string{}
	for _, element := range p {
		result = append(result, element.String())
	}
	return result
}

// Equals returns true when prefixes have the same cidr.
func (p *Prefix) Equals(other *Prefix) bool {
	return p.String() == other.String()
}

// Network is a network in a metal as a service infrastructure, routable between.
// TODO specify rethinkdb restrictions.
type Network struct {
	Base
	Prefixes            Prefixes `rethinkdb:"prefixes"`
	DestinationPrefixes Prefixes `rethinkdb:"destinationprefixes"`
	PartitionID         string   `rethinkdb:"partitionid"`
	ProjectID           string   `rethinkdb:"projectid"`
	TenantID            string   `rethinkdb:"tenantid"`
	ParentNetworkID     string   `rethinkdb:"parentnetworkid"`
	Vrf                 uint     `rethinkdb:"vrf"`
	Primary             bool     `rethinkdb:"primary"`
	Nat                 bool     `rethinkdb:"nat"`
	Underlay            bool     `rethinkdb:"underlay"`
}

// FindPrefix returns the prefix by cidr if contained in this network, nil otherwise
func (n *Network) FindPrefix(cidr string) *Prefix {
	var found *Prefix
	for _, p := range n.Prefixes {
		if p.String() == cidr {
			return &p
		}
	}
	return found
}

// SubstractPrefixes returns the prefixes of the network minus the prefixes passed in the arguments
func (n *Network) SubstractPrefixes(prefixes ...Prefix) []Prefix {
	var result []Prefix
	for _, p := range n.Prefixes {
		contains := false
		for _, prefixToSubstract := range prefixes {
			if p.Equals(&prefixToSubstract) {
				contains = true
				break
			}
		}
		if contains {
			continue
		}
		result = append(result, p)
	}
	return result
}

type NicMap map[MacAddress]*Nic

// ByMac creates a indexed map from a nic list.
func (nics Nics) ByMac() NicMap {
	res := make(map[MacAddress]*Nic)
	for i, n := range nics {
		res[n.MacAddress] = &nics[i]
	}
	return res
}
