/**
@author: yeebing
@date: 2022/9/24
**/

package code

// Database
const (
	// ErrDatabase - 500: Database error
	ErrDatabase int = iota + 100001
)

// Type Transform
const (
	// ErrTypeAssertion -500: Type assertion error
	ErrTypeAssertion int = iota + 100101

	// ErrTypeTransform -500: Type transform error
	ErrTypeTransform
)

// Gin Server
const (
	// ErrBind -400: Error occurred while binding the request body to the struct.
	ErrBind int = iota + 100201

	// ErrGetFormFile -400: Can not get post form file
	ErrGetFormFile

	// ErrSaveUploadFile -500: Can not save the upload file
	ErrSaveUploadFile
)
