package handler

import (
	"data"
	
	"encoding/json"

)
const(
	REGISTER = "register"
	LOGIN = "login"
)

func HandleApiRequest(req *data.ApiRequest) []byte {
	var response *data.ApiResponse = nil
	switch *req.Api{
	
	case LOGIN : 
		response = &data.ApiResponse{Status:200,ResponseData:"login successful",Err:""}
	case REGISTER:
		response = &data.ApiResponse{Status:200,ResponseData:"successfully register",Err:""}
	default :
		response = &data.ApiResponse{Status:200,ResponseData:nil,Err:"Invalid Api Request"}

	}
	resp,err :=	json.Marshal(response)
	if(nil != err){
		return nil
		//response = &data.ApiResponse{Status:500,ResponseData:nil,Err:err.Error()}
	}
	return resp
}

//		fmt.Sprintf("your api is %s",*apiRequest.Api)
