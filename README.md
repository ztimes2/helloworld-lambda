# AWS Lambda + Go
This repo serves as a place to keep learnings about using AWS Lambda together with Go.

## Deploying as container images
Initially, the only option to ship a Go application as a Lambda function was via a ZIP archive (which is not that great 🤢). Nowadays, AWS provides an option to make use of container images as Lambda functions.

### Requirements
* Your custom image needs to be based on `amazon/aws-lambda-go` (DockerHub) or `gallery.ecr.aws/lambda/go` (ECR). 
* Move your Go binary file to the specific directory expected by the base image (the directory path can be taken from `LAMBDA_TASK_ROOT` env variable).
* Add a `CMD ["app"]` step and substitute `app` with the actual name of your Go binary file. 

See an example of a Dockerfile based on the DockerHub image [here](./Dockerfile). 
