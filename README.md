# A HTTP Interface for [pdf2htmlEx](https://github.com/coolwanglu/pdf2htmlEX)

This a basic HTTP service that converts PDF files to HTML using the library
[pdf2htmlEx](https://github.com/coolwanglu/pdf2htmlEX).

The HTTP service is implemented in Go. The client must provide a PDF file in a
POST request and the service will return the tranformed file in HTML.

**Example using Curl for consuming the service**
``` sh
    curl -F "pdf=@/home/chris/document.pdf" http://localhost:8080/ > output.html
```

Currently this is basic implementation that does not allow any kind of configuration.
The listening port is hardcoded to be **`8080`** and there is no way to send
custom calling params to pdf2htmlEX.

A GO intallation is required for building the HTTP service. Execute `build.sh`
to build the final executable file.

Feel free to fork and enhance the implementation of the server.
