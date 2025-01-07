# Distributed Task Processor

This project demonstrates a simple distributed task processor in Go. It consists of workers that process tasks concurrently.

## Structure

- `Task`: Represents a task with an ID, name, description, and processing time.
- `Worker`: Represents a worker that processes tasks from a task channel.
- `main.go`: The main entry point of the application.

## How It Works

1. The main function initializes a set number of workers and tasks.
2. Tasks are generated and sent to a task channel.
3. Workers listen to the task channel and process tasks concurrently.
4. A stop channel is used to signal workers to stop processing.

## Running the Application

To run the application, use the following command:

```sh
go run main.go -tasks=<number_of_tasks> -workers=<number_of_workers>
```

## Configuration

- `-tasks`: The number of tasks to be processed (default is 20).
- `-workers`: The number of workers available to process tasks (default is 3).

## Example Output

```
Worker 0 is processing task 0 for 3s
Worker 1 is processing task 1 for 2s
Worker 2 is processing task 2 for 1s
Task 0 finished by worker 0
Task 1 finished by worker 1
Task 2 finished by worker 2
...
```

## Concurrency

This application demonstrates the use of goroutines and channels in Go to achieve concurrent task processing.

## License

This project is licensed under the MIT License.
