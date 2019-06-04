package main
import (
	"fmt"
	"net"
	"bufio"
)

func startServer() {
	fmt.Println("Iniciando Servidor...")
	ln, _:= net.Listen("tcp", "10.11.97.218:9000")
	defer ln.Close()
	fmt.Println("Escuchando por puerto 8000 nodo cliente")
	for {
		con, _ := ln.Accept()
		go handle(con)
	}
}

func handle(con net.Conn) {
	defer con.Close()
	r := bufio.NewReader(con)
	msg, _ := r.ReadString('\n')
	fmt.Println("Recibido msg nodo cliente: ", msg)
	startClient(msg)
}

func startClient(msg string) {
	fmt.Println("Iniciando Cliente...")
	con,_ := net.Dial("tcp","10.11.97.217:9000")
	defer con.Close()
	fmt.Fprint(con,msg)
}

func main(){
	startServer()
	//startClient("se oye ?")
}