// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudTasksGetReader is a Reader for the PcloudTasksGet structure.
type PcloudTasksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudTasksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudTasksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudTasksGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudTasksGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudTasksGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudTasksGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudTasksGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/tasks/{task_id}] pcloud.tasks.get", response, response.Code())
	}
}

// NewPcloudTasksGetOK creates a PcloudTasksGetOK with default headers values
func NewPcloudTasksGetOK() *PcloudTasksGetOK {
	return &PcloudTasksGetOK{}
}

/*
PcloudTasksGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudTasksGetOK struct {
	Payload *models.Task
}

// IsSuccess returns true when this pcloud tasks get o k response has a 2xx status code
func (o *PcloudTasksGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud tasks get o k response has a 3xx status code
func (o *PcloudTasksGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tasks get o k response has a 4xx status code
func (o *PcloudTasksGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud tasks get o k response has a 5xx status code
func (o *PcloudTasksGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tasks get o k response a status code equal to that given
func (o *PcloudTasksGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud tasks get o k response
func (o *PcloudTasksGetOK) Code() int {
	return 200
}

func (o *PcloudTasksGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetOK  %+v", 200, o.Payload)
}

func (o *PcloudTasksGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetOK  %+v", 200, o.Payload)
}

func (o *PcloudTasksGetOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *PcloudTasksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTasksGetBadRequest creates a PcloudTasksGetBadRequest with default headers values
func NewPcloudTasksGetBadRequest() *PcloudTasksGetBadRequest {
	return &PcloudTasksGetBadRequest{}
}

/*
PcloudTasksGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudTasksGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tasks get bad request response has a 2xx status code
func (o *PcloudTasksGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tasks get bad request response has a 3xx status code
func (o *PcloudTasksGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tasks get bad request response has a 4xx status code
func (o *PcloudTasksGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tasks get bad request response has a 5xx status code
func (o *PcloudTasksGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tasks get bad request response a status code equal to that given
func (o *PcloudTasksGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud tasks get bad request response
func (o *PcloudTasksGetBadRequest) Code() int {
	return 400
}

func (o *PcloudTasksGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTasksGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudTasksGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTasksGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTasksGetUnauthorized creates a PcloudTasksGetUnauthorized with default headers values
func NewPcloudTasksGetUnauthorized() *PcloudTasksGetUnauthorized {
	return &PcloudTasksGetUnauthorized{}
}

/*
PcloudTasksGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudTasksGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tasks get unauthorized response has a 2xx status code
func (o *PcloudTasksGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tasks get unauthorized response has a 3xx status code
func (o *PcloudTasksGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tasks get unauthorized response has a 4xx status code
func (o *PcloudTasksGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tasks get unauthorized response has a 5xx status code
func (o *PcloudTasksGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tasks get unauthorized response a status code equal to that given
func (o *PcloudTasksGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud tasks get unauthorized response
func (o *PcloudTasksGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudTasksGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudTasksGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudTasksGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTasksGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTasksGetForbidden creates a PcloudTasksGetForbidden with default headers values
func NewPcloudTasksGetForbidden() *PcloudTasksGetForbidden {
	return &PcloudTasksGetForbidden{}
}

/*
PcloudTasksGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudTasksGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tasks get forbidden response has a 2xx status code
func (o *PcloudTasksGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tasks get forbidden response has a 3xx status code
func (o *PcloudTasksGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tasks get forbidden response has a 4xx status code
func (o *PcloudTasksGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tasks get forbidden response has a 5xx status code
func (o *PcloudTasksGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tasks get forbidden response a status code equal to that given
func (o *PcloudTasksGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud tasks get forbidden response
func (o *PcloudTasksGetForbidden) Code() int {
	return 403
}

func (o *PcloudTasksGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudTasksGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudTasksGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTasksGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTasksGetNotFound creates a PcloudTasksGetNotFound with default headers values
func NewPcloudTasksGetNotFound() *PcloudTasksGetNotFound {
	return &PcloudTasksGetNotFound{}
}

/*
PcloudTasksGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudTasksGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tasks get not found response has a 2xx status code
func (o *PcloudTasksGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tasks get not found response has a 3xx status code
func (o *PcloudTasksGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tasks get not found response has a 4xx status code
func (o *PcloudTasksGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud tasks get not found response has a 5xx status code
func (o *PcloudTasksGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud tasks get not found response a status code equal to that given
func (o *PcloudTasksGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud tasks get not found response
func (o *PcloudTasksGetNotFound) Code() int {
	return 404
}

func (o *PcloudTasksGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudTasksGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudTasksGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTasksGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudTasksGetInternalServerError creates a PcloudTasksGetInternalServerError with default headers values
func NewPcloudTasksGetInternalServerError() *PcloudTasksGetInternalServerError {
	return &PcloudTasksGetInternalServerError{}
}

/*
PcloudTasksGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudTasksGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud tasks get internal server error response has a 2xx status code
func (o *PcloudTasksGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud tasks get internal server error response has a 3xx status code
func (o *PcloudTasksGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud tasks get internal server error response has a 4xx status code
func (o *PcloudTasksGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud tasks get internal server error response has a 5xx status code
func (o *PcloudTasksGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud tasks get internal server error response a status code equal to that given
func (o *PcloudTasksGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud tasks get internal server error response
func (o *PcloudTasksGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudTasksGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTasksGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/tasks/{task_id}][%d] pcloudTasksGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudTasksGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudTasksGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
