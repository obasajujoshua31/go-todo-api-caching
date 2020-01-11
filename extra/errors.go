package extra

type ServerError struct {
	message string
	status  int
}

var UnableToConnectToDBError = ServerError{
	message: "Unable to Connect to Database",
	status:  500,
}

var UnableToMarshalResponse = ServerError{
	message: "Unable to Marshal Response",
	status:  500,
}

var UnableToReadBodyError = ServerError{
	message: "Unable to read body error",
	status:  500,
}

var UnableToUnmarshalError = ServerError{
	message: "Unable to Unmarshal error",
	status:  500,
}

var UnableToConnectToRedis = ServerError{
	message: "Unable to Connect To Redis",
	status:  500,
}

var UnableToSetCacheError = ServerError{
	message: "Unable to Set Cache Error",
	status:  500,
}
