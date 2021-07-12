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

## Deploy
The deployment is handled through terraform scripts. 

Terraform scripts configure the following:
1. Go Lambda with the pack calc deployed
1. Node Lambda used as an authoriser for the API Gateway
1. API Gateway integrated with the pack calc lambda
1. A randomly generated token stored as a secure string in `ParamterStore` which is used as the auth token

### API 
The API only supports on endpoint as a POST request.

The API body is JSON and should look like this example:
```json
{
  "items":501, 
  "packSizes": [250,500,1000,2000,5000]
}
```

Example Request:

```bash
curl -X POST --data '{"items":501, "packSizes": [250,500,1000,2000,5000]}' \
 https://jcwz2ki77i.execute-api.eu-west-2.amazonaws.com/calculate-packs -H "Authorization: qD449xM9k0nvK@_f"
```
**The auth token is just an example**


## Notes/ Considerations
1. Unfortunately, I did not have enough time to implement a frontend as I was learning golang while implementing the backend
1. GitHub Actions have been used to do the following:
    1. Build and Test the Go Lang Code
    1. Terraform script validation (`fmt`, `init` and `validate`)
    1. On a PR merge, a GitHub release is created with the compiled Go code attached
    1. Once a GitHub release is published, a deployment workflow is triggered.
        1. This workflow will handle the deployment of the resources
