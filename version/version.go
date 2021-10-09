package version

type version struct {
	Name      string
	Version   string
	GitCommit string
	Developer string
}

var Version = version{
	Name:      "Garrus",
	Version:   "0.1.2",
	GitCommit: "HEAD",
	Developer: "Asaph",
}
