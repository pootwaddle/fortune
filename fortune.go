package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pootwaddle/dadjoke"
	"github.com/pootwaddle/ljemail"
)

func main() {

	var (
		logFile *os.File
		Control ljemail.EmailControl
	)

	joke, err := dadjoke.NewJokes("c:/autojob/fortune.dat")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	Control.From = "bjarvis@laughingj.com"
	Control.ReplyTo = "bjarvis@laughingj.com"
	Control.Recip = "bjarvis@laughingj.com"
	Control.CCRecip = "pootwaddle@pootwaddle.com"
	Control.BCCRecip = ""
	Control.ProgName = ""
	Control.Layout = ""
	Control.InputFile = "c:/autojob/fortune.dat"
	Control.Subject = joke.DadJokeOfTheDay(time.Now())

	//logFile
	logFileName := ljemail.MailFileName()
	logFile, err = os.Create(logFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create %s\r\n", logFileName)
		return
	}
	defer logFile.Close()

	ljemail.EmailHeaders(logFile, Control)

	logFile.WriteString("<p>" + joke.DadJokeOfTheDay(time.Now()) + "</p>\n")

	logFile.WriteString("<table>\n")
	logFile.WriteString(fmt.Sprintf("<tr><td>Today's date</td><td>%s</td></tr>\n", joke.Today))
	logFile.WriteString(fmt.Sprintf("<tr><td>Julian Day</td><td>%d</td></tr>\n", joke.JDay))
	logFile.WriteString(fmt.Sprintf("<tr><td>sysYear</td><td>%d</td></tr>\n", joke.SysYear))
	logFile.WriteString(fmt.Sprintf("<tr><td>yrMod (5)</td><td>%d</td></tr>\n", joke.YrMod))
	logFile.WriteString(fmt.Sprintf("<tr><td>file line</td><td>%d</td></tr>\n", joke.FileLine))
	logFile.WriteString(fmt.Sprintf("<tr><td>lines in file</td><td>%d</td></tr>\n", joke.LinesInFile))
	logFile.WriteString(fmt.Sprintf("<tr><td>years of</td><td>%d</td></tr>\n", joke.LinesInFile/365))
	logFile.WriteString(fmt.Sprintf("<tr><td>extra lines : </td><td>%d</td></tr>\n", joke.LinesInFile-366*(joke.LinesInFile/365)))
	logFile.WriteString("</table>\n")
	logFile.Sync()
	//	ljemail.Footer(logFile)

	logFile.WriteString("</body>\n</html>\n")
	logFile.Close()
	fmt.Printf("Done\r\n")
}