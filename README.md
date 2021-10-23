# Garrus

A **Discord bot** written in Golang that serves as a **soundboard** for users in
a voice channel. This bot is able to be deployed to multiple servers through a
single Bot user. This project was started from scratch with many ideas borrowed 
from [NootBot](http://github.com/AlexSafatli/NootBot) and is built to have
as minimal a resource footprint as possible while serving no superfluous
features.

## Implemented Features

- **Sound file** maintenance
    - Local server disk storage of sound files
    - A command to list these files and see them listed under
      **categories** (`.list`)
    - A command to search these files by keyword (can leverage a trie data
      structure)
- **Welcome** message and **entrance** sound playing when users enter a voice
  channel
    - This entrance should be able to be set by users (`.entrance`)
    - Users who have not set an entrance will not receive a welcome message

## Planned Features

- **Sound file** maintenance
    - Possible future use of S3 to store clips
    - Way to directly upload sound files to the Discord server and have the bot
      consume it/download it
- **String** responses read from files
    - These should be read only at launch and be selected at random at different
      instances including welcome messages
    - Each response type should be located in a different `.txt` file; for
      example, `welcomes.txt` and `snide_comments.txt`
- An optional **REST API** that can be enabled on bot launch
    - Can enable creation of consumer applications and plugins
    - Will serve sound file lists and category lists
