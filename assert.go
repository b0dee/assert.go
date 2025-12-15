package assert

func Assert(success bool, message string) {
	if !success { panic(message) }
}
