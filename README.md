# Go Mock

## URL
https://github.com/golang/mock

## gomock
go get github.com/golang/mock/gomock

## mockgen install
go install github.com/golang/mock/mockgen@v1.6.0

## mockgen cmd
mockgen -source=app/domain/repository/user.go -destination app/domain/repository/mock/user.go -package mock