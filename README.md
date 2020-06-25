# UDP Socket Server

UDP Socket Server with pluggable handlers

![Go](https://github.com/pavelkim/srcds_logserver/workflows/Go/badge.svg?branch=master&event=push) ![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/pavelkim/srcds_logserver?sort=semver) [![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/) ![GitHub go.mod Go version (subfolder of monorepo)](https://img.shields.io/github/go-mod/go-version/pavelkim/srcds_logserver?filename=go.mod)


# Handler plugin interface

Required symbols to be exported:
```python
PayloadHandlerDescription string
PayloadHandlerVersion     string
PayloadHandlerFunction    func([]byte)(bool, error)
```

# Usage

```bash
  -bind string
    	Address and port to listen (default "127.0.0.1:9001")
  -debug
    	Enable verbose output
  -handler string
    	Path to the payload handler shared library (default "handler.so")
  -version
    	Show version
```

# Examples

Run the server listening on address 127.0.0.1 and UDP port 28001
```bash
./srcds_logserver -handler srcds_handler.so -bind 127.0.0.1:28001
```
