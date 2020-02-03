package firebase

//NoUIDFoundErr - No uid found after token verification
type NoUIDFoundErr struct {
}

func (e *NoUIDFoundErr) Error() string {
	return "no uid found"
}

//UnexpectedUseFirestoreErr - Attempt at using firebase without initialization
//pass the param firebase = true on Connect
type UnexpectedUseFirestoreErr struct {
}

func (e *UnexpectedUseFirestoreErr) Error() string {
	return "Initialize Firestore if you want to use it"
}

//UnexpectedUseAuthErr - Attempt at using auth without initialization
//pass the param auth = true on Connect
type UnexpectedUseAuthErr struct {
}

func (e *UnexpectedUseAuthErr) Error() string {
	return "Initialize Auth if you want to use it"
}

//UnexpectedUseMessagingErr - Attempt at using messaging without initialization
//pass the param messaging = true on Connect
type UnexpectedUseMessagingErr struct {
}

func (e *UnexpectedUseMessagingErr) Error() string {
	return "Initialize Messaging if you want to use it"
}
