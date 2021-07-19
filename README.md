# sumocli
A CLI application that lets you manage/automate your Sumo Logic tenancy. 

Sumocli is currently in development so there could be bugs/incomplete functionality.
GA will be v1.0.0 which I am expecting to be ready for release towards the end of 2021.
## Installation

### Recommended
Download the binary for your platform from the Releases page. 

### Build Yourself
You can build the sumocli application for your platform by performing the following steps:

Clone the sumocli repo

`git clone https://github.com/wizedkyle/sumocli`

The repo is using Go modules so you can run go build:

`go build ./cmd/sumocli`

## Authentication

Sumocli uses two authentication methods;
- Environment variables
- Credentials file

When you run a command in sumocli it will first check to see if a credentials file exists, if it can't find one then it will fall back to environment variables. This is to ensure that sumocli can run in CI/CD pipelines. 
The sections below explain the requirements for each authentication type.

### Environment Variables

Environment variable authentication is useful when running sumocli in a CI/CD pipeline. The following environment variables need to be set to allow for proper authentication.

```
SUMO_ACCESS_ID: abcefghi

SUMO_ACCESS_KEY: AbCeFG123

SUMO_ENDPOINT: https://api.<regioncode>.sumologic.com/api
```

For a full list of Sumo Logic regions to API endpoints see this page: 
https://help.sumologic.com/APIs/General-API-Information/Sumo-Logic-Endpoints-and-Firewall-Security

### Credentials File

The credentials file stores the same information as the environment variables however, it can be generated interactively using `sumocli login`. 
The Sumo Logic Access ID and Access Key are encrypted (using AES-256) before being written to the credentials file to reduce the risk of the credentials being 
used outside of Sumocli.

Encryption of the Sumo Logic Access ID and Access Key was added in v0.9.0 of Sumocli, if you are running
an earlier version of Sumocli you will need to regenerate your credentials file by running `sumocli login` if you want 
to leverage encryption at rest.
If you need to know which Access ID sumocli is configured to use you can run `sumocli login --showAccessId` and
the plaintext access ID will be displayed.

The credential file is stored in the following locations depending on your operating system.

```
Windows: C:\Users\<username>\.sumocli\credentials\creds.json

Macos: /User/<username>/.sumocli/credentials/creds.json

Linux: /Usr/<username>/.sumocli/credentials/creds.json
```

The contents of the credential file is as follows:

```
{
  "version": "v1",
  "accessid": "abcefghi",
  "accesskey": "AbCeFG123",
  "region": "<regioncode>",
  "endpoint": "https://api.<regioncode>.sumologic.com/api"
}
```

## Documentation

Documentation for each command can be found by using `sumocli <command> --help`

You can find a list of API endpoints and sources that Sumocli supports in the [supported capabilities](COMPATIBILITY.md) documentation.

## Contributing

Clone or fork the repo, make your changes and create a pull request. 
I will then review it and all things looking good it gets merged!

If there is something in the code that you don't understand please feel free to email at kyle@thepublicclouds.com.

