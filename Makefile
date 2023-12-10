GOCMD=go
GOBUILD=${GOCMD} build -mod=mod
GOCLEAN=${GOCMD} clean

build: fuxibff

.PHONY: \
    fuixbff fuxi

clean:
	${GOCLEAN}

fuxibff:
	${GOBUILD} -o /Users/zongxingquan/goland/run/fuxibff github.com/zxq97/design/fuxi/bff/cmd

fuxi:
	${GOBUILD} -o /Users/zongxingquan/goland/run/fuxi github.com/zxq97/design/fuxi/service/cmd