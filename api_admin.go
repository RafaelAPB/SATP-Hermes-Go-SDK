/*
SATP Gateway Client (Business Logic Orchestrator)

SATP is a protocol operating between two gateways that conducts the transfer of a digital asset from one gateway to another. The protocol establishes a secure channel between the endpoints and implements a 2-phase commit to ensure the properties of transfer atomicity, consistency, isolation and durability.  This API defines the gateway client facing API (business logic orchestrator, or BLO), which is named API-Type 1 in the SATP-Core specification.  **Additional Resources**: - [Proposed SATP Charter](https://datatracker.ietf.org/doc/charter-ietf-satp/) - [SATP Core draft](https://datatracker.ietf.org/doc/draft-ietf-satp-core) - [SATP Crash Recovery draft](https://datatracker.ietf.org/doc/draft-belchior-satp-gateway-recovery/) - [SATP Architecture draft](https://datatracker.ietf.org/doc/draft-ietf-satp-architecture/) - [SATP Use-Cases draft](https://datatracker.ietf.org/doc/draft-ramakrishna-sat-use-cases/) - [SATP Data sharing draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-data-sharing) - [SATP View Addresses draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-views-addresses)

API version: v0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)


type AdminApi interface {

	/*
	CallContinue Continue a paused transaction session

	Attempts to continue a previously paused transaction intent, resuming its execution.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCallContinueRequest
	*/
	CallContinue(ctx context.Context) ApiCallContinueRequest

	// CallContinueExecute executes the request
	//  @return Continue200Response
	CallContinueExecute(r ApiCallContinueRequest) (*Continue200Response, *http.Response, error)

	/*
	GetAudit Audit transactions

	Audits transactions based on provided filters such as start and end dates. Optionally includes proofs generated from each gateway transaction.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAuditRequest
	*/
	GetAudit(ctx context.Context) ApiGetAuditRequest

	// GetAuditExecute executes the request
	//  @return GetAudit200Response
	GetAuditExecute(r ApiGetAuditRequest) (*GetAudit200Response, *http.Response, error)

	/*
	GetStatus Get SATP current session data

	Retrieve the status of a SATP session

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetStatusRequest
	*/
	GetStatus(ctx context.Context) ApiGetStatusRequest

	// GetStatusExecute executes the request
	//  @return Transact200ResponseStatusResponse
	GetStatusExecute(r ApiGetStatusRequest) (*Transact200ResponseStatusResponse, *http.Response, error)

	/*
	Pause Pause a transaction session

	Attempts to pause a previously submitted transaction intent, temporarily halting its execution.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPauseRequest
	*/
	Pause(ctx context.Context) ApiPauseRequest

	// PauseExecute executes the request
	//  @return Pause200Response
	PauseExecute(r ApiPauseRequest) (*Pause200Response, *http.Response, error)
}

// AdminApiService AdminApi service
type AdminApiService service

type ApiCallContinueRequest struct {
	ctx context.Context
	ApiService AdminApi
	continueRequest *ContinueRequest
}

func (r ApiCallContinueRequest) ContinueRequest(continueRequest ContinueRequest) ApiCallContinueRequest {
	r.continueRequest = &continueRequest
	return r
}

func (r ApiCallContinueRequest) Execute() (*Continue200Response, *http.Response, error) {
	return r.ApiService.CallContinueExecute(r)
}

/*
CallContinue Continue a paused transaction session

Attempts to continue a previously paused transaction intent, resuming its execution.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCallContinueRequest
*/
func (a *AdminApiService) CallContinue(ctx context.Context) ApiCallContinueRequest {
	return ApiCallContinueRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Continue200Response
func (a *AdminApiService) CallContinueExecute(r ApiCallContinueRequest) (*Continue200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Continue200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdminApiService.CallContinue")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/@hyperledger/cactus-plugin-satp-hermes/continue"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.continueRequest == nil {
		return localVarReturnValue, nil, reportError("continueRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.continueRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v TransactDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetAuditRequest struct {
	ctx context.Context
	ApiService AdminApi
	auditStartDate *time.Time
	auditEndDate *time.Time
	includeProofs *bool
}

// The start date for the audit period.
func (r ApiGetAuditRequest) AuditStartDate(auditStartDate time.Time) ApiGetAuditRequest {
	r.auditStartDate = &auditStartDate
	return r
}

// The end date for the audit period.
func (r ApiGetAuditRequest) AuditEndDate(auditEndDate time.Time) ApiGetAuditRequest {
	r.auditEndDate = &auditEndDate
	return r
}

// Include proofs generated from each gateway transaction.
func (r ApiGetAuditRequest) IncludeProofs(includeProofs bool) ApiGetAuditRequest {
	r.includeProofs = &includeProofs
	return r
}

func (r ApiGetAuditRequest) Execute() (*GetAudit200Response, *http.Response, error) {
	return r.ApiService.GetAuditExecute(r)
}

/*
GetAudit Audit transactions

Audits transactions based on provided filters such as start and end dates. Optionally includes proofs generated from each gateway transaction.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAuditRequest
*/
func (a *AdminApiService) GetAudit(ctx context.Context) ApiGetAuditRequest {
	return ApiGetAuditRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetAudit200Response
func (a *AdminApiService) GetAuditExecute(r ApiGetAuditRequest) (*GetAudit200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAudit200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdminApiService.GetAudit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/@hyperledger/cactus-plugin-satp-hermes/audit"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.auditStartDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "auditStartDate", r.auditStartDate, "")
	}
	if r.auditEndDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "auditEndDate", r.auditEndDate, "")
	}
	if r.includeProofs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeProofs", r.includeProofs, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetStatusRequest struct {
	ctx context.Context
	ApiService AdminApi
	sessionID *string
}

// Unique identifier for the session.
func (r ApiGetStatusRequest) SessionID(sessionID string) ApiGetStatusRequest {
	r.sessionID = &sessionID
	return r
}

func (r ApiGetStatusRequest) Execute() (*Transact200ResponseStatusResponse, *http.Response, error) {
	return r.ApiService.GetStatusExecute(r)
}

/*
GetStatus Get SATP current session data

Retrieve the status of a SATP session

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetStatusRequest
*/
func (a *AdminApiService) GetStatus(ctx context.Context) ApiGetStatusRequest {
	return ApiGetStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Transact200ResponseStatusResponse
func (a *AdminApiService) GetStatusExecute(r ApiGetStatusRequest) (*Transact200ResponseStatusResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Transact200ResponseStatusResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdminApiService.GetStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/@hyperledger/cactus-plugin-satp-hermes/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sessionID == nil {
		return localVarReturnValue, nil, reportError("sessionID is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "SessionID", r.sessionID, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPauseRequest struct {
	ctx context.Context
	ApiService AdminApi
	pauseRequest *PauseRequest
}

func (r ApiPauseRequest) PauseRequest(pauseRequest PauseRequest) ApiPauseRequest {
	r.pauseRequest = &pauseRequest
	return r
}

func (r ApiPauseRequest) Execute() (*Pause200Response, *http.Response, error) {
	return r.ApiService.PauseExecute(r)
}

/*
Pause Pause a transaction session

Attempts to pause a previously submitted transaction intent, temporarily halting its execution.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPauseRequest
*/
func (a *AdminApiService) Pause(ctx context.Context) ApiPauseRequest {
	return ApiPauseRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Pause200Response
func (a *AdminApiService) PauseExecute(r ApiPauseRequest) (*Pause200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Pause200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdminApiService.Pause")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/@hyperledger/cactus-plugin-satp-hermes/pause"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pauseRequest == nil {
		return localVarReturnValue, nil, reportError("pauseRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.pauseRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v TransactDefaultResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
