package shared_machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new shared machine API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for shared machine API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SharedMachineAdd shared machine add API
*/
func (a *Client) SharedMachineAdd(params *SharedMachineAddParams, authInfo runtime.ClientAuthInfoWriter) (*SharedMachineAddOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSharedMachineAddParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SharedMachine.add",
		Method:             "POST",
		PathPattern:        "/remote.api/SharedMachine.add",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SharedMachineAddReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SharedMachineAddOK), nil

}

/*
SharedMachineKick shared machine kick API
*/
func (a *Client) SharedMachineKick(params *SharedMachineKickParams, authInfo runtime.ClientAuthInfoWriter) (*SharedMachineKickOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSharedMachineKickParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SharedMachine.kick",
		Method:             "POST",
		PathPattern:        "/remote.api/SharedMachine.kick",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SharedMachineKickReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SharedMachineKickOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
