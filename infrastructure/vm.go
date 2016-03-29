package infrastructure

type VM struct {
	Cid      string
	JobName  string
	JobState string
	Index    int
	Ips      []string
}

type VMs []VM
