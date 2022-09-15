package constants

const (
	ValidationError       = "error in validating request"
	RequestError          = "error in request params"
	UnprocessableEntity   = "Unprocessable Entity"
	StartDateTime         = "0001-01-01 00:00:00 +0000 UTC"
	UserNotFound          = "user record not found"
	NoHoldingsFound       = "no holdings for user & scheme code found"
	NoTransactionFound    = "no transaction for user found"
	WalletAmountIsLow     = "wallet amount for this transaction is low"
	UnitsAmountIsLow      = "holding units amount for this transaction is low"
	SchemeCodeIsInvalid   = "scheme code in the request is invalid"
	NoSuchSchemCodeExists = "no such scheme code exists"
)
