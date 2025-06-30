package models

type Data struct {
	DeviceId    int `json:"device_id" gorm:"not null"`
	Temperature int `json:"temp" gorm:"not null"`
	Pressure    int `json:"press" gorm:"not null"`
	BatLvl      int `json:"bat_lvl" gorm:"not null"`
}
