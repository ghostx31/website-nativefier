# Nativefier Downloader 

[Nativefier](https://github.com/nativefier/nativefier) is a node cli tool which can convert any website into a native looking app using Electron. This aim of this repository is to provide this cli tool as a website. 

The repository needs to be built as a docker container to be used. 

### Technologies used:
- Node.js 
- Golang 
  - Echo server 
- Docker 

### How to build the docker container 

- From the root of the repository run the command:
```bash
docker build --network=host -t nativefier-downloader:latest .
```

- To run the built docker container: 
```bash
docker run -p 1323:1323 nativefier-downloader:latest
```

- Now browse to `localhost:1323` to get to the page. 

### What works currently

- Currently only building the Linux and macOS packages work. The Windows build requires wine to be installed on the docker container and that's a long term goal since it needs more optimization. 
- Front end is just a simple HTML page. The plan is to make a better front-end overall. 
- There is nearly no error handling currently or user OS checking for automation. This would require some work on the `server.go` and I aim to complete it. 

