package tue2

import (
	"fmt"
)

func Run() {
	qList := make([]float64, 0)
	etaList := make([]float64, 0)
	etaExList := make([]float64, 0)
	for i := 0.75; i <= 0.9501; i += 0.001 {
		q := i
		ws2 := -100000.0
		m := -ws2 / (3375.13 - (2357.65*q + 251.18))
		F_CH4 := m * (3375.13 - 267.88) / (842.626 - 74.520)
		ws1 := m * (267.88 - 251.18)
		//QL := m * (2357.65*q + 251.18 - 251.18)
		Ha := -74.520 * F_CH4
		Hb := -842.626 * F_CH4
		eta := -(ws1 + ws2) / (Ha - Hb)
		Exa := 824.60 * F_CH4
		Exb := 19.58 * F_CH4
		etaEx := -(ws1 + ws2) / (Exa - Exb)
		qList = append(qList, q)
		etaList = append(etaList, eta)
		etaExList = append(etaExList, etaEx)
	}

	q := 0.85
	ws2 := -100000.0
	m := -ws2 / (3375.13 - (2357.65*q + 251.18))
	F_CH4 := m * (3375.13 - 267.88) / (842.626 - 74.520)
	ws1 := m * (267.88 - 251.18)
	//QL := m * (2357.65*q + 251.18 - 251.18)
	Ha := -74.520 * F_CH4
	Hb := -842.626 * F_CH4
	eta := -(ws1 + ws2) / (Ha - Hb)
	Exa := 824.60 * F_CH4
	Exb := 19.58 * F_CH4
	etaEx := -(ws1 + ws2) / (Exa - Exb)

	fmt.Println("ws2:", ws2)
	fmt.Println("m:", m)
	fmt.Println("F_CH4:", F_CH4)
	fmt.Println("ws1:", ws1)
	fmt.Println("eta:", eta)
	fmt.Println("eta_Ex:", etaEx)
}
