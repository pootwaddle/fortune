package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/pootwaddle/date"
	"github.com/pootwaddle/ljemail"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type fortune struct {
	today         time.Time
	jDay          int
	sysYear       int
	yrMod         int
	fileLine      int
	linesInFile   int
	todaysFortune string
}

func (a *fortune) calc_line(t time.Time) int {
	a.today = t
	a.jDay = date.JDay(t)
	fmt.Println("jDay is ", a.jDay)
	fmt.Println("SysYear(t) is ", date.SysYear(t))
	a.sysYear = date.SysYear(t)
	a.yrMod = date.SysYear(t) % 5
	fmt.Println("yr mod is ", a.yrMod)
	a.fileLine = (a.yrMod * 365) + a.yrMod + a.jDay - 1
	return a.fileLine
}

func main() {

	var (
		logFile *os.File
		Control ljemail.EmailControl
	)

	/* test calc_line

	   t := time.Date(2016, time.January, 1, 15, 0, 0, 0, time.Local)
	   fmt.Println(t)
	   fmt.Println(calc_line(t))
	   t = time.Date(2016, time.December, 31, 15, 0, 0, 0, time.Local)
	   fmt.Println(t)
	   fmt.Println(calc_line(t))
	   t = time.Date(2017, time.January, 1, 15, 0, 0, 0, time.Local)
	   fmt.Println(t)
	   fmt.Println(calc_line(t))
	   t = time.Date(2017, time.December, 31, 15, 0, 0, 0, time.Local)
	   fmt.Println(t)
	   fmt.Println(calc_line(t))
	   t = time.Date(2018, time.January, 1, 15, 0, 0, 0, time.Local)
	   fmt.Println(t)
	   fmt.Println(calc_line(t))
	   t = time.Date(2018, time.December, 31, 15, 0, 0, 0, time.Local)
	   fmt.Println(t)
	   fmt.Println(calc_line(t))
	   t = time.Date(2019, time.January, 1, 15, 0, 0, 0, time.Local)
	   fmt.Println(t)
	   fmt.Println(calc_line(t))
	   t = time.Date(2019, time.December, 31, 15, 0, 0, 0, time.Local)
	   fmt.Println(t)
	   fmt.Println(calc_line(t))
	*/

	f := fortune{}

	ft := f.calc_line(time.Now())

	fmt.Println(ft)

	raw, err := ioutil.ReadFile("c:/autojob/fortune.dat")
	check(err)
	str := string(raw)
	linesArray := strings.Split(str, "\r\n")

	f.todaysFortune = ""
	f.linesInFile = len(linesArray)

	if ft < len(linesArray) {
		f.todaysFortune = linesArray[ft]
	} else {
		f.todaysFortune = "End of file"
	}

	Control.From = "bjarvis@laughingj.com"
	Control.ReplyTo = "bjarvis@laughingj.com"
	Control.Recip = "bjarvis@laughingj.com"
	Control.CCRecip = "pootwaddle@pootwaddle.com"
	Control.BCCRecip = ""
	Control.ProgName = ""
	Control.Layout = ""
	Control.InputFile = "c:/autojob/fortune.dat"
	Control.Subject = f.todaysFortune

	//logFile
	logFileName := ljemail.MailFileName()
	logFile, err = os.Create(logFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create %s\r\n", logFileName)
		return
	}
	defer logFile.Close()

	ljemail.EmailHeaders(logFile, Control)

	logFile.WriteString("<p>" + f.todaysFortune + "</p>\n")

	logFile.WriteString("<table>")
	logFile.WriteString(fmt.Sprintf("<tr><td>Today's date</td><td>%s</td></tr>", f.today))
	logFile.WriteString(fmt.Sprintf("<tr><td>Julian Day</td><td>%d</td></tr>", f.jDay))
	logFile.WriteString(fmt.Sprintf("<tr><td>sysYear</td><td>%d</td></tr>", f.sysYear))
	logFile.WriteString(fmt.Sprintf("<tr><td>yrMod (5)</td><td>%d</td></tr>", f.yrMod))
	logFile.WriteString(fmt.Sprintf("<tr><td>file line</td><td>%d</td></tr>", f.fileLine))
	logFile.WriteString(fmt.Sprintf("<tr><td>lines in file</td><td>%d</td></tr>", f.linesInFile))
	logFile.WriteString(fmt.Sprintf("<tr><td>years of</td><td>%d</td></tr>", f.linesInFile/365))
	logFile.WriteString("</table>")
	logFile.WriteString("</body>\n</html>\n")
	logFile.Sync()
	//	ljemail.Footer(logFile)
	fmt.Printf("Done\r\n")
}
