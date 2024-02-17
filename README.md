# :trumpet: bugle: send messages from command line

Bugle is a command-line tool designed for sending messages quickly and efficiently. Built with Go and packaged using Docker, it offers a straightforward way for users to integrate messaging capabilities into their scripts or automation workflows.

## Features

- Send messages from the command line
- Easy integration with scripts and automation tasks
- Lightweight and fast

## Installation

To get started with Bugle, clone this repository and build the Docker container:

### Golang
```bash
go install github.com/debeshmandal/bugle
./bugle --body="Hello!"
```

### Docker
```bash
git clone https://github.com/debeshmandal/bugle.git
cd bugle
docker build -t bugle .
docker run bugle --body="Hello!"
```

### Manual Build (Golang)
```bash
git clone https://github.com/debeshmandal/bugle.git
cd bugle
go install .
```

## Usage
After installation, use Bugle by running:
```bash
Copy code
docker run bugle [options]
Replace [options] with your specific command line options for sending messages.
```

## Contributing
Contributions to Bugle are welcome! Please refer to the contributing guidelines for more information on how to get involved.

License
Bugle is licensed under the GPL-3.0 license. See the LICENSE file for details.