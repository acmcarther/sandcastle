package main

import (
	"github.com/go-chat-bot/bot/irc"
	_ "github.com/go-chat-bot/plugins/crypto"
	_ "github.com/go-chat-bot/plugins/godoc"
	_ "github.com/go-chat-bot/plugins/treta"
	"github.com/go-chat-bot/bot"
	"os"
	"fmt"
	"os/exec"
)

func scala(command *bot.Cmd) (msg string, err error) {

	if command.Channel != "#devs" {
		return "nah dude", nil
	}
	f, err := os.Create("/tmp/scala.scala")
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.WriteString("object T { def main(args: Array[String]): Unit = {\n")

	if err != nil {
		return "", err
	}

	_, err = f.WriteString(command.RawArgs)

	if err != nil {
		return "", err
	}

	_, err = f.WriteString("} } \n")

	if err != nil {
		return "", err
	}

	cmd := exec.Command("scala", "/tmp/scala.scala")
	stdOutStdErr, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Sprintf("Scala had err: %s with message: %s", err.Error(), stdOutStdErr), nil
	}

	fmt.Printf("output: [%s]\n", stdOutStdErr)
	return fmt.Sprintf("Scala says: %s", stdOutStdErr), nil
}

func main() {
	bot.RegisterCommand(
		"scala",
		"runs a scalamadoo",
		"println(\"do your stuff here\")",
		scala)
	irc.Run(&irc.Config{
		Server:   "irc.iximeow.net:6697",
		Channels: []string{"#devs"},
		User:     "sandbot",
		Nick:     "sandbot",
		Password: "itjustworks",
		UseTLS:   true,
		Debug:    os.Getenv("DEBUG") != ""})
}
