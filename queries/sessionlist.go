package queries

type sessionServerResponse struct {
	Alias string
	Closed bool
	Founder string
	HasPassword bool
	Id string
	MaxUserCount int
	Nsfm bool
	Persistent bool
	Protocol string
	Size int
	StartTime string
	Title string
	UserCount int
}

type SessionInfo struct {
	Host      string   `json:"host"`
	Port      int      `json:"port"`
	Id        string   `json:"id"`
	Alias     string   `json:"alias"`
	Protocol  string   `json:"protocol"`
	Title     string   `json:"title"`
	Users     int      `json:"users"`
	Password  bool     `json:"password"`
	Closed    bool     `json:"closed"`
	Nsfm      bool     `json:"nsfm"`
	Owner     string   `json:"owner"`
	Started   string   `json:"started"`
	Size      int      `json:"size"`
}

func QuerySessionInfo(server QueryOpts) ([]SessionInfo, error) {
	var sessions []sessionServerResponse
	err := getJson(server.ServerAddr + "sessions/", &sessions, server.Cache)

	if err != nil {
		return nil, err
	}

	response := make([]SessionInfo, 0, len(sessions))
	for _, s := range sessions {
		response = append(response, SessionInfo{
			Host: "localhost",
			Port: 27750,
			Id: s.Id,
			Alias: s.Alias,
			Protocol: s.Protocol,
			Title: s.Title,
			Users: s.UserCount,
			Password: s.HasPassword,
			Closed: s.Closed,
			Nsfm: s.Nsfm,
			Owner: s.Founder,
			Started: s.StartTime,
			Size: s.Size,
		})
	}

	return response, nil
}

