package main

import (
	"constraints"
)

// Usamos uma interface para armazenar os tipos que iremos utilizar
// usamos ~ para fazer aproximação de tipos
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// Quando comparamos 2 valores que serão de mesmo tipo é possivel, mas se tentar usar essa função para comparar um int com um float por exemplo vai falhar pois os tipos são diferentes
// Uma vez que cada T refere-se ao sua própria variavel podendo todos terem tipos distintos.
func ComparacaoDeMesmoTipo[T Number] (number1 T, number2 T) T{
	if number1 == number2 {
		return number1
	}
	return number2
}

// Agora vamos a uma comparação mais generica
// Any é um tipo generico que pode ser utilizado para receber qualquer valor. entretanto o mesmo não é comparavel ja que golang é fortemente tipado
// A linha de código ficara comentada para não dar problema ao código, mas recomendo que descomente o código para visualizar o erro
// func ComparacaoGenerica[T any] (number1 T, number2 T) T{
// 	if number1 == number2 {
// 		return number1
// 	}
// 	return number2
// }

// Agora usamos uma tipagem generica que permite comparação
// Comparable é uma interface que compara mesmos tipos e ela resolve esse problema de comparar tipos diferentes garantido que eles se tornem o mesmo tipo
func ComparacaoGenerica[T comparable] (number1 T, number2 T) T{
	if number1 == number2 {
		return number1
	}
	return number2
}

// Como nem tudo são flores comparable permite comparações de igualdade, mas quando se trata de grandezas ele não consegue
// Vamos dizer q estamos buscando o valor maior entre os elementos abaixo (código ficara comentado para evitar problemas de compilação, mas recomendo descomentar para visualizar o problema)
// func Max[T comparable] (number1 T, number2 T) T{
// 	if number1 > number2 {
// 		return number1
// 	}
// 	return number2
// }

// Para resolver esses problemas precisamos de fato utilizar um pacote de constraints que tratam estes problemas
// existe um pacote oficial do go chamado constraints que você pode instalar
// use get para instalar -> go get golang.org/x/exp/constraints
func Max[T constraints.Ordered] (number1 T, number2 T) T{
	if number1 > number2 {
		return number1
	}
	return number2
}

func main() {
	var x, y, z MyNumber
	x = 1
	y = 2
	z = 3
	somafuncgenerica := Soma(1, 4)
	comparacaoMesmoTipo := ComparacaoDeMesmoTipo(2, 2)

	println(SomaGenerica(map[string]int64{"a": 1, "b": 2, "c": 3}))
	println(SomaGenerica(map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}))
	println(SomaGenerica(map[string]MyNumber{"a": x, "b": y, "c": z}))
	println(somafuncgenerica)
	println(comparacaoMesmoTipo)
}