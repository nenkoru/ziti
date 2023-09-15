// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package router

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/fabric/controller/rest_model"
)

// ListRouterTerminatorsReader is a Reader for the ListRouterTerminators structure.
type ListRouterTerminatorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRouterTerminatorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRouterTerminatorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListRouterTerminatorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListRouterTerminatorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListRouterTerminatorsOK creates a ListRouterTerminatorsOK with default headers values
func NewListRouterTerminatorsOK() *ListRouterTerminatorsOK {
	return &ListRouterTerminatorsOK{}
}

/* ListRouterTerminatorsOK describes a response with status code 200, with default header values.

A list of terminators
*/
type ListRouterTerminatorsOK struct {
	Payload *rest_model.ListTerminatorsEnvelope
}

func (o *ListRouterTerminatorsOK) Error() string {
	return fmt.Sprintf("[GET /routers/{id}/terminators][%d] listRouterTerminatorsOK  %+v", 200, o.Payload)
}
func (o *ListRouterTerminatorsOK) GetPayload() *rest_model.ListTerminatorsEnvelope {
	return o.Payload
}

func (o *ListRouterTerminatorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.ListTerminatorsEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRouterTerminatorsBadRequest creates a ListRouterTerminatorsBadRequest with default headers values
func NewListRouterTerminatorsBadRequest() *ListRouterTerminatorsBadRequest {
	return &ListRouterTerminatorsBadRequest{}
}

/* ListRouterTerminatorsBadRequest describes a response with status code 400, with default header values.

The supplied request contains invalid fields or could not be parsed (json and non-json bodies). The error's code, message, and cause fields can be inspected for further information
*/
type ListRouterTerminatorsBadRequest struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListRouterTerminatorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /routers/{id}/terminators][%d] listRouterTerminatorsBadRequest  %+v", 400, o.Payload)
}
func (o *ListRouterTerminatorsBadRequest) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListRouterTerminatorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRouterTerminatorsUnauthorized creates a ListRouterTerminatorsUnauthorized with default headers values
func NewListRouterTerminatorsUnauthorized() *ListRouterTerminatorsUnauthorized {
	return &ListRouterTerminatorsUnauthorized{}
}

/* ListRouterTerminatorsUnauthorized describes a response with status code 401, with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type ListRouterTerminatorsUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *ListRouterTerminatorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /routers/{id}/terminators][%d] listRouterTerminatorsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListRouterTerminatorsUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *ListRouterTerminatorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
