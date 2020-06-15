# manhattanite
> Like most native Manhattanites, she knew them by heart, could grab one that would take her to the appropriate location, if not always by the quickest route. 


*I, Robot - Isaac Asimov*.

[![Go Report Card](https://goreportcard.com/badge/github.com/rafaeltardivo/manhattanite)](https://goreportcard.com/report/github.com/rafaeltardivo/manhattanite)

Manhatannite cartesianservice provides a simple API with a single endpoint that returns a list of points that are within [manhattan distance](https://xlinux.nist.gov/dads/HTML/manhattanDistance.html) (`x`, `y` `distance`), based on pre-loaded two-dimentional points from [`data.json`](data/data.json).

## Table of Contents

- [Technology](#technology)
- [How it works](#how-it-works)
	- [Architecture](#architecture)
	- [Configurations](#configurations)
- [Developing](#developing)
    - [Installing](#installing)
    - [API](#using-the-api)
- [Logs](#logs)
- [Final Considerations](#final-considerations)
    

## Technology
- [Golang](https://www.python.org/) 1.14
- [logrus](github.com/sirupsen/logrus) v1.6.0
- [gomega](github.com/onsi/gomega) v1.10.1
- [Docker](https://www.docker.com/) 19.03.6



## How it Works

### Architecture
Manhattanite cartesian service is designed as a hexagonal (a.k.a. Ports and Adapters) microservice:

![alt text](https://user-images.githubusercontent.com/4626533/84607176-3a944b00-ae82-11ea-8f78-e96f0b51d257.jpg)

The domain logic is separated from data sources (repositories) and presentation formats (serializers), making it easy to switch from local file to database, or even from JSON to XML.

### Configurations

Following [twelve factor app](https://12factor.net/) rules, all configurations is strictly separated from the application:

|     Parameter                |     Description  |      Value        |
|------------------------------|------------------|-------------------|
| `HTTP_SERVER_PORT`           |    Server port   |  8080             |
| `POINTS_FILE_RELATIVE_PATH`  |  Data file path  |  `data.json`      |  



## Developing

### Installing
1 - Clone the project
```
git clone git@github.com:rafaeltardivo/manhattanite.git  
```
2 - Build the application:  
```
make build
```  
4 - Run the application:  
```  
make up
```  

**OBS**
 - All unit tests are executed during the building proccess;
  - Everything is dockerized with multi-stage build.

### Using the API

Single endpoint:

|  Resource                           | Port  | HTTP Method |  
|-------------------------------------|-------|-------------|  
| `http://localhost:8080/api/points`  | `8080`|  `GET`      |

With mandatory query parameters:

|  Parameter  | type   | example |  
|-------------|--------|---------|  
|     `x`     | integer|  -30    |
|     `y`     | integer|  -38    |
|  `distance` | integer|   10    |


Request example:


```bash
curl --request GET \
  --url 'http://0.0.0.0:8080/api/points?x=-30&y=-38&distance=5'
```

Response: 

```json
[
  {
    "x": -30,
    "y": -38
  }
]
```

Possible errors:

 - `HTTP 400` for bad requests with missing and/or invalid query parameters;
 - `HTTP 403` for HTTP methods not allowed (all except `GET`);
 - `HTTP 404` for no point(s) found within range).


### Logs
Logs are pretty descriptive. They come in two levels: `info` and `error`. Here is one example of log output generated by a request:

```bash
{"level":"info","msg":"loading key POINTS_FILE_RELATIVE_PATH from environment","time":"2020-06-14T23:47:46Z"}
{"level":"info","msg":"loading key HTTP_SERVER_PORT from environment","time":"2020-06-14T23:47:46Z"}
{"level":"info","msg":"server is ready to accept connections on port 8080","time":"2020-06-14T23:47:46Z"}
{"level":"info","msg":"validating request parameters map[distance:[5] x:[-30] y:[-38]]","time":"2020-06-14T23:47:51Z"}
{"level":"info","msg":"querying for points within distance x=-30 y=-38 distance=5 ","time":"2020-06-14T23:47:51Z"}
{"level":"info","msg":"sorting and returning 1 point(s) found within distance","time":"2020-06-14T23:47:51Z"}
```

### Final Considerations
- This project is not "production ready". For that, we would need a more robust HTTP multiplexer;
- The `data.json` is read all at once, which is dangerous for large files who may not fit on your available RAM. For better scalability, we would need to stream the reading process;
- The sorting is done using go sorting interface, which uses [quicksort](https://golang.org/src/sort/sort.go?s=5416:5441#L206) to sort data;
- Since the queries are deterministic, it could have a cache model (invalidated everytime a new `data.json` was loaded);
- There's probably a way to optimize the search for points by ordering the file.