package infrastructure

type Director interface {
	GetDeployments() BoshDeployments
}

type BoshDirector struct {
	url      string
	username string
	password string
}

func NewDirector(url, username, password string) Director {
	return &BoshDirector{
		url:      url,
		username: username,
		password: password,
	}
}

func (d *BoshDirector) GetDeployments() BoshDeployments {
	return BoshDeployments{}
}
