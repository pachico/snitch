# snitch

Snitch is a small (~2.5 MB) Go application designed to run within a Docker environment to provide context information. It captures and displays data about the Docker pod it operates in, including environment variables, system information, DNS resolutions, and HTTP request details.

Its purpose is to aid debugging in development, testing, and production environments.

- [snitch](#snitch)
  - [Usage](#usage)
  - [Environment Variables](#environment-variables)
  - [Featured Endpoints](#featured-endpoints)
    - [System Information](#system-information)
    - [Environment Variables](#environment-variables-1)
    - [DNS Resolution](#dns-resolution)
    - [HTTP Requests and Responses](#http-requests-and-responses)
    - [Filesystem Information](#filesystem-information)
  - [License](#license)

## Usage

Deploy Snitch as a Docker container. A ready-to-use image is available at <https://hub.docker.com/repository/docker/pachico/snitch/tags>.

To run it locally, use the following command:

```bash
docker run --rm -p 8080:1323 --env SNITCH_PORT=1323 pachico/snitch:latest
```

This makes the application available at `http://localhost:8080`.

Deploying it in a Kubernetes cluster allows for retrieving information about the pod it's running in.

## Environment Variables

The following environment variables can be set to customize the application:

- `SNITCH_PORT`: The port the application listens on. Default is `1323`.

## Featured Endpoints

Access these endpoints via HTTP GET requests as detailed in the
[internal/infrastructure/routes/routes.go](internal/infrastructure/routes/routes.go) file.

All responses are formatted according to the JSEND specification and are pretty-printed for readability.

### System Information

Displays system, OS, kernel, hardware, and network details.

```bash

$ curl {prot}://{host}:{port}/ -X GET

{
  "status": "success",
  "data": {
    "sysinfo": {
      "version": "1.0.2",
      "timestamp": "REDACTED"
    },
    "node": {
      "hostname": "REDACTED",
      "machineid": "REDACTED",
      "timezone": "Europe/Madrid"
    },
    "os": {
      "name": "Ubuntu 22.04.4 LTS",
      "vendor": "ubuntu",
      "version": "22.04",
      "release": "22.04.4",
      "architecture": "amd64"
    },
    "kernel": {
      "release": "REDACTED",
      "version": "REDACTED",
      "architecture": "x86_64"
    },
    "product": {
      "name": "REDACTED",
      "vendor": "Dell Inc.",
      "uuid": "REDACTED",
      "sku": "REDACTED"
    },
    "board": {
      "name": "REDACTED",
      "vendor": "Dell Inc.",
      "version": "A00"
    },
    "chassis": {
      "type": 10,
      "vendor": "Dell Inc."
    },
    "bios": {
      "vendor": "Dell Inc.",
      "version": "1.6.2",
      "date": "REDACTED"
    },
    "cpu": {
      "vendor": "GenuineIntel",
      "model": "REDACTED",
      "cache": 24576,
      "cpus": 1,
      "cores": 8,
      "threads": 16
    },
    "memory": {},
    "storage": [
      {
        "name": "REDACTED",
        "model": "REDACTED",
        "serial": "REDACTED",
        "size": 1024
      }
    ],
    "network": [
      {
        "name": "REDACTED",
        "driver": "REDACTED",
        "macaddress": "REDACTED",
        "port": "tp/mii",
        "speed": 1000
      },
      {
        "name": "REDACTED",
        "driver": "iwlwifi",
        "macaddress": "REDACTED"
      }
    ]
  }
}

```

### Environment Variables

Lists all environment variables.

> **WARNING**: This endpoint may expose sensitive information. Use it with caution.

```bash

$ curl {prot}://{host}:{port}/envvars -X GET

{
  "status": "success",
  "data": {
      "FOO": "bar",
      "BAZ": "qux"
    }
}

```

### DNS Resolution

Resolves the IP address of the specified domain, useful for debugging DNS issues.

```bash

$ curl {prot}://{host}:{port}/dnsresolution/photos.google.com -X GET

{
  "status": "success",
  "data": {
    "hostname": "photos.google.com",
    "ips": [
      {
        "ip": "142.250.186.142",
        "is_private": false,
        "is_global_unicast": true,
        "is_loopback": false,
        "is_unspecified": false
      }
    ]
  }
}

```

In case of not being able to resolve the domain, the response will be:

```bash

$ curl {prot}://{host}:{port}/dnsresolution/idontexist -X GET

{
  "status": "fail",
  "data": "error resolving idontexist: lookup idontexist on 127.0.0.53:53: server misbehaving"
}

```

### HTTP Requests and Responses

Performs an HTTP(S) request from the container and returns the response.

```bash

$ curl {prot}://{host}:{port}/httprequest/https://example.com/ -X GET

{
  "status": "success",
  "data": {
    "request": {
      "url": "https://example.com/"
    },
    "response": {
      "status_code": 200,
      "size": -1,
      "headers": {
        "Accept-Ranges": [
          "bytes"
        ],
        "Age": [
          "499083"
        ],
        "Cache-Control": [
          "max-age=604800"
        ],
        "Content-Type": [
          "text/html; charset=UTF-8"
        ],
        "Date": [
          "Wed, 08 May 2024 13:19:54 GMT"
        ],
        "Etag": [
          "\"3147526947+gzip\""
        ],
        "Expires": [
          "Wed, 15 May 2024 13:19:54 GMT"
        ],
        "Last-Modified": [
          "Thu, 17 Oct 2019 07:18:26 GMT"
        ],
        "Server": [
          "ECAcc (nyd/D139)"
        ],
        "Vary": [
          "Accept-Encoding"
        ],
        "X-Cache": [
          "HIT"
        ]
      },
      "body": "<!doctype html>\n<html>\n<head>\n    <title>Example Domain</title>\n\n    <meta charset=\"utf-8\" />\n    <meta http-equiv=\"Content-type\" content=\"text/html; charset=utf-8\" />\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" />\n    <style type=\"text/css\">\n    body {\n        background-color: #f0f0f2;\n        margin: 0;\n        padding: 0;\n        font-family: -apple-system, system-ui, BlinkMacSystemFont, \"Segoe UI\", \"Open Sans\", \"Helvetica Neue\", Helvetica, Arial, sans-serif;\n        \n    }\n    div {\n        width: 600px;\n        margin: 5em auto;\n        padding: 2em;\n        background-color: #fdfdff;\n        border-radius: 0.5em;\n        box-shadow: 2px 3px 7px 2px rgba(0,0,0,0.02);\n    }\n    a:link, a:visited {\n        color: #38488f;\n        text-decoration: none;\n    }\n    @media (max-width: 700px) {\n        div {\n            margin: 0 auto;\n            width: auto;\n        }\n    }\n    </style>    \n</head>\n\n<body>\n<div>\n    <h1>Example Domain</h1>\n    <p>This domain is for use in illustrative examples in documents. You may use this\n    domain in literature without prior coordination or asking for permission.</p>\n    <p><a href=\"https://www.iana.org/domains/example\">More information...</a></p>\n</div>\n</body>\n</html>\n"
    }
  }
}
```

In case of error, the response will be:

```bash

$ curl {prot}://{host}:{port}/httprequest/https://idontexist/ -X GET

{
  "status": "fail",
  "data": "error resolving or requesting https://idontexist/: Get \"https://idontexist/\": dial tcp: lookup idontexist on 127.0.0.53:53: server misbehaving"
}

```

### Filesystem Information

Provides details about files and directories up to 2 levels deep.

```bash

$ curl {prot}://{host}:{port}/fs -X GET

{
  "status": "success",
  "data": [
    {
      "name": "/",
      "size": 32,
      "is_dir": true,
      "owner": "root",
      "permission": "-rwxr-xr-x"
    },
    {
      "name": "/usr/bin",
      "size": 7,
      "is_dir": false,
      "owner": "root",
      "permission": "-rwxrwxrwx"
    }
    // ...
  ]
}

```

## License

This project is licensed under the Apache License. See the [LICENSE](LICENSE) file for details.
