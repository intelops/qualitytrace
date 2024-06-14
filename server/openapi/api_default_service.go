/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"encoding/json"
	"net/http"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// HealthzGet - Health Check
func (s *DefaultApiService) HealthzGet(ctx context.Context) (ImplResponse, error) {
	healthStatus := HealthzGet200Response{
		Status: "ok",
	}

	responseBody, err := json.Marshal(healthStatus)
	if err != nil {
		errorStatus := HealthzGet500Response{
			Status: "error",
		}
		errorBody, _ := json.Marshal(errorStatus)
		return Response(http.StatusInternalServerError, errorBody), err
	}

	return Response(http.StatusOK, responseBody), nil
}