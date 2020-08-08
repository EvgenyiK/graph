package midleware

import (
    "github.com/cheekybits/genny/generic"
)

//выбираем тип двоичного дерева поиска
type Item generic.Type

//единственный узел состовляющий дерево
type Node struct {
    value Item
}

/*func GraphsRoute() {
	graphs,err:= getAllGraph()
	if err != nil {
		log.Fatalf("error")
	}

	fmt.Println(graphs)

	

	
    
}*/