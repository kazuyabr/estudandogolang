package main

// Usamos uma interface para armazenar os tipos que iremos utilizar
type Number inferface {
	int | int32 | int64 | float32 | float64
}

// criamos aqui um tipo MyNumber que armazena inteiros
// Teóricamente o soma Generica deveria ler isto, mas não vai conseguir
type MyNumber int

// O que Number esta fazendo é o mesmo que as constraints ou seja elas são regras que definem quem pode passar por essa generic
func SomaGenerica[T Number ] (m map[string]T) T{
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func main(){
	var x, y, z MyNumber
	x = 1
	y = 2
	z = 3

	println(SomaGenerica(map[string]int64{"a": 1,"b": 2, "c": 3}))
	println(SomaGenerica(map[string]float64{"a": 1.1,"b": 2.2, "c": 3.3}))
	println(SomaGenerica(map[string]MyNumber{"a": x,"b": y, "c": z}))
}