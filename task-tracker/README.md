# Task Tracker CLI

A simple command-line interface (CLI) application to track and manage your tasks and to-do list using Go.

## Requirements
- Go programming language installed on your machine.
- Git for version control.

## Project Features
- Add, update, and delete tasks.
- Mark tasks as in progress or done.
- List all tasks or filter by status (todo, in-progress, done).

## Getting Started
1. Clone the repository:
   ```bash
   git clone https://github.com/PecheslavVenik/backend-projects.git
   ```

2. Navigate to the project directory:
   ```bash
   cd task-tracker
   ```

3. Build the CLI application:
   ```bash
   go build -o task-tracker
   ```

4. Run the application:
   ```bash
   ./task-tracker list
   ```

## Usage
To use the task tracker, you can execute commands like:
- Adding a new task:
  ```bash
  ./task-tracker add "Buy groceries"
  ```

- Listing all tasks:
  ```bash
  ./task-tracker list
  ```

- Listing tasks by status:
  ```bash
  ./task-tracker list [todo, in-progress, done]
  ```

- Updating a task:
  ```bash
  ./task-tracker update 1 "New task description"
  ```

- Deleting a task:
  ```bash
  ./task-tracker delete 1
  ```

- Marking a task as done:
  ```bash
  ./task-tracker mark-done 1
  ```

- Marking a task as in progress:
  ```bash
  ./task-tracker mark-in-progress 1
  ```

## Project Page
Visit the project page for more information: [Project URL](https://roadmap.sh/projects/task-tracker)
