package models

type Servers struct {
	Servers []Server `json:"servers"`
}

type Server struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	ProductStatus string `json:"product_status"`
	Datacenter    string `json:"datacenter"`
	BillingCycle  string `json:"billingcycle"`
	Price         string `json:"price"`
	Plan          string `json:"plan"`
	Ip            string `json:"ip"`
	Disk          int    `json:"disk"`
}

type SingleServer struct {
	Server struct {
		Id            int    `json:"id"`
		Name          string `json:"name"`
		ProductStatus string `json:"product_status"`
		Datacenter    string `json:"datacenter"`
		BillingCycle  string `json:"billingcycle"`
		Price         string `json:"price"`
		Plan          string `json:"plan"`
		Ip            string `json:"ip"`
		Disk          int    `json:"disk"`
	} `json:"server"`
}

type Status struct {
	Server struct {
		Status    string `json:"status"`
		Consuming struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
			Disk   string `json:"disk"`
		} `json:"consuming"`
	} `json:"server"`
}

type Result struct {
	Result string `json:"result"`
	Status string `json:"status"`
}
