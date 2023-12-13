package main

func SomaInteiros(m map[string]int64) int64{
	var soma int64
	for _, v := range m {
		soma += v
	}
	return soma
}

// Sem generics você tem muita repetiçãode código. pode até ser que vc não veja essa repetição em um mesmo modulo, mas ocorria com certa frequencia em versões antecessoras a 1.18
func SomaFlutuantes(m map[string]float64) float64{
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

func main(){
	println(SomaInteiros(map[string]int64{"a": 1,"b": 2, "c": 3}))
	println(SomaFlutuantes(map[string]float64{"a": 1.1,"b": 2.2, "c": 3.3}))
}