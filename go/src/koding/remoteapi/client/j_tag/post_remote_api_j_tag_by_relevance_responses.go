package j_tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJTagByRelevanceReader is a Reader for the PostRemoteAPIJTagByRelevance structure.
type PostRemoteAPIJTagByRelevanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJTagByRelevanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJTagByRelevanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJTagByRelevanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJTagByRelevanceOK creates a PostRemoteAPIJTagByRelevanceOK with default headers values
func NewPostRemoteAPIJTagByRelevanceOK() *PostRemoteAPIJTagByRelevanceOK {
	return &PostRemoteAPIJTagByRelevanceOK{}
}

/*PostRemoteAPIJTagByRelevanceOK handles this case with default header values.

Request processed successfully
*/
type PostRemoteAPIJTagByRelevanceOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJTagByRelevanceOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTag.byRelevance][%d] postRemoteApiJTagByRelevanceOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJTagByRelevanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJTagByRelevanceUnauthorized creates a PostRemoteAPIJTagByRelevanceUnauthorized with default headers values
func NewPostRemoteAPIJTagByRelevanceUnauthorized() *PostRemoteAPIJTagByRelevanceUnauthorized {
	return &PostRemoteAPIJTagByRelevanceUnauthorized{}
}

/*PostRemoteAPIJTagByRelevanceUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJTagByRelevanceUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJTagByRelevanceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTag.byRelevance][%d] postRemoteApiJTagByRelevanceUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJTagByRelevanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
