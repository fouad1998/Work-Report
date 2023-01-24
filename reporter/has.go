package reporter

import (
	"time"
)

func (r *Reporter) Has(date time.Time) (bool, error) {
	if err := r.read(); err != nil {
		return false, err
	}

	for _, i := range r.Items {
		if i.Date == date.Format("2006-01-02") {
			return true, nil
		}
	}

	return false, nil
}
