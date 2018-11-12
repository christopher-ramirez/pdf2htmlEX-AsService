# /usr/bin/sh

# Builds the HTTP executable. A GO installation is required
echo "Building the HTTP Service..."
go build -a -tags netgo -o frontend *.go

echo "Service executable built successfull"
