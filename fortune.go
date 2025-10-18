package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pootwaddle/dadjoke"
	"github.com/pootwaddle/dayplus"
	"github.com/pootwaddle/ljemail"
	"github.com/pootwaddle/shift"
	"github.com/pootwaddle/slogger"
)

func main() {
	var (
		logFile *os.File
		Control ljemail.EmailControl
	)

	slogger.Info("👋 Starting daily dad joke email job")

	joke, err := dadjoke.NewJokes("c:/autojob/fortune.dat")
	if err != nil {
		slogger.Errorf("Could not load dad joke data: %v", err)
		os.Exit(1)
	}
	slogger.Info("✅ Loaded dad joke data")

	Control.From = "bjarvis@laughingj.com"
	Control.ReplyTo = "bjarvis@laughingj.com"
	Control.Recip = "bjarvis@laughingj.com"
	Control.CCRecip = "pootwaddle88@gmail.com"
	Control.BCCRecip = ""
	Control.ProgName = "autojob"
	Control.Layout = ""
	Control.InputFile = "c:/autojob/fortune.dat"
	Control.Subject = joke.DadJokeOfTheDay(time.Now())

	logFileName := ljemail.MailFileName()
	logFile, err = os.Create(logFileName)
	if err != nil {
		// STRUCTURED LOG (fields as JSON in file, if file logging enabled)
		slogger.GetLogger().Error("Unable to create output file",
			"file", logFileName,
			"err", err,
		)
		return
	}
	defer logFile.Close()

	// STRUCTURED LOG (fields as JSON in file)
	slogger.GetLogger().Info("Created output file",
		"file", logFileName,
	)

	ljemail.EmailHeaders(logFile, Control)
	slogger.Debug("Wrote email headers")

	logFile.WriteString("<p>" + joke.DadJokeOfTheDay(time.Now()) + "</p>\n")

	slogger.Debug("Writing main joke and daily stats to email")
	logFile.WriteString("<table>\n")
	logFile.WriteString(fmt.Sprintf("<tr><td>Today's date</td><td>%s</td></tr>\n", joke.Today))
	logFile.WriteString(fmt.Sprintf("<tr><td>Today's FD Shift is:</td><td>%s</td></tr>\n", shift.GetShift(time.Now())))
	logFile.WriteString(fmt.Sprintf("<tr><td>Julian Day</td><td>%d</td></tr>\n", joke.JDay))
	logFile.WriteString(fmt.Sprintf("<tr><td>sysYear</td><td>%d</td></tr>\n", joke.SysYear))
	logFile.WriteString(fmt.Sprintf("<tr><td>lines in file</td><td>%d</td></tr>\n", joke.LinesInFile))
	logFile.WriteString(fmt.Sprintf("<tr><td>years of</td><td>%d</td></tr>\n", joke.LinesInFile/366))
	logFile.WriteString(fmt.Sprintf("<tr><td>extra lines : </td><td>%d</td></tr>\n", joke.LinesInFile-(366*(joke.LinesInFile/366))))
	logFile.WriteString(fmt.Sprintf("<tr><td>yrMod(years of)</td><td>%d</td></tr>\n", joke.YrMod))
	logFile.WriteString(fmt.Sprintf("<tr><td>file line</td><td>%d</td></tr>\n", joke.FileLine))
	//	logFile.WriteString(fmt.Sprintf("<tr><td>2020/09/19 - Dad passed+</td><td>%d</td></tr>\n", int(dayplus.Days(2020, 9, 19, time.Now()))))
	//	logFile.WriteString(fmt.Sprintf("<tr><td>2021/03/18 - Transplant+</td><td>%d</td></tr>\n", int(dayplus.Days(2021, 3, 18, time.Now()))))
	//	logFile.WriteString(fmt.Sprintf("<tr><td>2021/09/20 - Dad Interred+</td><td>%d</td></tr>\n", int(dayplus.Days(2021, 9, 20, time.Now()))))
	//	logFile.WriteString(fmt.Sprintf("<tr><td>2025/03/20 - Mom passed</td><td>%d</td></tr>\n", int(dayplus.Days(2025, 3, 20, time.Now()))))
	logFile.WriteString(fmt.Sprintf("<tr><td>2020/09/19 - Dad passed+</td><td>%s</td></tr>\n", dayplus.ElapsedTime("2020/09/19")))
	logFile.WriteString(fmt.Sprintf("<tr><td>2021/03/18 - Transplant+</td><td>%s</td></tr>\n", dayplus.ElapsedTime("2021/03/18")))
	logFile.WriteString(fmt.Sprintf("<tr><td>2021/09/20 - Dad Interred+</td><td>%s</td></tr>\n", dayplus.ElapsedTime("2021/09/20")))
	logFile.WriteString(fmt.Sprintf("<tr><td>2025/03/20 - Mom passed</td><td>%s</td></tr>\n", dayplus.ElapsedTime("2025/03/20")))
	logFile.WriteString("</table>\n")
	logFile.Sync()

	// ljemail.Footer(logFile)
	logFile.WriteString("</body>\n</html>\n")

	slogger.Info("✅ Wrote complete email HTML")
	// STRUCTURED LOG (fields as JSON in file)
	slogger.GetLogger().Info("✅ Daily dad joke email job completed successfully",
		"output_file", logFileName,
	)
}
