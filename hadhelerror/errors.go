package hadhelerror

import "errors"

// ErrXLSXFail is returned when the XLSX file could not be generated
var ErrXLSXFail = errors.New("Could not generate the XLSX file")

// ErrFailedRequest is returned when Go could not make the request (client or network problem)
var ErrFailedRequest = errors.New("Failed request")

// ErrAPIEmpty is returned when the API responded with a 204 status code
var ErrAPIEmpty = errors.New("204 (empty) response")

// ErrAPITimeout is returned when the API timed out during the call
var ErrAPITimeout = errors.New("Timeout during the call")

// ErrAPIUnauthorized is returned when the API returns 401
var ErrAPIUnauthorized = errors.New("401 (unauthorized) response")

// ErrEncoding is returned when the encoding of the request body failed
var ErrEncoding = errors.New("Could not encode request body")

// ErrDecoding is returned when the decoding of the response body failed
var ErrDecoding = errors.New("Could not decode response body")
