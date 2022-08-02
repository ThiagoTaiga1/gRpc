package main

import (
	"io"
	"main/common"
	"net/http"
	"net/rpc"
)

func main() {
	//creando o objeto `*college`
	mit := common.NewCollege()
	//registra `mit` objeto em `rpc.DefaultServer`
	rpc.Register(mit)
	// / registra um manipulador HTTP para comunicação RPC em `http.DefaultServeMux` (padrão)
	// registra um manipulador no endpoint `rpc.DefaultRPCPath` para responder a mensagens RPC
	// registra um manipulador no endpoint `rpc.DefaultDebugPath` para depuração
	rpc.HandleHTTP()
	//teste de endpoint
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "RPC SERVER LIVE!")

	})
	http.ListenAndServe(":9000", nil)
}
