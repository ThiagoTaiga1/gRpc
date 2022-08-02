package main

import (
	"fmt"
	"main/common"
	"net/rpc"
)

func main() {
	// // obtém o cliente RPC discando no terminal `rpc.DefaultRPCPath`
	client, _ := rpc.DialHTTP("tcp", "localhost:9000")
	// // cria a variável john do tipo `common.Student`
	var thiago common.Student
	// pega um student id 1
	if err := client.Call("College.Get", 1, &thiago); err != nil {
		fmt.Println("Error:1College.Get()", err)
	} else {
		fmt.Printf("Sucess:1 Student '%s'found with id=1 \n", thiago.FullName())
	}
	// adiciona student com ids `1`

	if err := client.Call("College.Add", common.Student{
		ID:        1,
		FirstName: "thiago",
		LastName:  "silva",
	}, &thiago); err != nil {
		fmt.Println("Error:2 College.Add()", err)
	} else {
		fmt.Printf("Sucess:2 Student '%s' created with id=1 \n", thiago.FullName())
	}

}
