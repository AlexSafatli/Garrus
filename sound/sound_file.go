package sound

type File struct {
	ID                 string
	FilePath           string
	Categories         []string
	NumberPlays        uint
	ExcludedFromRandom bool
}
