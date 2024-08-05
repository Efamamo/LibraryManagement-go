# Library Management System

## Overview
This is a simple console-based library management system implemented in Go. It allows you to manage books and members of a library, including adding, removing, borrowing, and returning books.

## Folder Structure
- `main.go`: Entry point of the application.
- `controllers/`: Contains the code for console interaction.
- `models/`: Defines data structures for Book and Member.
- `services/`: Contains business logic and data manipulation functions.
- `docs/`: Contains documentation.

## How to Run
1. Build the application:
    ```bash
    go build -o library_management main.go
    ```

2. Run the application:
    ```bash
    ./library_management
    ```

## Features
- Add a new book
- Remove an existing book
- Borrow a book
- Return a book
- List all available books
- List all borrowed books by a member

## Error Handling
The system provides error messages for:
- Non-existent books or members
- Books already borrowed or not borrowed
