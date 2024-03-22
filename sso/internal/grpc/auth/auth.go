package auth

import (
	ssov1 "github.com/17HIERARCH70/messageService/api-contracts/gen/go/sso"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}
