# API Documentation for AuditCue

## Overview
This document provides an overview of the API endpoints available in the AuditCue application. Each endpoint includes details about the request method, URL, parameters, and response format.

## Endpoints

### 1. User Signup
- **Method**: POST
- **URL**: `/api/signup`
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```
- **Response**:
  - **Success (201 Created)**:
    ```json
    {
      "message": "User created successfully",
      "userId": "12345"
    }
    ```
  - **Error (400 Bad Request)**:
    ```json
    {
      "error": "Invalid input data"
    }
    ```

### 2. User Login
- **Method**: POST
- **URL**: `/api/login`
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```
- **Response**:
  - **Success (200 OK)**:
    ```json
    {
      "message": "Login successful",
      "token": "jwt.token.here"
    }
    ```
  - **Error (401 Unauthorized)**:
    ```json
    {
      "error": "Invalid credentials"
    }
    ```

### 3. OAuth Authentication
- **Method**: GET
- **URL**: `/api/auth/oauth`
- **Response**:
  - **Redirects to OAuth provider** for authentication.

### 4. Get User Profile
- **Method**: GET
- **URL**: `/api/user/profile`
- **Headers**:
  - `Authorization: Bearer <token>`
- **Response**:
  - **Success (200 OK)**:
    ```json
    {
      "userId": "12345",
      "email": "user@example.com"
    }
    ```
  - **Error (401 Unauthorized)**:
    ```json
    {
      "error": "Unauthorized access"
    }
    ```

## Usage Examples
- To sign up a new user, send a POST request to `/api/signup` with the required fields.
- To log in, send a POST request to `/api/login` with the user's credentials.
- Use the provided token in the `Authorization` header for protected routes.

## Conclusion
This API documentation serves as a guide for developers to understand and interact with the AuditCue application. Ensure to handle errors appropriately and validate input data for a smooth user experience.