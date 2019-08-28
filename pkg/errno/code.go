package errno

import "fmt"

var (
	_codes = map[int]struct{}{}
)

func New(e int) int {
	if e <= 0 {
		panic("business ecode must greater than zero")
	}
	return add(e)
}

func add(e int) int {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return e
}

var (
	// Common errors
	OK                  	= &Errno{Code: New(1), 	 Message: "OK"}
	ErrInternalServer	 	= &Errno{Code: New(1001), Message: "error Internal server"}
	ErrBind             	= &Errno{Code: New(1002), Message: "error bind struct"}

	ErrValidation 			= &Errno{Code: New(2001), Message: "error Validation failed"}
	ErrDatabase   			= &Errno{Code: New(2002), Message: "error Database"}
	ErrToken      			= &Errno{Code: New(2003), Message: "error signing the JSON web token."}

	// user errors
	ErrPassword           	= &Errno{Code: New(2101), Message: "error user password not right"}
	ErrUserNotFound      	= &Errno{Code: New(2102), Message: "error can not found this user"}
	ErrUserHadRegister      = &Errno{Code: New(2103), Message: "error the user had register"}
)
