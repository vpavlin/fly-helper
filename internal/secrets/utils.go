package secrets

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/superfly/graphql"
)

func ToSecretName(prefix string, n string) string {
	r := regexp.MustCompile("[^a-zA-Z0-9_]+")
	return fmt.Sprintf("%s_%s", PREFIX, strings.ToUpper(r.ReplaceAllString(n, "_")))
}

func IsUnchangedErr(err error) bool {
	var gqlErr *graphql.GraphQLError

	if errors.As(err, &gqlErr) {
		return gqlErr.Extensions.Code == "UNCHANGED"
	}
	return false
}
