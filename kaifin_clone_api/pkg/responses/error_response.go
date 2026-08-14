package error_responses

import "fmt"

type ErrorResponse struct {
	MessageID string `json:"message_id"`
	Err       error  `json:"error"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("Message: %s Error %v", e.MessageID, e.Err)
}

func (e *ErrorResponse) ErrorString() string {
	return fmt.Sprintf("Message: %s Error %v", e.MessageID, e.Err)
}

func (e *ErrorResponse) NewErrorResponse(MessageId string, err error) *ErrorResponse {
	return &ErrorResponse{
		MessageID: MessageId,
		Err:       err,
	}
}

// server,service អ្នកផ្ដល់សេវាកម្ម & client អ្នកស្មើរសេវាកម្ម
// http server ជាអ្នករងចាំផ្ដល់សេវាកម្មនៅពេលមានសំណើរពី client
