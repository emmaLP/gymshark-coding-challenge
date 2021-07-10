# Gymshark Coding Challenge

This is split into the following components:

| Component | Description |
| :---: | :---:|
| Backend | Written in Go 1.16 and will run as an AWS Lambda function. This function can also be called by an API|


## Backend
All code can be found in the `backend/` folder of this repo 
### Building
Every commit to this repo will trigger a Github Action to build and test the code.

To build/test locally:
```bash
cd backend/
go mod tidy
go build .
go test .
```

