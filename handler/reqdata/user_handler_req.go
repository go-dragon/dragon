package reqdata

// UserReq contains user information
// https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Less_Than
type UserReq struct {
	UserId int64 `validate:"required,number"`
}
