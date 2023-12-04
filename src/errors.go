package client

// error types and error handling logic specific to your SDK
type SDKError struct {
    Message string
}

func (e *SDKError) Error() string {
    return e.Message
}

// need to add the rest of the error handling functions that relates to the client side
