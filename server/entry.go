package main

import "time"

type ServerConfig struct {
	Server struct {
		Host   string `yaml:"host"`
		Port   string `yaml:"port"`
		Secret string `yaml:"secret"`
	}
	ReleaseServer map[string][]string `yaml:"release_server"`
}

type JsonData struct {
	Action string `json:"action"`
	Release struct {
		ID int `json:"id"`
		TagName string `json:"tag_name"`
		TargetCommitish string `json:"target_commitish"`
		Name string `json:"name"`
		Body string `json:"body"`
		Draft bool `json:"draft"`
		Prerelease bool `json:"prerelease"`
		Author struct {
			ID int `json:"id"`
			Login string `json:"login"`
			FullName string `json:"full_name"`
			Email string `json:"email"`
			AvatarURL string `json:"avatar_url"`
			Username string `json:"username"`
		} `json:"author"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"release"`
	Repository struct {
		ID int `json:"id"`
		Owner struct {
			ID int `json:"id"`
			Login string `json:"login"`
			FullName string `json:"full_name"`
			Email string `json:"email"`
			AvatarURL string `json:"avatar_url"`
			Username string `json:"username"`
		} `json:"owner"`
		Name string `json:"name"`
		FullName string `json:"full_name"`
		Description string `json:"description"`
		Private bool `json:"private"`
		Fork bool `json:"fork"`
		Parent interface{} `json:"parent"`
		Empty bool `json:"empty"`
		Mirror bool `json:"mirror"`
		Size int `json:"size"`
		HTMLURL string `json:"html_url"`
		SSHURL string `json:"ssh_url"`
		CloneURL string `json:"clone_url"`
		Website string `json:"website"`
		StarsCount int `json:"stars_count"`
		ForksCount int `json:"forks_count"`
		WatchersCount int `json:"watchers_count"`
		OpenIssuesCount int `json:"open_issues_count"`
		DefaultBranch string `json:"default_branch"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"repository"`
	Sender struct {
		ID int `json:"id"`
		Login string `json:"login"`
		FullName string `json:"full_name"`
		Email string `json:"email"`
		AvatarURL string `json:"avatar_url"`
		Username string `json:"username"`
	} `json:"sender"`
}

