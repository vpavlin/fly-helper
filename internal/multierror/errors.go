package multierror

import "fmt"

type Multierror []error

func (m Multierror) ToError() error {
	var joined string
	if len(m) == 0 {
		return nil
	}

	for _, e := range m {
		joined = fmt.Sprintf("%s; %s", joined, e)
	}

	return fmt.Errorf(joined)
}
