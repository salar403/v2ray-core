package net

func (n Network) AsList() *NetworkList {
	return &NetworkList{
		Network: []Network{n},
	}
}

func (n Network) SystemString() string {
	switch n {
	case Network_TCP:
		return "tcp"
	case Network_UDP:
		return "udp"
	default:
		return "unknown"
	}
}

func HasNetwork(list []Network, network Network) bool {
	for _, value := range list {
		if string(value) == string(network) {
			return true
		}
	}
	return false
}

// HasNetwork returns true if the given network is in v NetworkList.
func (l NetworkList) HasNetwork(network Network) bool {
	for _, value := range l.Network {
		if string(value) == string(network) {
			return true
		}
	}
	return false
}

func (l NetworkList) Get(idx int) Network {
	return l.Network[idx]
}

// Size returns the number of networks in this network list.
func (l NetworkList) Size() int {
	return len(l.Network)
}
