package main

// Podemos usar uma variavel qualquer para definir que ela seja a tipagem abordada
// para isso na frente da função acrescentamos [Variavel tipos...]  como demonstrado na função SomaGenerica
// Problema disto é que ela pode ficar fortemente verbosa.
func SomaGenerica[T int32 | int64 | float32 |float64 ] (m map[string]T) T{
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