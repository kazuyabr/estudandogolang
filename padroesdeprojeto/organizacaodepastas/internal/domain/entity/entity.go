package entity

//Exemplo de metodo que temos acesso só no mesmo package
type productInterno struct{
	Nome string
}

//Exemplo de metodo que temos acesso ao metodo exportado
type ProductExterno struct{
	Nome string
}