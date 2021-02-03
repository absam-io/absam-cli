package models

type FirewallStatus struct {
	Enabled bool `json:"enabled"`
}

type Rule struct {
	Pos     int    `json:"pos"`
	Port    string `json:"port"`
	Proto   string `json:"proto"`
	IP      string `json:"ip"`
	Expires string `json:"expires"`
	Comment string `json:"comment"`
}

type FirewallRules struct {
	Rules struct {
		In  []Rule `json:"in"`
		Out []Rule `json:"out"`
	} `json:"rules"`
}

type ChangeFirewallSuccess struct {
	Success string `json:"success"`
}
