package sounds

type SoundFile struct {
	ID                 string
	FilePath           string
	Category           string
	NumberPlays        uint
	ExcludedFromRandom bool
}
