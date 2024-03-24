Golang_api_server is a Golang-based REST API server that provides functionality for managing shopping items. This project includes endpoints for retrieving a list of items, creating new items, and deleting items. It is designed to showcase API development in Golang and can be easily downloaded and run on a local machine for testing and learning purposes.

Requirements Go programming language installed on your machine. You can follow the official Go installation guide here. Go modules enabled. Run go mod init server in the project directory to initialize modules. Clone the repository to your local machine using git clone https://github.com/developer-jos/golang_api_server. Getting Started Navigate to the project directory. Run the following command to start the server: go run main.go The server will start running on port 8080. Access the API endpoints using a web browser or tools like cURL or Postman. API Endpoints GET /items: Retrieve a list of shopping items. POST /items: Create a new shopping item. DELETE /items/{id}: Delete a shopping item by ID. Contributing Feel free to contribute to this project by submitting pull requests. For major changes, please open an issue first to discuss the changes you would like to make.