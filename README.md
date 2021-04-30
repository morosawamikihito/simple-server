# Overview

Simple servers for trying out some technologies.

Example docker-compose.yml is following.
And execute `docker-compose up -d` only.

After that, check `http://localhost:8081/`, `http://localhost:8082/`, `http://localhost:8083/`.

```
version: '3.9'

services:
  server1:
    image: morosawamikihito/simple-server:server1
    ports:
    - 8081:8080
  server2:
    image: morosawamikihito/simple-server:server2
    ports:
    - 8082:8080
  server3:
    image: morosawamikihito/simple-server:server3
    ports:
    - 8083:8080
```

## endpoints

|server|path|
|---|---|
|server1|/|
|server1|/server1/hello|
|server1|/server1/hey|
|server2|/|
|server2|/server2/hello|
|server2|/server2/hey|
|server3|/|
|server3|/server3/hello|
|server3|/server3/hey|


## Links

- [dockerhub](https://hub.docker.com/r/morosawamikihito/simple-server)