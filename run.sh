docker run -w $GOPATH/src/github.com/franela/cli-kit -e GOPATH=$GOPATH -v /home/jonas:/home/jonas golang:1.8 go run main.go $@
