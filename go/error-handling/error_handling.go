package erratum

// Use takes a ResourceOpener and applies a given input string to it
func Use(o ResourceOpener, input string) (err error) {
	var r Resource
	r, err = o()
	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return err
		}
		r, err = o()
	}

	defer r.Close()
	defer func() {
		if rec := recover(); rec != nil {
			if f, ok := rec.(FrobError); ok {
				r.Defrob(f.defrobTag)
				err = f.inner
			}
			err = rec.(error)
		}
	}()
	r.Frob(input)

	return nil
}
