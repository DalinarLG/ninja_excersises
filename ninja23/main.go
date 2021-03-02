package main

import "fmt"

func main(){
	x := make([]string, 32, 32)
	x = []string{
		"Amazonas",
		"Antioquia", 
		"Arauca", 
		"Atlántico",
		"Bolívar",
		"Boyacá",
		"Caldas",
		"Caquetá",
		"Casanare",
		"Cauca",
		"Cesar",
		"Chocó",
		"Córdoba",
		"Cundinamarca",
		"Guainía",
		"Guaviare",
		"Huila",
		"La Guajira",
		"Magdalena",
		"Meta",
		"Nariño",
		"Norte de Santander",
		"Putumayo",
		"Quindío",
		"Risaralda",
		"San Andrés y Providencia",
		"Santander",
		"Sucre",
		"Tolima",
		"Valle del Cauca",
		"Vaupés",
		"Vichada",
	}

	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i:=0; i<=len(x)-1; i++{
		fmt.Printf("%v: %s\n",i,x[i])
	}
}