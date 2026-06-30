package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"net"
)

type isPrime struct {
	Method string      `json:"method"`
	Number interface{} `json:"number"`
}

func main() {
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			if err == io.EOF {
				break
			}
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}

		var req isPrime
		if err := json.Unmarshal(buf[:n], &req); err != nil {
			continue
		}

		v, ok := req.Number.(float64)
		if !ok {
			continue
		}

		if v != math.Trunc(v) {
			continue
		}

		isPrime := checkPrime(int64(v))
		res := fmt.Sprintf(`{"method":"isPrime","prime": %s}`, isPrime)

		_, err = conn.Write([]byte(res))
		if err != nil {
			return
		}
	}
}

func checkPrime(n int64) string {
	if big.NewInt(n).ProbablyPrime(0) {
		return "true"
	}
	return "false"
}
