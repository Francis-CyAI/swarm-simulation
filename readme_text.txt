This project was conceived to simulates the behavior of a swarm of drones navigating from an origin to a specified destination in 3D space. It consists of two main programs: a server and a client, which communicate in real-time over TCP connections. The server handles the core logic for drone movement, while the client receives real-time updates and logs the drone’s position to the standard output.

Purpose 
The primary purpose of this project is to model and simulate the coordinated movement of a swarm of drones, demonstrating various aspects of Go development, including concurrency management and real-time server-client communication in a high-performance computing environment. This simulation can serve as a foundational framework for further development in drone swarm navigation and robotics.

Features
•	Concurrent Drone Simulation: Each drone is represented by a goroutine on the server, simulating its movement and communicating with the client.
•	Real-Time Communication: The server and client communicate over TCP, with the server sending real-time position updates to the clients.
•	3D Movement Logic: The server calculates the next position for each drone in 3D space based on its current location and destination.
•	Main Path Axis Prioritization: Each drone determines the main path axis and prioritizes movement along this axis, using other axes primarily for collision avoidance.
•	Mutual Exclusion: Critical section, such as space point access and unique ID generation for drones are protected using mutex locks to ensure thread-safe operations.
•	Synchronization: A `sync.WaitGroup` is used to synchronize the completion of all drone movement simulations, ensuring the main goroutine wait for all node goroutines to finish.
•	Robust Error Handling: Implemented error handling for TCP connections, ensuring robust and reliable communication between the server and clients.
•	Performance Optimization: Utilized Go’s lightweight goroutines for efficient concurrent execution, allowing the simulation of a large number of drones.

