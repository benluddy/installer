// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// ListClustersReader is a Reader for the ListClusters structure.
type ListClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListClustersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClustersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewListClustersMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListClustersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewListClustersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListClustersOK creates a ListClustersOK with default headers values
func NewListClustersOK() *ListClustersOK {
	return &ListClustersOK{}
}

/* ListClustersOK describes a response with status code 200, with default header values.

Success.
*/
type ListClustersOK struct {
	Payload models.ClusterList
}

func (o *ListClustersOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters][%d] listClustersOK  %+v", 200, o.Payload)
}
func (o *ListClustersOK) GetPayload() models.ClusterList {
	return o.Payload
}

func (o *ListClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersUnauthorized creates a ListClustersUnauthorized with default headers values
func NewListClustersUnauthorized() *ListClustersUnauthorized {
	return &ListClustersUnauthorized{}
}

/* ListClustersUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type ListClustersUnauthorized struct {
	Payload *models.InfraError
}

func (o *ListClustersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/clusters][%d] listClustersUnauthorized  %+v", 401, o.Payload)
}
func (o *ListClustersUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *ListClustersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersForbidden creates a ListClustersForbidden with default headers values
func NewListClustersForbidden() *ListClustersForbidden {
	return &ListClustersForbidden{}
}

/* ListClustersForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type ListClustersForbidden struct {
	Payload *models.InfraError
}

func (o *ListClustersForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/clusters][%d] listClustersForbidden  %+v", 403, o.Payload)
}
func (o *ListClustersForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *ListClustersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersMethodNotAllowed creates a ListClustersMethodNotAllowed with default headers values
func NewListClustersMethodNotAllowed() *ListClustersMethodNotAllowed {
	return &ListClustersMethodNotAllowed{}
}

/* ListClustersMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type ListClustersMethodNotAllowed struct {
	Payload *models.Error
}

func (o *ListClustersMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /v1/clusters][%d] listClustersMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *ListClustersMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListClustersMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersInternalServerError creates a ListClustersInternalServerError with default headers values
func NewListClustersInternalServerError() *ListClustersInternalServerError {
	return &ListClustersInternalServerError{}
}

/* ListClustersInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type ListClustersInternalServerError struct {
	Payload *models.Error
}

func (o *ListClustersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/clusters][%d] listClustersInternalServerError  %+v", 500, o.Payload)
}
func (o *ListClustersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListClustersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersServiceUnavailable creates a ListClustersServiceUnavailable with default headers values
func NewListClustersServiceUnavailable() *ListClustersServiceUnavailable {
	return &ListClustersServiceUnavailable{}
}

/* ListClustersServiceUnavailable describes a response with status code 503, with default header values.

Unavailable.
*/
type ListClustersServiceUnavailable struct {
	Payload *models.Error
}

func (o *ListClustersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /v1/clusters][%d] listClustersServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ListClustersServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListClustersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
