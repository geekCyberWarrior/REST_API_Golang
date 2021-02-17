# Student REST API in Golang
A RESTful API example for simple todo application with Go

It is a just simple tutorial or example for making simple RESTful API with Go using gorilla/mux (A nice mux library)

# Installation & Run

```bash
# Download this project
go get github.com/geekCyberWarrior/REST_API_Golang
```

Before running API server, you should set the database config with yours or set the your database config with the right values in [main.go](https://github.com/geekCyberWarrior/REST_API_Golang/blob/main/main.go)

```bash
# Build and Run
cd REST_API_Golang
go build
./students

# API Endpoint : http://127.0.0.1:12345
```

# API

### /students
* `GET`: Get all students' list
* `POST`: Create a new Student

### /students/:name
* `GET`: Get specific student by name

### /students/:name
* `PUT`: Update specific student by name

### /students/:name
* `DELETE`: Delete specific student by name
