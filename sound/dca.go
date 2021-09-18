package sound

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jonas747/dca"
	"io"
	"os"
	"time"
)

func encodeSoundFileToDCA(pathToSoundFile string) (io.Reader, error) {
	return dca.EncodeFile(pathToSoundFile, dca.StdEncodeOptions)
}

func saveSoundFileToDCA(source, target string) error {
	enc, err := encodeSoundFileToDCA(source)
	if err != nil {
		return err
	}
	o, err := os.Create(target)
	if err != nil {
		return err
	}
	_, err = io.Copy(o, enc)
	return err
}

func playDCA(path string, vc *discordgo.VoiceConnection) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	dec := dca.NewDecoder(f)
	for {
		frame, err := dec.OpusFrame()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		// Send frame to discord
		select {
		case vc.OpusSend <- frame:
		case <-time.After(time.Second):
			return nil
		}
	}
	return nil
}
