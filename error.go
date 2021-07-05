package gtils

// if error isn't nil,panic error
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
