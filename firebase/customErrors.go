package firebase

type NoUidFoundErr struct {
}

func (e *NoUidFoundErr) Error() string {
	return "no uid found"
}

type UnexpectedUseFirestoreErr struct {
}

func (e *UnexpectedUseFirestoreErr) Error() string {
	return "Initialize Firestore if you want to use it"
}

type UnexpectedUseAuthErr struct {
}

func (e *UnexpectedUseAuthErr) Error() string {
	return "Initialize Auth if you want to use it"
}

type UnexpectedUseMessagingErr struct {
}

func (e *UnexpectedUseMessagingErr) Error() string {
	return "Initialize Messaging if you want to use it"
}
