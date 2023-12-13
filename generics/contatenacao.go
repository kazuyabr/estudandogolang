package main

type stringer interface {
	String() string
}


//Para concatenação não podemos usar any, pois não é possivel garantir que o tipo sera concatenavel
// código comentado para evitar erro de compilação, mas descomente para ver o problema
// func concat[T any](vals []T) string{
// 	result := ""
// 	for _, val := range vals{
// 		result += val.String()
// 	}
// 	return result
// }

// Agora o T sabe que ele é um string, então garantimos que ele é concatenavel
func concat[T stringer](vals []T) string{
	result := ""
	for _, val := range vals{
		result += val.String() //Metodo que vai dar muito erro se não for tratado
	}
	return result
}

// Criamos um tipo para manipular strings
type MyString string

// funcao responsavel por criar um metodo String() para o tipo MyString evitando assim erros para concatenação
// essa funcao faz com que a nossa string ganhe um metodo que respeita a interface stringer e a partir dai nossa concatenacao para de dar erro
func(m MyString) String() string{
	return string(m)
}

func main(){
	// Aqui seria uma tentativa de uso da concatenação, mas é erronea pois o tipo string não possui um metodo String()
	// E string tambem não possui metodo de conversão para string ja q ele e uma string ou seja não seria possivel fazer ele se tornar uma string que contenha o metodo String()
	// código comentado para evitar erros, mas descomente para visualizar o problema
	// println(concat([]string{"a", "b", "c"}))

	// Esse aqui tambem teria problema ja q ele não possui o metodo String()
	// Entretanto permite que usemos uma função para resolver isso	
	println(concat([]MyString{"a", "b", "c"}))


}