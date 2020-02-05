// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/acongos92/andyauth/models"
)

// VerifyReader is a Reader for the Verify structure.
type VerifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VerifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVerifyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewVerifyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVerifyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVerifyNoContent creates a VerifyNoContent with default headers values
func NewVerifyNoContent() *VerifyNoContent {
	return &VerifyNoContent{}
}

/*VerifyNoContent handles this case with default header values.

valid token
*/
type VerifyNoContent struct {
}

func (o *VerifyNoContent) Error() string {
	return fmt.Sprintf("[POST /verify][%d] verifyNoContent ", 204)
}

func (o *VerifyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVerifyForbidden creates a VerifyForbidden with default headers values
func NewVerifyForbidden() *VerifyForbidden {
	return &VerifyForbidden{}
}

/*VerifyForbidden handles this case with default header values.

token was not valid
*/
type VerifyForbidden struct {
}

func (o *VerifyForbidden) Error() string {
	return fmt.Sprintf("[POST /verify][%d] verifyForbidden ", 403)
}

func (o *VerifyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVerifyInternalServerError creates a VerifyInternalServerError with default headers values
func NewVerifyInternalServerError() *VerifyInternalServerError {
	return &VerifyInternalServerError{}
}

/*VerifyInternalServerError handles this case with default header values.

internal server error
*/
type VerifyInternalServerError struct {
	Payload *models.Error
}

func (o *VerifyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /verify][%d] verifyInternalServerError  %+v", 500, o.Payload)
}

func (o *VerifyInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *VerifyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
