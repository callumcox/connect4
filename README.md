# Connect 4

## Requirements

This solution was built with the following core requirements

* `go` v1.22.3

## Installation Instructions

If you have `go` installed on your system you can simply run the following `make` command to build and run the solution:

```bash
make run
```

If you do not have `go` installed on your system you can leverage the supplied `Dockerfile` to build and run the solution.  You can do so by running the following command:

```bash
make docker-build
```

or if you do not have `make` on your system you can run:

```bash
docker build -t connect4:latest .
```

## Running the solution

### Build and Run

To run the solution from the local source, you can run:

```bash
make run-dev
```

This will skip the build step.  Alternatively you can run the application directly via:

```bash
go run cmd/main.go
```

### Running Tests

To run unit tests, you can run the following:

```bash
make test
```

or

```bash
go test ./...
```

## Design, Trade-Offs and Future Thought

### Design

This repository is structured in the following manner:

|           | Description                                                                                                                                    |
| --------- | ---------------------------------------------------------------------------------------------------------------------------------------------- |
| /cmd      | Contains the main entrypoint for the application, where the players and game are initialised                                                   |
| /constants| Contains any fixed constants used throughout the game                                                                                          |
| /game     | Contains the main game logic/models                                                                                                            |
| Makefile  | Makefile containing a list of commonly used commands in this project                                                                           |
| Dockerfile| Supports building the application if one does not have `go` installed on the local system                                                      |

The core principals of the design of this application include injecting dependencies where necessary, to better aid with future extensibility and testing especially when dealing with `stdin`/`stdout`.

I've chosen to expose interfaces to components to hide internal workings, particularly when dealing with the `Board` and `cells` so that external consumers have to deal with public interfaces rather than being able to manipulate the board in an uncontrolled manner.

### Trade-Offs

To keep the solution simple/due to time constraints I've made the following trade-offs:

* The win condition checks are currently performed synchronously but the performance impact with a board of this size is negligible.
* The Game struct still takes Player structs as arguments rather than relying on an interface and allowing the consumer to provide alternative implementations.  This would be refactored out if there was more time 

### Future Thoughts

Thoughts about possible future extensions to this solution include:

* Add in additional prompts to allow the players to supply their names/tokens
* Allow for selection of the board size/win condition, defaulting to the values in the constants
* Leverage concurrency when checking the win conditions via goroutines to speed up performance (more beneficial when looking at AI computation)
* Colorize the output to see visible tokens.
* Add an AI Player which implements a GetMove function that does not take from `stdin` but rather simply computes the next move (minmax algorithm).
  * Extend to include selection of Human vs AI Player
* Extend to allow for play via a GUI.