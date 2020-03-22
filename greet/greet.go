package greet
//El paquete debe ser especifico e indicar cual es su funcionalidad
//para la facilidad de lectura, ejemplo aqui sabemos que el paquete greet se encargara 
//de los saludos, entonces se llama greet y sus funciones indican en que idioma saludara, 

//Su la primera letra es en mayuscula Go la exporta, caso contrario No, exportada indica que pueden ser visibles desde otros paquetes. 
var emoji = " by greet"  //las variables declaradas aqui deben ir con var 

//GreetEnglish retorna saludo en ingles 
func English() string {
	return "Hi" + emoji
}

func Italian() string {
	return "Ciao" + emoji
}