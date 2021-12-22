# Notfound

Example app used to test docker and kubernetes services.
It's a simple go app that runs on port 4400 and renders not found page.

## Usage

    ./notfound

## Docker usage

    docker build . -t andrzejtrzaska/notfound
    docker run -p 4400:4400 andrzejtrzaska/notfound
    curl -v http://localhost:4400

## License

MIT
