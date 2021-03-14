# AWS Lambda + Go
This repo serves as a place to keep learnings about using AWS Lambda together with Go.

## Deploying as container images
Initially, the only option to ship a Go application as a Lambda function was via a ZIP archive (which is not that great 🤢). Nowadays, AWS provides an option to make use of container images as Lambda functions.

### Requirements
* Your custom image needs to be based on `amazon/aws-lambda-go` (DockerHub) or `gallery.ecr.aws/lambda/go` (ECR). 
* Move your Go binary file to the specific directory expected by the base image (the directory path can be taken from `LAMBDA_TASK_ROOT` env variable).
* Add a `CMD ["app"]` step and substitute `app` with the actual name of your Go binary file. 

See an example of a Dockerfile based on the DockerHub image [here](./Dockerfile). 

## Networking
By default a Lambda function gets attached to a mysterious VPC that is fully managed and controlled by AWS with the following characteristics:
- CIDR Blocks: Unknown and doesn't matter.
- Ingress Rules: None, since external resources can never 'call' a Lambda function.
- Egress Rules: All. It doesn't block outbound access. However it is stateful, so responses would be allowed back in.

We won't have control or visibility about this VPC. If you wish to have more control over the VPC in which Lambda runs, you can create your own VPC and configure the Lambda function to run in that VPC.

Please note that, in this situation, the Lambda function receives a private IP address and does not have direct access to the Internet. If you wish the function to access the Internet, you will need to treat it like a resource in a private subnet and use a NAT Gateway to provide Internet access.