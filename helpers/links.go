package helpers

import (
	"strconv"
	"strings"

	"github.com/aptible/aptible-api-go/aptibleapi"
)

// ExtractIDFromHref parses a numeric ID from the last segment of a HAL link href.
func ExtractIDFromHref(href string) int32 {
	if href == "" {
		return 0
	}
	segments := strings.Split(href, "/")
	if len(segments) == 0 {
		return 0
	}
	val, _ := strconv.ParseInt(segments[len(segments)-1], 10, 32)
	return int32(val)
}

// GetStackOrganizationID returns the organization ID for a stack, extracted from its HAL links.
// Returns empty string if the organization link is not present.
func GetStackOrganizationID(stack *aptibleapi.Stack) string {
	if stack.Links == nil || stack.Links.Organization == nil || stack.Links.Organization.Href == nil {
		return ""
	}
	href := *stack.Links.Organization.Href
	segments := strings.Split(href, "/")
	if len(segments) == 0 {
		return ""
	}
	return segments[len(segments)-1]
}
