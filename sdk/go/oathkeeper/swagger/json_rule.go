/*
 * ORY Oathkeeper
 *
 * ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.am
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type JsonRule struct {
	Description string `json:"description,omitempty"`

	Id string `json:"id,omitempty"`

	MatchesMethods []string `json:"matchesMethods,omitempty"`

	MatchesUrl string `json:"matchesUrl,omitempty"`

	Mode string `json:"mode,omitempty"`

	RequiredAction string `json:"requiredAction,omitempty"`

	RequiredResource string `json:"requiredResource,omitempty"`

	RequiredScopes []string `json:"requiredScopes,omitempty"`

	Upstream Upstream `json:"upstream,omitempty"`
}