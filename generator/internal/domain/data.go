package domain

type Data struct {
	DeviceId    int `json:"device_id"`
	Temperature int `json:"temp"`
	Pressure    int `json:"press"`
	BatLvl      int `json:"bat_lvl"`
}
