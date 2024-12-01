package erratum

// https://exercism.org/tracks/go/exercises/error-handling

func isTransit(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(TransientError)
	return ok
}

// func safeGuard(frobFn func(input string)) func() {
// 	return func() {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				if fErr, ok := r.(FrobError); ok {

// 				}
// 			}
// 		}()
// 		frobFn("")
// 	}

// }

func safeGuardRes(res Resource, input string) (retErr error) {
	// handle panic, still return explicit err to the caller for normal control
	// deferred functions may read and assign to the returning function’s named return values.
	defer func() {
		if r := recover(); r != nil {
			if fErr, ok := r.(FrobError); ok {
				res.Defrob(fErr.defrobTag)
				retErr = fErr
			} else if oerr, ok := r.(error); ok {
				retErr = oerr
			}
			// return
		}
	}()

	res.Frob(input)
	// if no panic from Frob, otherwise control goes to defer recover
	return nil
}

func Use(opener ResourceOpener, input string) error {
	res, err := opener()
	for isTransit(err) {
		res, err = opener()
	}
	// err is not transient
	if err != nil {
		return err
	}

	// res is opened
	defer res.Close()

	// safeguard resource Frob ops
	// if FrobError need to call Defrob with defrobTag and close
	// otherwise return error and close res
	// res.Frob(input)
	err = safeGuardRes(res, input)
	return err
}
