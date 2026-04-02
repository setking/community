package docs

import (
	"myApp/forms"
)

// swagger:route POST /create comment createCommentRequest
// Create a comment in memory.
// responses:
//   200: createCommentResponse
//   default: errResponse

// swagger:route GET /list comment getCommentRequest
// Get a comment from memory.
// responses:
//   200: getCommentResponse
//   default: errResponse

// swagger:parameters createCommentRequest
type commentParamsWrapper struct {
	// This text will appear as description of your request body.
	// in:body
	Body forms.CommentForm
}

// This text will appear as description of your request url path.
// swagger:parameters createCommentResponse
type createCommentResponseWrapper struct {
	// This text will appear as description of your request body.
	// in:body

}

// This text will appear as description of your request url path.
// swagger:parameters getCommentRequest
type getCommentParamsWrapper struct {
	// in:path
	P int `json:"p"`
	N int `json:"n"`
}

// This text will appear as description of your request url path.
// swagger:parameters getCommentResponse
type getCommentResponseWrapper struct {
	// in:path
	Data  []forms.CommentForm
	Total int `json:"n"`
}

// This text will appear as description of your error response body.
// swagger:response errResponse
type errResponseWrapper struct {
	// Error code.
	Code int `json:"code"`

	// Error message.
	Message string `json:"message"`
}
