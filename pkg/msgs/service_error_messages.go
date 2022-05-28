package msgs

var (
	ErrInvalidValue            = "Invalid value"
	ErrConvFailedProtoToStruct = "failed to convert struct to proto"
	ErrConvFailedStructToProto = "failed to convert structure to proto"
	ErrUnmarshalError          = "failed to unmarshal"
	ErrDBCreate                = "failed to create object in database"
	ErrDBUpdate                = "failed to update object in database"
	ErrDBGetAll                = "failed to retrieve all object in database"
	ErrDBDelete                = "failed to delete object in database"
	ErrDBGet                   = "failed to get object in database"
	ErrUpdStage                = "error while updating stage"
	WrongStage                 = "wrong stage"
)
