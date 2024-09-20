package patterns

import (
	"fmt"
	"net"
	"os"
	"regexp"
)

var r, _ = regexp.Compile("GET (.+) HTTP/1.1\r\n")

func handleHttpRequest(conn net.Conn) {
	buff := make([]byte, 1024)
	size, _ := conn.Read(buff)
	if r.Match(buff[:size]) {
		file, err := os.ReadFile(fmt.Sprintf("../resources/%s", r.FindSubmatch(buff[:size])[1]))
		if err == nil {
			conn.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Length: %d\r\n\r\n", len(file))))
			conn.Write(file)
		} else {
			conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n<html>Not Found</html>"))
		}
	} else {
		conn.Write([]byte("HTTP/1.1 500 Internal Server Error\r\n\r\n<html>Internal Server Error</html>"))
	}
}

func StartHttpWorkers(n int, incomingConnections <-chan net.Conn) {
	for i := 0; i < n; i++ {
		go func() {
			for c := range incomingConnections {
				handleHttpRequest(c)
			}
		}()
	}
}

func Run3() {
	incomingConnections := make(chan net.Conn)
	StartHttpWorkers(3, incomingConnections)

	server, _ := net.Listen("tcp", ":8080")
	defer server.Close()
	for {
		conn, _ := server.Accept()
		incomingConnections <- conn
	}
}

func Run4() {
	incomingConnections := make(chan net.Conn)
	server, _ := net.Listen("tcp", ":8080")
	defer server.Close()
	for {
		conn, _ := server.Accept()
		select {
		case incomingConnections <- conn:
		default:
			fmt.Println("Server is busy")
			conn.Write([]byte("HTTP/1.1 429 Too Many Requests\r\n\r\n<html>Too Many Requests</html>"))
			conn.Close()
		}
	}
}
