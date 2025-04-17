# Microservices E-commerce System

This project implements a microservices architecture for an e-commerce system using Go, gRPC, and RESTful APIs.

## Architecture

The system consists of the following services:

1. **API Gateway**: Entry point for clients, exposes RESTful endpoints and forwards requests to backend services via gRPC
2. **Inventory Service**: Manages products and categories via gRPC
3. **Order Service**: Manages customer orders
4. **User Service**: Handles user authentication and profiles

## Running the Services

1. Start the Inventory Service:
   ```
   cd inventory_service/cmd
   go run main.go
   ```

2. Start the API Gateway:
   ```
   cd api_gateway/cmd
   go run main.go
   ```

By default, the API Gateway runs on port 8080, and the Inventory Service's gRPC server runs on port 50051.

## API Documentation

### Authentication

Most endpoints require authentication via a JWT token. Include the token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

### Product Management

#### Create a Product

Creates a new product in the inventory.

- **Method**: POST
- **URL**: `http://localhost:8080/inventory/products`
- **Headers**:
  - Content-Type: application/json
  - Authorization: Bearer <your_jwt_token>
- **Request Body**:
  ```json
  {
    "name": "Smartphone X",
    "description": "Latest smartphone with advanced features",
    "category_id": "1",
    "price": 899.99,
    "stock": 50
  }
  ```
- **Response**:
  ```json
  {
    "id": "product_id",
    "name": "Smartphone X",
    "description": "Latest smartphone with advanced features",
    "category_id": "1",
    "price": 899.99,
    "stock": 50,
    "created_at": "2023-04-18T12:00:00Z",
    "updated_at": "2023-04-18T12:00:00Z"
  }
  ```
- **Notes**: The `category_id` must reference an existing category, otherwise the API will return a 404 error.

#### Get a Product

Retrieves a specific product by ID.

- **Method**: GET
- **URL**: `http://localhost:8080/inventory/products/{product_id}`
- **Headers**:
  - Authorization: Bearer <your_jwt_token>
- **Response**:
  ```json
  {
    "id": "product_id",
    "name": "Smartphone X",
    "description": "Latest smartphone with advanced features",
    "category_id": "1",
    "price": 899.99,
    "stock": 50,
    "created_at": "2023-04-18T12:00:00Z",
    "updated_at": "2023-04-18T12:00:00Z"
  }
  ```

#### Update a Product

Updates an existing product.

- **Method**: PUT
- **URL**: `http://localhost:8080/inventory/products/{product_id}`
- **Headers**:
  - Content-Type: application/json
  - Authorization: Bearer <your_jwt_token>
- **Request Body**:
  ```json
  {
    "name": "Smartphone X Pro",
    "description": "Updated description with new features",
    "category_id": "1",
    "price": 999.99,
    "stock": 45
  }
  ```
- **Response**:
  ```json
  {
    "id": "product_id",
    "name": "Smartphone X Pro",
    "description": "Updated description with new features",
    "category_id": "1",
    "price": 999.99,
    "stock": 45,
    "created_at": "2023-04-18T12:00:00Z",
    "updated_at": "2023-04-18T13:00:00Z"
  }
  ```
- **Notes**: The `category_id` must reference an existing category, otherwise the API will return a 404 error.

#### Delete a Product

Deletes a product from the inventory.

- **Method**: DELETE
- **URL**: `http://localhost:8080/inventory/products/{product_id}`
- **Headers**:
  - Authorization: Bearer <your_jwt_token>
- **Response**:
  ```json
  {
    "message": "Product deleted successfully"
  }
  ```

#### List Products

Lists products with optional filtering and pagination.

- **Method**: GET
- **URL**: `http://localhost:8080/inventory/products`
- **Headers**:
  - Authorization: Bearer <your_jwt_token>
- **Query Parameters**:
  - `name`: Filter by product name (optional)
  - `category_id`: Filter by category ID (optional)
  - `min_price`: Filter by minimum price (optional)
  - `max_price`: Filter by maximum price (optional)
  - `page`: Page number for pagination (default: 1)
  - `limit`: Number of items per page (default: 10)
- **Response**:
  ```json
  [
    {
      "id": "product_id_1",
      "name": "Smartphone X",
      "description": "Latest smartphone with advanced features",
      "category_id": "1",
      "price": 899.99,
      "stock": 50,
      "created_at": "2023-04-18T12:00:00Z",
      "updated_at": "2023-04-18T12:00:00Z"
    },
    {
      "id": "product_id_2",
      "name": "Laptop Pro",
      "description": "High performance laptop",
      "category_id": "1",
      "price": 1299.99,
      "stock": 20,
      "created_at": "2023-04-18T12:30:00Z",
      "updated_at": "2023-04-18T12:30:00Z"
    }
  ]
  ```

### Category Management

#### Create a Category

Creates a new product category.

- **Method**: POST
- **URL**: `http://localhost:8080/inventory/categories`
- **Headers**:
  - Content-Type: application/json
  - Authorization: Bearer <your_jwt_token>
- **Request Body**:
  ```json
  {
    "name": "Electronics",
    "description": "Electronic devices and gadgets"
  }
  ```
- **Response**:
  ```json
  {
    "id": "category_id",
    "name": "Electronics",
    "description": "Electronic devices and gadgets"
  }
  ```

#### Get a Category

Retrieves a specific category by ID.

- **Method**: GET
- **URL**: `http://localhost:8080/inventory/categories/{category_id}`
- **Headers**:
  - Authorization: Bearer <your_jwt_token>
- **Response**:
  ```json
  {
    "id": "category_id",
    "name": "Electronics",
    "description": "Electronic devices and gadgets"
  }
  ```

#### Update a Category

Updates an existing category.

- **Method**: PUT
- **URL**: `http://localhost:8080/inventory/categories/{category_id}`
- **Headers**:
  - Content-Type: application/json
  - Authorization: Bearer <your_jwt_token>
- **Request Body**:
  ```json
  {
    "name": "Modern Electronics",
    "description": "Latest electronic devices and gadgets"
  }
  ```
- **Response**:
  ```json
  {
    "id": "category_id",
    "name": "Modern Electronics",
    "description": "Latest electronic devices and gadgets"
  }
  ```

#### Delete a Category

Deletes a category.

- **Method**: DELETE
- **URL**: `http://localhost:8080/inventory/categories/{category_id}`
- **Headers**:
  - Authorization: Bearer <your_jwt_token>
- **Response**:
  ```json
  {
    "message": "Category deleted successfully"
  }
  ```
- **Notes**: If the category is referenced by existing products, the deletion might fail.

#### List Categories

Lists all available categories.

- **Method**: GET
- **URL**: `http://localhost:8080/inventory/categories`
- **Headers**:
  - Authorization: Bearer <your_jwt_token>
- **Response**:
  ```json
  [
    {
      "id": "category_id_1",
      "name": "Electronics",
      "description": "Electronic devices and gadgets"
    },
    {
      "id": "category_id_2",
      "name": "Clothing",
      "description": "Apparel and accessories"
    }
  ]
  ```

## Health Check

Check if services are running properly.

- **Method**: GET
- **URL**: `http://localhost:8080/health`
- **Response**:
  ```json
  {
    "status": "ok",
    "time": "2023-04-18T15:00:00Z"
  }
  ```

## Error Responses

The API returns appropriate HTTP status codes and error messages:

- **400 Bad Request**: Invalid request format or parameters
- **401 Unauthorized**: Missing or invalid authentication token
- **404 Not Found**: Requested resource not found
- **500 Internal Server Error**: Server-side error

Example error response:
```json
{
  "error": "Category with ID 123 not found"
}
```

## Authentication

### Register a New User

- **Method**: POST
- **URL**: `http://localhost:8080/users/register`
- **Headers**:
  - Content-Type: application/json
- **Request Body**:
  ```json
  {
    "username": "newuser",
    "password": "securepassword",
    "email": "user@example.com",
    "full_name": "New User"
  }
  ```
- **Response**:
  ```json
  {
    "id": "user_id",
    "username": "newuser",
    "email": "user@example.com",
    "full_name": "New User",
    "created_at": "2023-04-18T12:00:00Z"
  }
  ```

### Login

- **Method**: POST
- **URL**: `http://localhost:8080/users/login`
- **Headers**:
  - Content-Type: application/json
- **Request Body**:
  ```json
  {
    "username": "newuser",
    "password": "securepassword"
  }
  ```
- **Response**:
  ```json
  {
    "token": "jwt_token_string",
    "expire_at": "2023-04-19T12:00:00Z",
    "user": {
      "id": "user_id",
      "username": "newuser",
      "email": "user@example.com",
      "full_name": "New User",
      "created_at": "2023-04-18T12:00:00Z"
    }
  }
  ``` 