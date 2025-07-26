# 🐮 Go-Cowsay

A simple implementation of the classic `cowsay` command in Go. This program displays a cow with a message provided by the user. 🐄

## ✨ Features

- Display a cow with a custom message.
- Option to display Tux instead of a cow. 🐧
- Formats the message to fit within a dialogue box.

## 🚀 Installation

To build the executable, you need to have Go installed on your system.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/go-cowsay.git
    cd go-cowsay
    ```

2.  **Build the executable:**
    ```bash
    go build
    ```

## Usage

To run the program, use the following command:

```bash
./go-cowsay "your message here"
```

### Options

-   `-f <character>`: Choose the character to display.
    -   `cow`: (Default) Displays a cow.
    -   `tux`: Displays Tux.

### Examples

**Default cow:**

```bash
./go-cowsay "Hello, I am a cow."
```

**Output:**

```
 ____________________
< Hello, I am a cow. >
 --------------------
         \  ^__^
          \ (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

**Tux:**

```bash
./go-cowsay -f tux "Hello, I am Tux."
```

**Output:**

```
 __________________
< Hello, I am Tux. >
 ------------------
 \
  \
        .--.
       |o o |
       |:_/ |
      //    \\
     (|     | )
    /'\\_   _/'\\
    \\___)=(___/
```

## 🛠️ Development

### Prerequisites

-   Go 1.18 or higher.

### Running Tests

To run the tests for this project, use the following command:

```bash
go test
```

### Project Structure

-   `main.go`: The main application logic.
-   `main_test.go`: Tests for the main application.
-   `go.mod`: Go module definition.
-   `go.sum`: Checksums for dependencies.
-   `.gitignore`: Specifies files to be ignored by Git.

## 🤝 Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.
