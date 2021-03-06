GoLang Wiki
=======

A PoC using Docker to build and develop an application from a container. This application is based on this article: https://golang.org/doc/articles/wiki/.

 > Note: This project is using DevContainer plugin on VSCode for more information check: https://code.visualstudio.com/docs/remote/containers. 

# Requirements

 - Docker
 - docker-compose

# Build

```bash
docker-compose build
```

# Setting

Copy the `docker-compose.override.yml.dist` file to `docker-compose.override.yml`.

 ```bash
cp docker-compose.override.yml.dist docker-compose.override.yml
 ```

# Run

```bash
docker-compose up -d
```

# How it works?

You could navigate to http://localhost:8080/ to list all the pages stored into `tmp` folder.

 - `http://localhost:8080/edit/MyNewTitle`.- Edit an existent or new page 
 - `http://localhost:8080/save/MyNewTitle`.- Save a new page
 - `http://localhost:8080/view/MyNewTitle`.- View a page

```bash
curl -d 'body=Hello World!' http://localhost:8080/save/MyNewTitle
```

> Note: Every path after edit, save or view must satisfies this regular expresion: `^/(edit|save|view)/([a-zA-Z0-9]+)$`

# Resources

 - https://code.visualstudio.com/docs/remote/containers
 - https://golang.org/doc/articles/wiki/
