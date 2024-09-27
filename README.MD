# Go Microservices

## Project Introduction
This project aims to illustrate the advantages of microservices architecture by developing a collection of small, self-contained services that are loosely coupled and can be independently developed, deployed, and scaled. By leveraging Go, I created a dynamic architecture that facilitates rapid and reliable delivery of complex applications. The project encompasses several essential microservices that enhance functionality and user experience.

## Features
**Authentication Service:** A dedicated service for user authentication, utilizing a Postgres database for secure user data management.
**Logging Service**: Collects and stores logs using MongoDB, enabling efficient data retrieval and analysis.
**Message Listener:** Receives messages from RabbitMQ, facilitating inter-service communication and event handling.
**Broker Service:** Serves as an optional single entry point into the microservice cluster, streamlining requests to various services.
**Mail Service:** Processes JSON payloads to generate and send formatted emails, enhancing user interaction.
**Dynamic Communication:** Integrates various communication methods, including REST API, RPC, gRPC, and AMQP, to enable seamless service interactions.

## Technologies Used
**Programming Language:** Go (Golang)
**Databases:** Postgres for authentication and MongoDB for logging
**Message Broker:** RabbitMQ for message queuing
**Containerization:** Docker for packaging services
**Orchestration:** Docker Swarm for deploying and managing services

## Running the Project
### Prerequisites
Before running the project, make sure you have the following installed on your machine:
* Docker

### Method 1: Using Docker Compose
This method allows you to build and run the microservices using Docker Compose.
1. **Navigate to the Project Directory:** Open a terminal and run:
```cd project```
2. **Build the Services:** To build all services except the front-end, execute:
```make up_build```
This command will create Docker images and containers for all microservices defined in the docker-compose.yml file.
3. **Start the Front-End Service:** After building the services, start the front-end service with:
```make start```
This command will run the front-end application, allowing it to communicate with the other services.
4. **Access the Application:** Open your web browser and navigate to:
[http://localhost:8081]
You should see the front-end application running.

### Method 2: Using Docker Swarm
This method utilizes Docker Swarm for deploying the microservices with Caddy as a reverse proxy.
1. **Initialize Docker Swarm**: In the project directory, run the following command to initialize Docker Swarm:
```docker swarm init```
This command will configure your Docker environment for Swarm mode.
2. **Deploy the Stack**: Deploy your application stack by executing:
```docker stack deploy -c swarm.yml myapp```
This command reads the swarm.yml file and sets up the services defined in it.
3. **Access the Application**: Once the deployment is complete, open your web browser and go to:
[http://localhost]
This URL will direct you to the front-end application running through the Caddy reverse proxy.