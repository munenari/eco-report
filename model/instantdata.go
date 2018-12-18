package model

// InstantData struct
type InstantData struct {
	Circuits []Circuit `json:"circuits"`
	PV       PV        `json:"pv"` // 太陽光発電
	SB       int       `json:"sb"` // 蓄電池
	DB       int       `json:"db"` // 買電量
	FC       FC        `json:"fc"` // エネファーム
}

// Circuit struct (power consumption)
type Circuit struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

// PV struct (solar power)
type PV struct {
	Generate int      `json:"generate"`
	Sell     int      `json:"sell"`
	List     []Device `json:"list"`
}

// FC struct (enefarm)
type FC struct {
	Consumption int `json:"comsumption"`
	Generate    int `json:"generate"`
}

// Device struct
type Device struct {
	MakerCode string `json:"makerCode"`
	Generate  int    `json:"generate"`
	Sell      int    `json:"sell"`
	GUID      string `json:"guid"`
}

// GetCircuitsMap data for easy visualization in kibana
func (d *InstantData) GetCircuitsMap() *map[string]int {
	res := make(map[string]int)
	for _, v := range d.Circuits {
		res[v.Name] = v.Value
	}
	return &res
}
