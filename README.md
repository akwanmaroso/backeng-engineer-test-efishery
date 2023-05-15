# Backend Engineer Test by EFishery

## Prerequisites

- Node.js (v14 or higher) and npm (Node Package Manager) installed on your machine. You can download Node.js from the official website: [https://nodejs.org](https://nodejs.org)
- Go (v1.16 or higher) installed on your machine. You can download Go from the official website: [https://golang.org](https://golang.org)
- Docker and Docker Compose installed on your machine. You can download Docker from the official website: [https://www.docker.com](https://www.docker.com)

## Getting Started

1. Clone the repository:

   ```shell
   git clone https://github.com/your-username/my-program.git
   cd my-program
   ```

2. Set up the environment variables:

   - Rename the .env.example file to .env.
   - Open the .env file and provide the necessary values for the variables.

3. Build and run the application:

   ```shell
   docker-compose up --build -d
   ```

   This command will build the Docker images and start the containers for the application.

4. Access the application:
   - The Auth service will be available at http://localhost:4999.
   - The Core service will be available at http://localhost:5001.

## Stopping the Application

To stop the application and remove the Docker containers, use the following command:

```shell
   docker-compose down
```

This will stop and remove the containers, but your data will be preserved.

## Contributing

Contributions are welcome! If you encounter any issues
