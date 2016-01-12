package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/examples/todo-list/models"
)

// NewAddOneParams creates a new AddOneParams object
// with the default values initialized.
func NewAddOneParams() *AddOneParams {
	var ()
	return &AddOneParams{}
}

/*AddOneParams contains all the parameters to send to the API endpoint
for the add one operation typically these are written to a http.Request
*/
type AddOneParams struct {

	/*Body*/
	Body *models.Item
}

// WithBody adds the body to the add one params
func (o *AddOneParams) WithBody(body *models.Item) *AddOneParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *AddOneParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.Item)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
