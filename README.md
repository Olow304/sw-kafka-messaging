# sw-kafka-messaging

This application is a distributed system comprising multiple services, including Kafka, Zookeeper, a React front-end, a Node.js back-end, and a Go server. The services are defined using Docker Compose. Follow the instructions below to run the application.

### Prerequisites
Install <strong>Docker</strong> and <strong>Docker Compose</strong>. Ensure you have the latest version of Docker Compose installed, as the configuration file format version is 3.8.


### Steps to Run the Application 

- Download or clone this repository to your local machine.
- Open a terminal and navigate to the folder containing the Docker Compose configuration file (docker-compose.yml).
- Before running the application, ensure the create-topic.sh script is executable. Run the following command to grant execute permissions:

```
  chmod +x create-topic.sh
```
- To start the application, run the following command in the terminal:
```
  docker-compose up
```
This command will build and run all the services defined in the docker-compose.yml file. Docker Compose will download the required images and create the necessary containers for each service.

- Once all services are up and running, access the React front-end by opening a web browser and navigating to http://localhost:3000.
- The Node.js back-end service will be available on http://localhost:3535, and the Go server will be available on http://localhost:8080.
- To send a message using the /send-message endpoint via curl, open a terminal and run the following command:
```
  curl -X POST -H "Content-Type: application/json" -d '{"message": "Your message here"}' http://localhost:3535/send-message
```
or you can use PostMan

### Stopping the Application
- To stop the application and remove the containers, networks, and volumes defined in the docker-compose.yml file, run the following command in the terminal:
```
  docker-compose down
```
