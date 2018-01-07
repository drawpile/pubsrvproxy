package queries

type UserInfo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Ip string `json:"ip,omitempty"`
	Auth bool `json:"auth"`
	Op bool `json:"op"`
	Muted bool `json:"muted"`
	Mod bool `json:"mod"`
	Session string `json:"session"`
}

func QueryUserList(server QueryOpts) ([]UserInfo, error) {
	var users []UserInfo
	err := getJson(server.ServerAddr + "users/", &users, server.Cache)

	return users, err
}

