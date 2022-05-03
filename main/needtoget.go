package main

type Need struct {
	Needmap map[string]float64
}
type Standard struct {
	Stand map[string]float64
}

//GetNeed() map[string]float64{}
func GetNeed(TEst *Already) *Need {
	Stand_ := new(Standard)
	Need := new(Need)
	Stand_.GetStandard()
	Need.Needmap = make(map[string]float64)
	Need.Needmap["Protein"] = Stand_.Stand["Protein"] - TEst.Get["Protein"]

	Need.Needmap["Va"] = Stand_.Stand["Va"] - TEst.Get["Va"]

	Need.Needmap["Vb"] = Stand_.Stand["Vb"] - TEst.Get["Vb"]

	Need.Needmap["Vc"] = Stand_.Stand["Vc"] - TEst.Get["Vc"]

	Need.Needmap["Fat"] = Stand_.Stand["Fat"] - TEst.Get["Fat"]
	Need.Needmap["Carbon"] = Stand_.Stand["Carbon"] - TEst.Get["Carbon"]
	return Need
}

//
func (this *Standard) GetStandard() {
	this.Stand = make(map[string]float64)
	this.Stand["Protein"] = 60
	this.Stand["Va"] = 0.8
	this.Stand["Vb"] = 6
	this.Stand["Vc"] = 100
	this.Stand["Fat"] = 31.5
	this.Stand["Carbon"] = 75
}
