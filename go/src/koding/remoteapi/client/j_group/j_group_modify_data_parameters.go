package j_group

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

// NewJGroupModifyDataParams creates a new JGroupModifyDataParams object
// with the default values initialized.
func NewJGroupModifyDataParams() *JGroupModifyDataParams {
	var ()
	return &JGroupModifyDataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJGroupModifyDataParamsWithTimeout creates a new JGroupModifyDataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJGroupModifyDataParamsWithTimeout(timeout time.Duration) *JGroupModifyDataParams {
	var ()
	return &JGroupModifyDataParams{

		timeout: timeout,
	}
}

// NewJGroupModifyDataParamsWithContext creates a new JGroupModifyDataParams object
// with the default values initialized, and the ability to set a context for a request
func NewJGroupModifyDataParamsWithContext(ctx context.Context) *JGroupModifyDataParams {
	var ()
	return &JGroupModifyDataParams{

		Context: ctx,
	}
}

/*JGroupModifyDataParams contains all the parameters to send to the API endpoint
for the j group modify data operation typically these are written to a http.Request
*/
type JGroupModifyDataParams struct {

	/*Body
	  body of the request

	*/
	Body JGroupModifyDataBody
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j group modify data params
func (o *JGroupModifyDataParams) WithTimeout(timeout time.Duration) *JGroupModifyDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j group modify data params
func (o *JGroupModifyDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j group modify data params
func (o *JGroupModifyDataParams) WithContext(ctx context.Context) *JGroupModifyDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j group modify data params
func (o *JGroupModifyDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j group modify data params
func (o *JGroupModifyDataParams) WithBody(body JGroupModifyDataBody) *JGroupModifyDataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j group modify data params
func (o *JGroupModifyDataParams) SetBody(body JGroupModifyDataBody) {
	o.Body = body
}

// WithID adds the id to the j group modify data params
func (o *JGroupModifyDataParams) WithID(id string) *JGroupModifyDataParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j group modify data params
func (o *JGroupModifyDataParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JGroupModifyDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
