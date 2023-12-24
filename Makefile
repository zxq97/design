GOCMD=go
GOBUILD=${GOCMD} build -mod=mod
GOCLEAN=${GOCMD} clean

build: fuxibff

.PHONY: \
    fuixbff fuxi accountbff account archivebff archive likebff like

clean:
	${GOCLEAN}

fuxibff:
	${GOBUILD} -o /Users/zongxingquan/goland/run/fuxibff github.com/zxq97/design/fuxi/bff/cmd

fuxi:
	${GOBUILD} -o /Users/zongxingquan/goland/run/fuxi github.com/zxq97/design/fuxi/service/cmd

accountbff:
	${GOBUILD} -o /Users/zongxingquan/goland/run/accountbff github.com/zxq97/design/weitter/account/bff/cmd

account:
	${GOBUILD} -o /Users/zongxingquan/goland/run/account github.com/zxq97/design/weitter/account/service/cmd

archivebff:
	${GOBUILD} -o /Users/zongxingquan/goland/run/archivebff github.com/zxq97/design/weitter/archive/bff/cmd

archive:
	${GOBUILD} -o /Users/zongxingquan/goland/run/archive github.com/zxq97/design/weitter/archive/service/cmd

likebff:
	${GOBUILD} -o /Users/zongxingquan/goland/run/likebff github.com/zxq97/design/weitter/like/bff/cmd

like:
	${GOBUILD} -o /Users/zongxingquan/goland/run/like github.com/zxq97/design/weitter/like/service/cmd
