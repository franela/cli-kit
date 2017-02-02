docker run -w $GOPATH/src/github.com/franela/cli-kit/plugins/lala -e GOPATH=$GOPATH -v /home/jonas:/home/jonas golang:1.8 go build -buildmode=plugin -o /home/jonas/.foo/plugins/lala.so
docker run -w $GOPATH/src/github.com/franela/cli-kit/plugins/lolo -e GOPATH=$GOPATH -v /home/jonas:/home/jonas golang:1.8 go build -buildmode=plugin -o /home/jonas/.foo/plugins/lolo.so
docker run -w $GOPATH/src/github.com/franela/cli-kit/plugins/docker -e GOPATH=$GOPATH -v /home/jonas:/home/jonas golang:1.8 go build -buildmode=plugin -o /home/jonas/.foo/plugins/docker.so
sudo chown -R jonas:jonas ~/.foo
