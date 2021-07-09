package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

type request struct {
	method  string
	path    string
	headers []header
}

type header struct {
	key   string
	value string
}

func main() {
	fmt.Println("Start stupid webserver....")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go handle(conn)
	}
}

func handle(conn io.ReadWriteCloser) {
	defer conn.Close()

	r := getHttpRequest(conn)

	body, err := retrieveBody(r.path)
	if err != nil {
		fmt.Fprintf(conn, "HTTP/1.1 404 NOT FOUND\r\n")
		fmt.Fprintf(conn, "Content-Type: text/html; charset=utf-8\r\n")
		fmt.Fprintf(conn, "\r\n")
	} else {
		fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprintf(conn, "Content-Type: text/html; charset=utf-8\r\n")
		fmt.Fprintf(conn, "\r\n")
		fmt.Fprintf(conn, body)
	}
}

func retrieveBody(path string) (string, error) {
	if path == "/" {
		path = "/index.html"
	}

	staticFolder := os.Getenv("HTTP_LOCATION")

	dat, err := ioutil.ReadFile(staticFolder + path)
	if err != nil {
		return "", err
	}
	return string(dat), nil
}

func getHttpRequest(conn io.ReadWriter) request {
	result := request{}

	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		h := scanner.Text()
		if h == "" {
			// headers are complete
			break
		}
		if i == 0 {
			requestElements := strings.Fields(h)
			result.method = requestElements[0]
			result.path = requestElements[1]
		} else {
			e := strings.SplitN(h, " ", 2)
			if len(e) == 2 {
				key := e[0][:len(e[0])-1]
				result.headers = append(result.headers, header{key: key, value: e[1]})
			}
		}
		i++
	}

	return result
}
