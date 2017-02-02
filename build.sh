docker run -i -w $GOPATH/src/github.com/franela/cli-kit/plugins/lala -e GOPATH=$GOPATH -v $HOME:$HOME golang:1.8 go build -buildmode=plugin -o $HOME/.foo/plugins/lala.so
docker run -i -w $GOPATH/src/github.com/franela/cli-kit/plugins/lolo -e GOPATH=$GOPATH -v $HOME:$HOME golang:1.8 go build -buildmode=plugin -o $HOME/.foo/plugins/lolo.so
#docker run -i -w $GOPATH/src/github.com/franela/cli-kit/plugins/docker -e GOPATH=$GOPATH -v $HOME:$HOME golang:1.8 go build -buildmode=plugin -o $HOME/.foo/plugins/docker.so
sudo chown -R $USER:$USER ~/.foo
