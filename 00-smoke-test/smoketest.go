package smoketest

import (
	"io"
	"log"
	"net"
)

func Run(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
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

		_, err = conn.Write(buf[:n])
		if err != nil {
			return
		}
	}
}
