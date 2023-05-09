// JSend is a package to make standardizing API responses easy
// It is based on the jsend specification that can be found at:
// https://github.com/omniti-labs/jsend
package jsend

type response struct {
	// Required. Either "success" for a successful request, "fail"
	// for a failed request, or "error" for a server error.
	Status string `json:"status"`
	// Required for "success" and "fail". Data is a wrapper for
	// any information the API returns.
	// If the reasons for failure correspond to POST values,
	// the response object's keys SHOULD correspond to those POST values.
	Data map[string]interface{} `json:"data,omitempty"`
	// Required for "error".  A meaningful, end-user-readable
	// message, explaining what went wrong.
	Message string `json:"message,omitempty"`
	// Optional for "error". A numerical error code.
	Code int `json:"code,omitempty"`
}
