package internal

import "github.com/paysuper/paysuper-billing-server/pkg/proto/grpc"

var (
	errorRequestParamsIncorrect = newManagementApiResponseError("ma000023", "incorrect request parameters")
	errorUnknown                = newManagementApiResponseError("ma000001", "unknown error. try request later")
)

func newManagementApiResponseError(code, msg string, details ...string) *grpc.ResponseErrorMessage {
	var det string
	if len(details) > 0 && details[0] != "" {
		det = details[0]
	} else {
		det = ""
	}
	return &grpc.ResponseErrorMessage{Code: code, Message: msg, Details: det}
}
