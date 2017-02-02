docker run -v $GOPATH:$GOPATH -w $GOPATH/src/github.com/franela/cli-kit -e GOPATH=$GOPATH -v $HOME:/root golang:1.8 go run main.go $@
