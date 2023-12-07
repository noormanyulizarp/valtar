# Valtar - Go API Project

## Description

Valtar is a RESTful API developed in Go, designed to calculate the area of a rectangle. This project demonstrates a clean and efficient way to structure a Go-based API with modular code organization.

## Features

- RESTful endpoint to calculate the area of a rectangle.
- Centralized handling of API requests.
- Efficient JSON response handling.
- Organized project structure for maintainability and scalability.

## Project Structure
valtar/
├── cmd/
│   └── server/
│       └── main.go       // Main application entry point
├── internal/
│   ├── api/
│   │   ├── handlers/
│   │   │   ├── apiInfoHandler.go
│   │   │   ├── rectangleAreaHandler.go
│   │   │   └── handler.go
│   │   └── utils/
│   │       └── utils.go
│   └── calculator/
│       └── calculator.go
└── go.mod


## Getting Started

### Prerequisites

- Go 1.19 or higher

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/valtar.git


# Fish History Management Script

This Bash script is designed to help you manage your Fish shell command history. It provides functionality to export and import Fish history, making it convenient to version control and transfer your command history between sessions.

## Usage

### Export Fish History

Run the following command to create a backup of your Fish history:

```bash
./manage_fish_history.sh create
