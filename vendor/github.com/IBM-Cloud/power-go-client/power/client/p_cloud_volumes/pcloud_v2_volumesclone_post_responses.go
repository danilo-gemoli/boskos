// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudV2VolumesclonePostReader is a Reader for the PcloudV2VolumesclonePost structure.
type PcloudV2VolumesclonePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2VolumesclonePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudV2VolumesclonePostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudV2VolumesclonePostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudV2VolumesclonePostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudV2VolumesclonePostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudV2VolumesclonePostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone] pcloud.v2.volumesclone.post", response, response.Code())
	}
}

// NewPcloudV2VolumesclonePostAccepted creates a PcloudV2VolumesclonePostAccepted with default headers values
func NewPcloudV2VolumesclonePostAccepted() *PcloudV2VolumesclonePostAccepted {
	return &PcloudV2VolumesclonePostAccepted{}
}

/*
PcloudV2VolumesclonePostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudV2VolumesclonePostAccepted struct {
	Payload *models.VolumesClone
}

// IsSuccess returns true when this pcloud v2 volumesclone post accepted response has a 2xx status code
func (o *PcloudV2VolumesclonePostAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud v2 volumesclone post accepted response has a 3xx status code
func (o *PcloudV2VolumesclonePostAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone post accepted response has a 4xx status code
func (o *PcloudV2VolumesclonePostAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 volumesclone post accepted response has a 5xx status code
func (o *PcloudV2VolumesclonePostAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone post accepted response a status code equal to that given
func (o *PcloudV2VolumesclonePostAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud v2 volumesclone post accepted response
func (o *PcloudV2VolumesclonePostAccepted) Code() int {
	return 202
}

func (o *PcloudV2VolumesclonePostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumesclonePostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudV2VolumesclonePostAccepted) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumesclonePostAccepted  %+v", 202, o.Payload)
}

func (o *PcloudV2VolumesclonePostAccepted) GetPayload() *models.VolumesClone {
	return o.Payload
}

func (o *PcloudV2VolumesclonePostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumesClone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesclonePostBadRequest creates a PcloudV2VolumesclonePostBadRequest with default headers values
func NewPcloudV2VolumesclonePostBadRequest() *PcloudV2VolumesclonePostBadRequest {
	return &PcloudV2VolumesclonePostBadRequest{}
}

/*
PcloudV2VolumesclonePostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudV2VolumesclonePostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone post bad request response has a 2xx status code
func (o *PcloudV2VolumesclonePostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone post bad request response has a 3xx status code
func (o *PcloudV2VolumesclonePostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone post bad request response has a 4xx status code
func (o *PcloudV2VolumesclonePostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone post bad request response has a 5xx status code
func (o *PcloudV2VolumesclonePostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone post bad request response a status code equal to that given
func (o *PcloudV2VolumesclonePostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud v2 volumesclone post bad request response
func (o *PcloudV2VolumesclonePostBadRequest) Code() int {
	return 400
}

func (o *PcloudV2VolumesclonePostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumesclonePostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudV2VolumesclonePostBadRequest) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumesclonePostBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudV2VolumesclonePostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesclonePostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesclonePostUnauthorized creates a PcloudV2VolumesclonePostUnauthorized with default headers values
func NewPcloudV2VolumesclonePostUnauthorized() *PcloudV2VolumesclonePostUnauthorized {
	return &PcloudV2VolumesclonePostUnauthorized{}
}

/*
PcloudV2VolumesclonePostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudV2VolumesclonePostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone post unauthorized response has a 2xx status code
func (o *PcloudV2VolumesclonePostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone post unauthorized response has a 3xx status code
func (o *PcloudV2VolumesclonePostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone post unauthorized response has a 4xx status code
func (o *PcloudV2VolumesclonePostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone post unauthorized response has a 5xx status code
func (o *PcloudV2VolumesclonePostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone post unauthorized response a status code equal to that given
func (o *PcloudV2VolumesclonePostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud v2 volumesclone post unauthorized response
func (o *PcloudV2VolumesclonePostUnauthorized) Code() int {
	return 401
}

func (o *PcloudV2VolumesclonePostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumesclonePostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV2VolumesclonePostUnauthorized) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumesclonePostUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudV2VolumesclonePostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesclonePostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesclonePostForbidden creates a PcloudV2VolumesclonePostForbidden with default headers values
func NewPcloudV2VolumesclonePostForbidden() *PcloudV2VolumesclonePostForbidden {
	return &PcloudV2VolumesclonePostForbidden{}
}

/*
PcloudV2VolumesclonePostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudV2VolumesclonePostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone post forbidden response has a 2xx status code
func (o *PcloudV2VolumesclonePostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone post forbidden response has a 3xx status code
func (o *PcloudV2VolumesclonePostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone post forbidden response has a 4xx status code
func (o *PcloudV2VolumesclonePostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone post forbidden response has a 5xx status code
func (o *PcloudV2VolumesclonePostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone post forbidden response a status code equal to that given
func (o *PcloudV2VolumesclonePostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud v2 volumesclone post forbidden response
func (o *PcloudV2VolumesclonePostForbidden) Code() int {
	return 403
}

func (o *PcloudV2VolumesclonePostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumesclonePostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudV2VolumesclonePostForbidden) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumesclonePostForbidden  %+v", 403, o.Payload)
}

func (o *PcloudV2VolumesclonePostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesclonePostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesclonePostInternalServerError creates a PcloudV2VolumesclonePostInternalServerError with default headers values
func NewPcloudV2VolumesclonePostInternalServerError() *PcloudV2VolumesclonePostInternalServerError {
	return &PcloudV2VolumesclonePostInternalServerError{}
}

/*
PcloudV2VolumesclonePostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudV2VolumesclonePostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone post internal server error response has a 2xx status code
func (o *PcloudV2VolumesclonePostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone post internal server error response has a 3xx status code
func (o *PcloudV2VolumesclonePostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone post internal server error response has a 4xx status code
func (o *PcloudV2VolumesclonePostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 volumesclone post internal server error response has a 5xx status code
func (o *PcloudV2VolumesclonePostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud v2 volumesclone post internal server error response a status code equal to that given
func (o *PcloudV2VolumesclonePostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud v2 volumesclone post internal server error response
func (o *PcloudV2VolumesclonePostInternalServerError) Code() int {
	return 500
}

func (o *PcloudV2VolumesclonePostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumesclonePostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2VolumesclonePostInternalServerError) String() string {
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone][%d] pcloudV2VolumesclonePostInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudV2VolumesclonePostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesclonePostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
