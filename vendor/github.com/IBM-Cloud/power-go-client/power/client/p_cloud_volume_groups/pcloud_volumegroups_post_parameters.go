// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volume_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudVolumegroupsPostParams creates a new PcloudVolumegroupsPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudVolumegroupsPostParams() *PcloudVolumegroupsPostParams {
	return &PcloudVolumegroupsPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudVolumegroupsPostParamsWithTimeout creates a new PcloudVolumegroupsPostParams object
// with the ability to set a timeout on a request.
func NewPcloudVolumegroupsPostParamsWithTimeout(timeout time.Duration) *PcloudVolumegroupsPostParams {
	return &PcloudVolumegroupsPostParams{
		timeout: timeout,
	}
}

// NewPcloudVolumegroupsPostParamsWithContext creates a new PcloudVolumegroupsPostParams object
// with the ability to set a context for a request.
func NewPcloudVolumegroupsPostParamsWithContext(ctx context.Context) *PcloudVolumegroupsPostParams {
	return &PcloudVolumegroupsPostParams{
		Context: ctx,
	}
}

// NewPcloudVolumegroupsPostParamsWithHTTPClient creates a new PcloudVolumegroupsPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudVolumegroupsPostParamsWithHTTPClient(client *http.Client) *PcloudVolumegroupsPostParams {
	return &PcloudVolumegroupsPostParams{
		HTTPClient: client,
	}
}

/*
PcloudVolumegroupsPostParams contains all the parameters to send to the API endpoint

	for the pcloud volumegroups post operation.

	Typically these are written to a http.Request.
*/
type PcloudVolumegroupsPostParams struct {

	/* Body.

	   Parameters for the creation of a new volume group
	*/
	Body *models.VolumeGroupCreate

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud volumegroups post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudVolumegroupsPostParams) WithDefaults() *PcloudVolumegroupsPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud volumegroups post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudVolumegroupsPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud volumegroups post params
func (o *PcloudVolumegroupsPostParams) WithTimeout(timeout time.Duration) *PcloudVolumegroupsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud volumegroups post params
func (o *PcloudVolumegroupsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud volumegroups post params
func (o *PcloudVolumegroupsPostParams) WithContext(ctx context.Context) *PcloudVolumegroupsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud volumegroups post params
func (o *PcloudVolumegroupsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud volumegroups post params
func (o *PcloudVolumegroupsPostParams) WithHTTPClient(client *http.Client) *PcloudVolumegroupsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud volumegroups post params
func (o *PcloudVolumegroupsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud volumegroups post params
func (o *PcloudVolumegroupsPostParams) WithBody(body *models.VolumeGroupCreate) *PcloudVolumegroupsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud volumegroups post params
func (o *PcloudVolumegroupsPostParams) SetBody(body *models.VolumeGroupCreate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud volumegroups post params
func (o *PcloudVolumegroupsPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudVolumegroupsPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud volumegroups post params
func (o *PcloudVolumegroupsPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudVolumegroupsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
