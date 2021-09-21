package sound

type Entrance struct {
	UserID              string
	SoundID             string
	PersonalizedMessage string
}

var entrances []Entrance

func GetEntranceForUser(userID string) *Entrance {
	for _, e := range entrances {
		if e.UserID == userID {
			return &e
		}
	}
	return nil
}
