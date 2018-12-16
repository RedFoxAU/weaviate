// Code generated by go-swagger; DO NOT EDIT.

package things

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

// NewWeaviateThingsDeleteParams creates a new WeaviateThingsDeleteParams object
// with the default values initialized.
func NewWeaviateThingsDeleteParams() *WeaviateThingsDeleteParams {
	var ()
	return &WeaviateThingsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateThingsDeleteParamsWithTimeout creates a new WeaviateThingsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateThingsDeleteParamsWithTimeout(timeout time.Duration) *WeaviateThingsDeleteParams {
	var ()
	return &WeaviateThingsDeleteParams{

		timeout: timeout,
	}
}

// NewWeaviateThingsDeleteParamsWithContext creates a new WeaviateThingsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateThingsDeleteParamsWithContext(ctx context.Context) *WeaviateThingsDeleteParams {
	var ()
	return &WeaviateThingsDeleteParams{

		Context: ctx,
	}
}

// NewWeaviateThingsDeleteParamsWithHTTPClient creates a new WeaviateThingsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateThingsDeleteParamsWithHTTPClient(client *http.Client) *WeaviateThingsDeleteParams {
	var ()
	return &WeaviateThingsDeleteParams{
		HTTPClient: client,
	}
}

/*WeaviateThingsDeleteParams contains all the parameters to send to the API endpoint
for the weaviate things delete operation typically these are written to a http.Request
*/
type WeaviateThingsDeleteParams struct {

	/*ThingID
	  Unique ID of the Thing.

	*/
	ThingID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate things delete params
func (o *WeaviateThingsDeleteParams) WithTimeout(timeout time.Duration) *WeaviateThingsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate things delete params
func (o *WeaviateThingsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate things delete params
func (o *WeaviateThingsDeleteParams) WithContext(ctx context.Context) *WeaviateThingsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate things delete params
func (o *WeaviateThingsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate things delete params
func (o *WeaviateThingsDeleteParams) WithHTTPClient(client *http.Client) *WeaviateThingsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate things delete params
func (o *WeaviateThingsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithThingID adds the thingID to the weaviate things delete params
func (o *WeaviateThingsDeleteParams) WithThingID(thingID strfmt.UUID) *WeaviateThingsDeleteParams {
	o.SetThingID(thingID)
	return o
}

// SetThingID adds the thingId to the weaviate things delete params
func (o *WeaviateThingsDeleteParams) SetThingID(thingID strfmt.UUID) {
	o.ThingID = thingID
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateThingsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param thingId
	if err := r.SetPathParam("thingId", o.ThingID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}