package domain

import (
	"math/rand/v2"
)

func GenerateData() []Data {
	return []Data{
		{DeviceId: 0, Temperature: rand.IntN(60), Pressure: rand.IntN(1000), BatLvl: rand.IntN(100)},
		{DeviceId: 1, Temperature: rand.IntN(60), Pressure: rand.IntN(1000), BatLvl: rand.IntN(100)},
		{DeviceId: 2, Temperature: rand.IntN(60), Pressure: rand.IntN(1000), BatLvl: rand.IntN(100)},
	}
}
