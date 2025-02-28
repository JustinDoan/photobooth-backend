# Photobooth Backend

## Overview
The Photobooth Backend is a Go-based server application designed to handle image processing and management for a photobooth application. It provides APIs to process images and maintain a list of images with their metadata.

## Features
- **Image Processing**: Backend logic to handle image processing tasks.
- **Image Metadata Management**: Stores metadata for each image in a database.
- **API Endpoints**:
  - `POST /api/images/`: Process a new image and store its metadata.
  - `GET /api/images/`: Retrieve a list of all stored images.
- **CORS Support**: Configured to allow cross-origin requests.

## Repository Structure
- `main.go`: Entry point of the application, initializing the server and database.
- `server.go`: Defines the server's HTTP handlers and routes.
- `controllers/`: Contains business logic (`ImageController`) for handling image-related operations.
- `models/`: Contains data models (`Image`) representing the application's data structures.
- `services/`: Business logic services used by controllers.
- `repository/`: Data access layer for interacting with the database.
- `handlers/`: HTTP request handlers to process API requests.
- `utils/`: Utility functions and helpers.

## Setup Instructions
1. **Prerequisites**:
   - Go 1.16 or higher installed.
   - A running instance of the database compatible with the configuration in `config/`.

2. **Installation**:
   - Clone the repository: `git clone <repository-url>`
   - Navigate to the project directory: `cd photobooth-backend`
   - Install dependencies: `go mod download`

3. **Running the Server**:
   - Start the server: `go run main.go`
   - The server will be available at `http://localhost:8080`

## Development Notes
- **Extending the API**: New endpoints can be added by expanding the `controllers` and `handlers` packages.
- **Database Migrations**: Ensure database schema is up-to-date with `gorm` migrations.
- **Configuration**: Check the `config/` directory for environment-specific configurations.

## Contribution
Contributions are welcome! Please fork the repository and submit a pull request for any changes.

## License
This project is licensed under the MIT License.