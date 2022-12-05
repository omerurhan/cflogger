# CFLOGGER 

## What is cflogger?

---
**cflogger** CloudFormation streaming logger for stack events.
It logs single stack events in given time. Not gives all events history.
Suitable for tracking stack build processes in ci/cd pipelines.
## Examples
> **Warning**
> You must set AWS region. Otherwise you get error.
---
```
# Follows a single stack events occurring after the specified time (UTC+0).
$ cflogger --stack-name test --since "2022-12-15 15:00"  --region eu-central-1

# Follows events from the time it was run until events finish.
$ cflogger --stack-name test --region eu-central-1

# Select stack id. 
$ cflogger --stack-name "arn:aws:cloudformation:eu-central-1:890724537041:stack/test/1448c0d9-7173-114d-af33-0ac9xyzd9fb36" --region eu-central-1

# Get stack-id from aws cloudformation command output and follow stream.
$ aws cloudformation update-stack --stack-name test --template-body file:///Users/omer/sample-template.yml  | cflogger --region eu-central-1 --stack-name -

# Define timeout in minutes for cflogger.
$ cflogger --stack-name test --timeout 20 --region eu-central-1
```

## Configuring Credentials
---

When using the cflogger you'll need your AWS credentials to authenticate with AWS services. cflogger supports multiple methods of supporting these credentials.

**Environment Credentials** - Set of environment variables that are useful when sub processes are created for specific roles.

Shared Credentials file (**~/.aws/credentials**) - This file stores your credentials based on a profile name and is useful for local development.

## Installation 
---
```
# Clone repository and go directory
$ git clone https://github.com/omerurhan/cflogger.git
$ cd cflogger

# Set OS, ARCH environment variables for build. 
# Example linux, amd64.
export OS=linux
export ARCH=amd64


# Build the binary according to your operating system.
$ make all

# Copy convenient location and run. 
$ cp target/cflogger /usr/local/bin 
```