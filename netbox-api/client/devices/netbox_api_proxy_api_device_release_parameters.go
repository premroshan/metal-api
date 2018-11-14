// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewNetboxAPIProxyAPIDeviceReleaseParams creates a new NetboxAPIProxyAPIDeviceReleaseParams object
// with the default values initialized.
func NewNetboxAPIProxyAPIDeviceReleaseParams() *NetboxAPIProxyAPIDeviceReleaseParams {
	var ()
	return &NetboxAPIProxyAPIDeviceReleaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNetboxAPIProxyAPIDeviceReleaseParamsWithTimeout creates a new NetboxAPIProxyAPIDeviceReleaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNetboxAPIProxyAPIDeviceReleaseParamsWithTimeout(timeout time.Duration) *NetboxAPIProxyAPIDeviceReleaseParams {
	var ()
	return &NetboxAPIProxyAPIDeviceReleaseParams{

		timeout: timeout,
	}
}

// NewNetboxAPIProxyAPIDeviceReleaseParamsWithContext creates a new NetboxAPIProxyAPIDeviceReleaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewNetboxAPIProxyAPIDeviceReleaseParamsWithContext(ctx context.Context) *NetboxAPIProxyAPIDeviceReleaseParams {
	var ()
	return &NetboxAPIProxyAPIDeviceReleaseParams{

		Context: ctx,
	}
}

// NewNetboxAPIProxyAPIDeviceReleaseParamsWithHTTPClient creates a new NetboxAPIProxyAPIDeviceReleaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNetboxAPIProxyAPIDeviceReleaseParamsWithHTTPClient(client *http.Client) *NetboxAPIProxyAPIDeviceReleaseParams {
	var ()
	return &NetboxAPIProxyAPIDeviceReleaseParams{
		HTTPClient: client,
	}
}

/*NetboxAPIProxyAPIDeviceReleaseParams contains all the parameters to send to the API endpoint
for the netbox api proxy api device release operation typically these are written to a http.Request
*/
type NetboxAPIProxyAPIDeviceReleaseParams struct {

	/*UUID
	  The product serial of the device (unique identifier of this device)

	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the netbox api proxy api device release params
func (o *NetboxAPIProxyAPIDeviceReleaseParams) WithTimeout(timeout time.Duration) *NetboxAPIProxyAPIDeviceReleaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the netbox api proxy api device release params
func (o *NetboxAPIProxyAPIDeviceReleaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the netbox api proxy api device release params
func (o *NetboxAPIProxyAPIDeviceReleaseParams) WithContext(ctx context.Context) *NetboxAPIProxyAPIDeviceReleaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the netbox api proxy api device release params
func (o *NetboxAPIProxyAPIDeviceReleaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the netbox api proxy api device release params
func (o *NetboxAPIProxyAPIDeviceReleaseParams) WithHTTPClient(client *http.Client) *NetboxAPIProxyAPIDeviceReleaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the netbox api proxy api device release params
func (o *NetboxAPIProxyAPIDeviceReleaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUID adds the uuid to the netbox api proxy api device release params
func (o *NetboxAPIProxyAPIDeviceReleaseParams) WithUUID(uuid string) *NetboxAPIProxyAPIDeviceReleaseParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the netbox api proxy api device release params
func (o *NetboxAPIProxyAPIDeviceReleaseParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NetboxAPIProxyAPIDeviceReleaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}