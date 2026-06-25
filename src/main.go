package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"
	"github.com/joho/godotenv"
)

func doClientConnection(connection net.Conn, listener net.Listener) {
    defer connection.Close()

    buffer := make([]byte, 1024)

    mLen, err := connection.Read(buffer)
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }
	clientIP, _, err := net.SplitHostPort(connection.RemoteAddr().String())
	if err != nil {
		clientIP = connection.RemoteAddr().String()
	}

    fmt.Println("[" + time.Now().Format(time.RFC822) + "] "+clientIP+": " + string(buffer[:mLen]))
	haveInformation(clientIP)

    listener.Close()
}

func haveInformation(ip string) {
    token := os.Getenv("TOKEN_ABUSEIPDP")

    out, err := exec.Command(
        "curl",
        "-G",
        "https://api.abuseipdb.com/api/v2/check",
        "--data-urlencode", "ipAddress="+ip,
        "-d", "maxAgeInDays=90",
        "-d", "verbose",
        "-H", "Key: "+token,
        "-H", "Accept: application/json",
    ).Output()

    if err != nil {
        fmt.Println("Erreur :", err)
        return
    }

    fmt.Println(string(out))
}

func loadEnv() (string, string, string){
	err := godotenv.Load("../.env")
    if err != nil {
        fmt.Println("Erreur lors du chargement du .env")
		os.Exit(1)
    }

    serverHost := os.Getenv("SERVER_HOST")
    serverPort := os.Getenv("SERVER_PORT")
    serverType := os.Getenv("SERVER_TYPE")

	return serverHost, serverPort, serverType
}

func main(){
	
	server_host, server_port, server_type := loadEnv()

	socket_listen, errSocketListen := net.Listen(server_type, server_host+":"+server_port)
	if errSocketListen != nil{
        fmt.Println("Error listening:", errSocketListen.Error())
		os.Exit(1)
	}


	fmt.Println("Listening on "+ server_host+":"+server_port+".")
	fmt.Println("Waiting for clients...")

	for {
		acceptSocket, err := socket_listen.Accept()
		if err != nil {
			fmt.Println("Server stopped")
			break
		}
		fmt.Println("Client connected")
		go doClientConnection(acceptSocket, socket_listen)
	}

}