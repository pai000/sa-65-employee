import { ProvincesInterface } from "./IProvince";
import { GendersInterface } from "./IGender";
import { Job_positionsInterface } from "./IJob_Position";
export interface EmployeeInterface {
    ID?:        		number,
	Name?:  			string,
	Email?:     		string,
	Personal_ID?:  		string, 
	Password?:			string,
	ProvinceID?:    	number,
	Province?:    		ProvincesInterface,  
	GenderID?:        	number,    
	Gender?:        	GendersInterface,       
	Job_PositionID?:	number, 
	Job_Position?: 		Job_positionsInterface, 
}