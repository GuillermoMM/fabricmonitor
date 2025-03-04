rm handler.zip
rm handler
# GOOS=linux go build -o handler -a -ldflags "-w -s -extldflags \"-static\"" -installsuffix cgo 
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o handler -a -ldflags "-w -s -extldflags \"-static\"" -installsuffix cgo 
zip handler.zip ./handler
