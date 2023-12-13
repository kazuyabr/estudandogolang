package main

// Usamos uma interface para armazenar os tipos que iremos utilizar
type Number inferface {
	int | int32 | int64 | float32 | float64
}


// O que Number esta fazendo é o mesmo que as constraints ou seja elas são regras que definem quem pode passar por essa generic
func SomaGenerica[T Number ] (m map[string]T) T{
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main(){
	println(SomaGenerica(map[string]int64{"a": 1,"b": 2, "c": 3}))
	println(SomaGenerica(map[string]float64{"a": 1.1,"b": 2.2, "c": 3.3}))
}