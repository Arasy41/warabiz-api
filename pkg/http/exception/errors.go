package exception

type MapErrors map[string][]string

func NewMapErr() MapErrors {
	return make(MapErrors)
}

func (m *MapErrors) JoinErrors(errs ...MapErrors) {
	if len(errs) == 0 {
		return
	}
	if m.IsNill() {
		return
	}
	for _, err := range errs {
		if !err.IsNotEmpty() {
			continue
		}
		for key, value := range err {
			(*m)[key] = append((*m)[key], value...)
		}
	}
	return
}

func (m *MapErrors) AppendErrors(key string, errs ...string) {
	if m.IsNill() {
		return
	}
	(*m)[key] = append((*m)[key], errs...)
}

func (m MapErrors) IsNotEmpty() bool {
	if m == nil {
		return false
	}
	if len(m) == 0 {
		return false
	}
	return true
}

func (m MapErrors) IsNill() bool {
	if m == nil {
		return true
	}
	return false
}
