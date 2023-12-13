package main

// Usamos uma interface para armazenar os tipos que iremos utilizar
// usamos ~ para fazer aproximação de tipos
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// criamos aqui um tipo MyNumber que armazena inteiros
// Teóricamente o soma Generica deveria ler isto, mas não vai conseguir
type MyNumber int

// O que Number esta fazendo é o mesmo que as constraints ou seja elas são regras que definem quem pode passar por essa generic
func SomaGenerica[T Number] (m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// Outra maneira de uso para generics
// Aqui não sabemos o tipo da variavel que sera passada, mas usamos uma interface do tipo Number para garantir tipos que serão permitidos.
func Soma[T Number] (number1 T, number2 T) T{
	return number1 + number2
}

func main() {
	var x, y, z MyNumber
	x = 1
	y = 2
	z = 3
	somafuncgenerica := Soma(1, 4)


	println(SomaGenerica(map[string]int64{"a": 1, "b": 2, "c": 3}))
	println(SomaGenerica(map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}))
	println(SomaGenerica(map[string]MyNumber{"a": x, "b": y, "c": z}))
	println(somafuncgenerica)
}