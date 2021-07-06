package reqdata

// UserReq contains user information
// https://pkg.go.dev/github.com/go-playground/validator/v10#hdr-Less_Than
type UserReq struct {
	FirstName      string `validate:"required"`
	LastName       string `validate:"required"`
	Age            uint8  `validate:"gte=0,lte=130"`
	FavouriteColor string `validate:"required_if=Age 10"` // alias for 'hexcolor|rgb|rgba|hsl|hsla'
}
