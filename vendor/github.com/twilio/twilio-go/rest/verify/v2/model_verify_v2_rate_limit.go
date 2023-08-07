/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Verify
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// VerifyV2RateLimit struct for VerifyV2RateLimit
type VerifyV2RateLimit struct {
	// A 34 character string that uniquely identifies this Rate Limit.
	Sid *string `json:"sid,omitempty"`
	// The SID of the [Service](https://www.twilio.com/docs/verify/api/service) the resource is associated with.
	ServiceSid *string `json:"service_sid,omitempty"`
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the Rate Limit resource.
	AccountSid *string `json:"account_sid,omitempty"`
	// Provides a unique and addressable name to be assigned to this Rate Limit, assigned by the developer, to be optionally used in addition to SID. **This value should not contain PII.**
	UniqueName *string `json:"unique_name,omitempty"`
	// Description of this Rate Limit
	Description *string `json:"description,omitempty"`
	// The date and time in GMT when the resource was created specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date and time in GMT when the resource was last updated specified in [RFC 2822](https://www.ietf.org/rfc/rfc2822.txt) format.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
	// The URLs of related resources.
	Links *map[string]interface{} `json:"links,omitempty"`
}