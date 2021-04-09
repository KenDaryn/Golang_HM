//1)https://github.com/urfave/cli/blob/master/docs/v1/manual.md
//почитать и ознакомиться
//что такое команды и как их создаваться (commands)
//а так же как создвать флаги для каждой команды
//2)
//необходимо создать приложение которое имеет два функционала
//1) калькулятор
//go run main.go calc -z=+ -n1=3 -n2=4
//func Calc()

//2)ауидо плеер
//go run main.go player -f=1.mp3

package main

import (
	"errors"
	"fmt"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/urfave/cli"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	znak string
	num1 int
	num2 int
)

func main() {
	app := cli.NewApp()
	app.Name = "App for downloading video"
	app.Description = "Just set the url and video files will bee downloaded"
	app.Usage = "Set parameter url to link"
	app.Commands = []cli.Command{
		{
			Name:  "Player",
			Usage: "Play mp3 file",
			Action: func(c *cli.Context) error {
				f, err := os.Open("11.mp3")
				if err != nil {
					log.Fatal(err)
				}
				st, format, err := mp3.Decode(f)
				if err != nil {
					log.Fatal(err)
				}
				defer st.Close()
				err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/1))
				speaker.Play(st)
				select {}
			},
		},
		{
			Name:  "Calc",
			Usage: "Play mp3 file",
			Action: func(c *cli.Context) error {
				znak := c.Args().Get(0)
				a := c.Args().Get(1)
				b := c.Args().Get(2)
				if c.NArg() == 3 {
					total := 0
					num1, err := strconv.Atoi(a)
					if err != nil {
						return err
					}
					num2, err := strconv.Atoi(b)
					if err != nil {
						return err
					}
					if znak == "+" {
						total = num1 + num2
					}
					if znak == "-" {
						total = num1 - num2
					} else {
						return errors.New("первый аргумент должен быть + или -")
					}
					fmt.Println(total)
				} else {
					return errors.New("dolzhno byt 3 argumenta")
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
