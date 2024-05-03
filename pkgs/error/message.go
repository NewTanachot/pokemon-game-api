package customerror

const (
	InvalidInput              string = "invalid_input"
	InvalidResponse           string = "invalid_response"
	UnableToParseJsonToStruct string = "unable_to_parse_json_to_struct"
	UnableToParseStructToJson string = "unable_to_parse_struct_to_json"
	UnableToCreate            string = "unable_to_create"
	PingDbFail                string = "ping_mongo_db_fail"
	WrongPassword             string = "wrong_password"
	UnableToCreateJWT         string = "unable_to_create_jwt"
)
