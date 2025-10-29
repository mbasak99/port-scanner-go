# Background
I've wanted to properly learn about networking and how things work when you make requests and call endpoints and stuff like that. So this is my excuse to learn two things I've been wanting to understand, Go and Networking, and combine them into one. This project aims to expand my understanding of what ports are, how a device (my PC in this instance) uses its ports to communicate with the external world, and the security implications of ports.

# Current Level
I know TCP is used for network calls in REST APIs. HTTP is 80 and HTTPS is 443. No real knowledge or understanding beyond that as of now.

# Reflection Post Project
- `ScanPort` function essentially checks what port is actively being used at the moment. When I first implemented this I didn't realize it'll only return true if there's an active process on a given port. Which is why the call to google.com worked when all my other ports failed. After starting a basic http server on port 3001 did I get a true back.
- `InitialScan` function leverages the previously created function and reuses the logic to check ports 1-1024 and return whether the given port is open or closed along with the port number all wrapped in a nice struct. Pretty straightforward.
- After that I basically added checks for both tcp and udp and extended my function to handle more than ports 1-1024 -- up to 65535, checking all 131,070 ports.

Checking all the ports definitely makes it take some time, even on my Ryzen 9 9900X :skull:. I want to see if using Go threads would speed this process up. Gonna time both.

# Steps to run this project
`go run main.go`
