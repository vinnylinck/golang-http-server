package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const fallbackport string = "8080"

// main function - entry point
func main() {
	// retrieving env vars
	port := getPort()

	// creating / starting tcp listener
	tcpli, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer tcpli.Close()

	// reading buffer from connections
	for {
		conn, err := tcpli.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

// Returns TCP port
func getPort() string {
	value := getenv("PORT", fallbackport)
	return fmt.Sprintf(":%s", value)
}

// Reads env var and returns fallback value if it is not set
func getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// connection handler
func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

// request handler
func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

// multiplexer - router
func mux(conn net.Conn, ln string) {
	// request line
	m := strings.Fields(ln)[0] // method
	u := strings.Fields(ln)[1] // uri
	fmt.Println("***METHOD", m)
	fmt.Println("***URI", u)

	// multiplexer
	if m == "GET" && u == "/" {
		index(conn)
	}

	/*
		if m == "GET" && u == "/about" {
			about(conn)
		}
		if m == "GET" && u == "/contact" {
			contact(conn)
		}
		if m == "GET" && u == "/apply" {
			apply(conn)
		}
		if m == "POST" && u == "/apply" {
			applyProcess(conn)
		}
	*/
}

// HANDLER - GET /
func index(conn net.Conn) {
	body := `{ "ok": true }`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: application/json\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
