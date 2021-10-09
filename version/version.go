package version

type version struct {
	Name      string
	Version   string
	GitCommit string
	Developer string
}

var (
	VersionStr   = "0.0.0"
	GitCommitStr = "HEAD"
)

var Version = version{
	Name:      "Garrus",
	Version:   VersionStr,
	GitCommit: GitCommitStr,
	Developer: "Asaph",
}
