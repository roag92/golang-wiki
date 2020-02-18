GoLang Wiki
=======

A PoC using Docker to build and develop an application from a container. This application is based on this article: https://golang.org/doc/articles/wiki/.

 > Note: This project is using DevContainer plugin on VSCode for more information check: https://code.visualstudio.com/docs/remote/containers. 

# Requirements

 - Docker

# Build

```bash
docker build -t roag92/golang-wiki:dockerfile ./
```

# Run

```bash
docker run --rm -it -p 8080:8080 -v ./tmp:/golang-wiki/tmp roag92/golang-wiki:dockerfile
```

# Resources

 - https://code.visualstudio.com/docs/remote/containers
 - https://golang.org/doc/articles/wiki/
 