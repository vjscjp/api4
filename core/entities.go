package core

type App struct {
	Id          string  `json:"id"`
	ProjectName string  `json:"projectname"`
	ProjectId   string  `json:"projectid"`
	EnvName     string  `json:"envname"`
	EnvId       string  `json:"envid"`
	ServiceName string  `json:"servicename"`
	ServiceId   string  `json:"serviceid"`
	Tasks       []*Task `json:"tasks,omitempty"`
}
type ListApp struct {
	Id      string `json:"id"`
	Display string `json:"text"`
}
type Task struct {
	Host  string `json:"host"`
	Ports []int  `json:"ports"`
	Id    string `json:"id"`
	AppId string `json:"appid"`
}

type User struct {
	ID   string `json:"id"`
	Pass string `json:"pass"`
	Api  string `json:"api"`
}
