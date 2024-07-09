# Drone Swarm Simulation

## Project Overview

This project simulates the behavior of a swarm of drones navigating from an origin to a specified destination in 3D space. It consists of two main programs: a server and a client, which communicate in real-time over TCP connections. The server handles the core logic for drone movement, while the client receives real-time updates and logs the drones' positions to the standard output.

## Purpose

The primary purpose of this project is to model and simulate the coordinated movement of a swarm of drones, demonstrating various aspects of Go development, including concurrency management and real-time client-server communication. This simulation can serve as a foundational framework for further development in drone swarm behavior, robotics, and distributed systems.

## Features

- **Concurrent Drone Simulation**: Each drone is represented by a goroutine on the server, simulating its movement and communicating with the client.
- **Real-Time Communication**: The server and clients communicate over TCP, with the server sending real-time position updates to the clients.
- **3D Movement Logic**: The server calculates the next position for each drone in 3D space based on its current location and the destination.
- **Main Path Axis Prioritization**: Each drone determines the main path axis and prioritizes movement along this axis, using other axes primarily for collision avoidance.
- **Mutual Exclusion**: Critical sections, such as space point access and unique ID generation for drones, are protected using mutex locks to ensure thread-safe operations.
- **Synchronization**: A `sync.WaitGroup` is used to synchronize the completion of all drone simulations, ensuring the main goroutine waits for all node goroutines to finish.
- **Robust Error Handling**: Implemented error handling for TCP connections, ensuring robust and reliable communication between the server and clients.
- **Performance Optimization**: Utilized Go’s lightweight goroutines for efficient concurrent execution, allowing the simulation of a large number of drones.

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- A working TCP/IP network setup for server-client communication

### Installation

1. **Clone the Repository**:
    ```sh
    git clone https://github.com/Francis-CyAI/swarm-simulation.git
    cd swarm-simulation
    ```

2. **Build the Server**:
    ```sh
    go build -o server ./swarm_server
    ```

3. **Build the Client**:
    ```sh
    go build -o client ./swarm_client
    ```

### Usage

1. **Start the Server**:
    ```sh
    ./swarm_server
    ```

2. **Start the Client(s)**: (do this for as many times as the number of clients needed)
    ```sh
    ./swarm_client
    ```

### Usage on Windows

1. **Start the Server**:
    ```sh
    ./start_server.bat
    ```

2. **Start the Client(s)**: (do this for as many times as the number of clients needed)
    ```sh
    ./start_client.bat


### Configuration

The server and clients can be configured via command-line arguments when prompted.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request for any enhancements or bug fixes.

## License

This project is licensed under the GPLv3 License. See the [License](LICENSE) for details.

## Contact

For questions or inquiries, please contact [kalungafrancis28@gmail.com](mailto:kalungafrancis28@gmail.com).

---

Thank you for checking out the Drone Swarm Simulation project! Your feedback and contributions are greatly appreciated.
