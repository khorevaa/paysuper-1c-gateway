package internal

import "github.com/paysuper/paysuper-proto/go/billingpb"

var (
	errorRequestParamsIncorrect = newManagementApiResponseError("ma000023", "incorrect request parameters")
	errorUnknown                = newManagementApiResponseError("ma000001", "unknown error. try request later")
)

func newManagementApiResponseError(code, msg string, details ...string) *billingpb.ResponseErrorMessage {
	var det string
	if len(details) > 0 && details[0] != "" {
		det = details[0]
	} else {
		det = ""
	}
	return &billingpb.ResponseErrorMessage{Code: code, Message: msg, Details: det}
}
