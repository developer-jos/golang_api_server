Golang_api_server is a Golang-based REST API server that provides functionality for managing shopping items. This project includes endpoints for retrieving a list of items, creating new items, and deleting items. It is designed to showcase API development in Golang and can be easily downloaded and run on a local machine for testing and learning purposes.

Requirements Go programming language installed on your machine. You can follow the official Go installation guide here. Go modules enabled. Run go mod init server in the project directory to initialize modules. Clone the repository to your local machine using git clone https://github.com/developer-jos/golang_api_server. Getting Started Navigate to the project directory. Run the following command to start the server: go run main.go The server will start running on port 8080. Access the API endpoints using a web browser or tools like cURL or Postman. API Endpoints GET /items: Retrieve a list of shopping items. POST /items: Create a new shopping item. DELETE /items/{id}: Delete a shopping item by ID. Contributing Feel free to contribute to this project by submitting pull requests. For major changes, please open an issue first to discuss the changes you would like to make.

Swagger documentation for the provided API endpoints:

YAML
swagger: "2.0"
info:
  description: "Shopping List API"
  title: "Shopping List"
  version: "1.0"
servers:
  - url: "http://localhost:8080"
paths:
  /items:
    get:
      summary: "Get all shopping items"
      description: "Returns a list of all shopping items in the system."
      responses:
        200:
          description: "OK"
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/Item"
        500:
          description: "Internal Server Error"
  /items/create:
    post:
      summary: "Create a new shopping item"
      description: "Creates a new shopping item and adds it to the list."
      requestBody:
        required: true
        content:
          application/json:
            schema:
              $ref: "#/components/schemas/Item"
      responses:
        201:
          description: "Created"
        400:
          description: "Bad Request"
          content:
            text/plain:
              example: "Error decoding request body: ..."
  /items/get:
    get:
      summary: "Get a specific shopping item by ID"
      description: "Retrieves a specific shopping item from the list based on its ID."
      parameters:
        - in: path
          name: id
          required: true
          type: string
          format: uuid
      responses:
        200:
          description: "OK"
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/Item"
        404:
          description: "Not Found"
          content:
            text/plain:
              example: "item not found"
  /items/delete:
    delete:
      summary: "Delete a shopping item by ID"
      description: "Removes a specific shopping item from the list based on its ID."
      parameters:
        - in: path
          name: id
          required: true
          type: string
          format: uuid
      responses:
        204:
          description: "No Content"
        404:
          description: "Not Found"
          content:
            text/plain:
              example: "item not found"
components:
  schemas:
    Item:
      type: object
      properties:
        id:
          type: string
          format: uuid
          description: "The unique identifier of the shopping item."
        name:
          type: string
          description: "The name of the shopping item."
      required:
        - name
Use code with caution.
This Swagger definition describes the following:

Info: General information about the API, including title, description, and version.
Servers: The base URL where the API is running (localhost:8080 in this case).
Paths: Each path represents an API endpoint.
/items (GET): Retrieves a list of all shopping items.
/items/create (POST): Creates a new shopping item.
/items/get (GET): Gets a specific shopping item by ID.
/items/delete (DELETE): Deletes a shopping item by ID.
Responses: The possible HTTP status codes returned by each endpoint and their descriptions.
Request Body: The expected format of the request body for the POST endpoint (/items/create).
Parameters: The required parameter for the GET and DELETE endpoints by ID (/items/get and /items/delete).
Schemas: The definition of the Item object schema, including its properties and required fields.
Remember to replace logRequest with your actual implementation for logging request details. You can further enhance this Swagger definition by adding additional information like security, examples, and tags for each endpoint.

Sources
github.com/ArturPromon/Rakenduste-programeerimine
github.com/kasra73/todo-api
github.com/Pocik-lab/PIUS_Lab3
