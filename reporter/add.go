package reporter

func (r *Reporter) Add(report *Report) error {
	err := r.read()
	if err != nil {
		return err
	}

	defer r.write()

	for index, i := range r.Items {
		if i.Date == report.Date.Format("2006-01-02") {
			r.Items[index] = report.object()
			return nil
		}
	}

	r.Items = append(r.Items, report.object())

	return nil
}
