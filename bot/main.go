package main

import (
  "github.com/bwmarrin/discordgo"
  "fmt"
)

func main() {
  discord, err := discordgo.New("n","n")
  fmt.Println(discord)
  fmt.Println(err)
  fmt.Println("Loaded") 
}
