package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"swarm_server/node"
	"swarm_server/space"
	"sync"
)

func main() {
	fmt.Println("Hello and Welcome to the Drone Swarm Simulation software!")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the IP Adress to use, and press ENTER: ")
	ip, _ := reader.ReadString('\n') // Read until newline

	fmt.Println("Enter the Port to use, and press ENTER: ")
	port, _ := reader.ReadString('\n')

	listener := estNetLisener(reader, ip, port)

	fmt.Println("\nSwarm server running...")

	fmt.Println("\nPlease note that all units used are SI-based.")

	fmt.Println(`
	You'll now be required to enter the space coordinate of the destination on the X,Y,Z plane.
	The longest axis will be regarded as the main path. The others will be used for side movements to avoind collisions.
	`)

	x, y, z := dest(reader)

	fmt.Printf("\nWaiting for node clients to connect...\n\n")

	var (
		sp            = space.Plane{End: space.Point{X: x, Y: y, Z: z}}
		occupiedSpace []space.Point
	)

	var (
		wg sync.WaitGroup

		mu    sync.Mutex
		count int
	)

	for {
		conn, err := listener.Accept()
		// conn closed in MoveTo method
		if err != nil {
			log.Print(err)
			break
		}

		wg.Add(1)

		go func() {
			mu.Lock()
			count++
			num := count
			mu.Unlock()
			defer wg.Done()
			var n node.Node
			point, done := n.MoveTo(conn, occupiedSpace, sp.End)
			fmt.Printf("Node # %d: done: %t, at: (%d, %d, %d)\n", num, done, point.X, point.Y, point.Z)
			count--
		}()
	}
	wg.Wait()
	fmt.Println("\nShutdown.")
}

func estNetLisener(reader *bufio.Reader, ip string, port string) net.Listener {

	var count int8

	if count > 0 {

		fmt.Println("Enter 1 to change the IP Adress or 2 to change the Port, and press ENTER: ")

		input, _ := reader.ReadString('\n') // Read until newline
		if input != "" {
			// num64, err := strconv.ParseInt(str, 10, 64) // base 10, 64-bit size
			num8, err := strconv.ParseInt(input, 10, 8)

			if err != nil {
				fmt.Println("Please enter a valid number.")
				estNetLisener(reader, ip, port)
			}

			if num8 == 1 {
				// Change IP
				fmt.Println("Please enter the IP Address to use and press ENTER: ")
				ip, _ = reader.ReadString('\n')
			} else {
				// Change Port
				fmt.Println("Please enter the Port number to use and press ENTER: ")
				port, _ = reader.ReadString('\n')
			}

		} else {
			estNetLisener(reader, ip, port)
		}
	}

	var listener net.Listener
	var builder strings.Builder

	builder.WriteString(strings.TrimSpace(ip))
	builder.WriteString(":")
	builder.WriteString(strings.TrimSpace(port))

	netAddr := builder.String()

	fmt.Printf("Listening at: %s\n", netAddr)

	listener, err := net.Listen("tcp", netAddr)
	if err != nil {
		log.Fatal(err)
		count++
		estNetLisener(reader, ip, port)
	}
	return listener
}

func dest(reader *bufio.Reader) (int, int, int) {
	fmt.Println("Enter the value of X on the destination point, and press ENTER: ")
	xString, _ := reader.ReadString('\n')
	x, err := strconv.ParseInt(strings.TrimSpace(xString), 10, 32)
	if err != nil {
		fmt.Println("Please enter a valid integer.")
		dest(reader)
	}
	

	fmt.Println("Enter the value of Y on the destination point, and press ENTER: ")
	yString, _ := reader.ReadString('\n')
	y, err := strconv.ParseInt(strings.TrimSpace(yString), 10, 32)
	if err != nil {
		fmt.Println("Please enter a valid integer.")
		dest(reader)
	}

	fmt.Println("Enter the value of Z on the destination point, and press ENTER: ")
	zString, _ := reader.ReadString('\n')
	z, err := strconv.ParseInt(strings.TrimSpace(zString), 10, 32)
	if err != nil {
		fmt.Println("Please enter a valid integer.")
		dest(reader)
	}

	return int(x), int(y), int(z)
}
