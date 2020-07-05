package main
// ------------ ByPython Learner for GoLang on 05-JUL-20
import (
	"fmt"
	"io"
	"os"
)

func createfile(name string) {
	fmt.Println("creating a file")
	fmt.Println(os.Getwd())

	newfile, error := os.Create(name)
	defer newfile.Close()
	if error != nil {
		fmt.Println("Error in creating a file")
	} else {
		fmt.Println("File got created")
	}

}

func writeinfile(name string) {
	data := `Hi Team,
	This is the content of a file 
	Or(CheckForIncorrectCharacters($full_record_line), Equal($event_datetime, DateTime(9999, 01, 01, 0, 0, 0, 0))), 
	IfForBoolean(Or(Equal(RouteSwitchEventDateLookup(MSC_ID, MSC_ID, $inRoute, $outRoute, $event_datetime), true ),
	Equal(RouteSwitchEventDateLookup($Parent_Switch, $Parent_Switch, $inRoute, $outRoute, $event_datetime), true)),
	Or(And(Equal(RouteSwitchEventDateLookup(MSC_ID, MSC_ID, $inRoute, $outRoute, $event_datetime), true), 
	Equal(RouteTransitFlDateLookUp(MSC_ID, MSC_ID, $inRoute, $outRoute, $event_datetime), true)),
	 And(Equal(RouteSwitchEventDateLookup($Parent_Switch, $Parent_Switch, $inRoute, $outRoute, $event_datetime), true), 
	 Equal(RouteTransitFlDateLookUp($Parent_Switch, $Parent_Switch, $inRoute, $outRoute, $event_datetime), true))),Or(And(Equal
		( RouteSwitchEventDateLookup(MSC_ID, MSC_ID, ChangeInRouteValue($inRoute, MSRN), $outRoute, $event_datetime), true),
		Equal( RouteTransitFlDateLookUp(MSC_ID, MSC_ID, ChangeInRouteValue($inRoute, MSRN), $outRoute, $event_datetime), true)),
		And(Equal( RouteSwitchEventDateLookup($Parent_Switch, $Parent_Switch, ChangeInRouteValue($inRoute, MSRN), 
		$outRoute, $event_datetime), true),Equal( RouteTransitFlDateLookUp($Parent_Switch, $Parent_Switch, 
			ChangeInRouteValue($inRoute, MSRN), $outRoute, $event_datetime), true))))
	,GreaterThan(SecondDiff($event_datetime,Now() ),10368000),
	Not(ValidateDateTime($event_datetime)))
	`
	file, _ := os.OpenFile(name, os.O_RDWR, 0666)
	defer file.Close()
	file.WriteString(data)
	fmt.Println("File got updated")

}

func readfile(name string) {
	readfilee, _ := os.Open(name) //os.OpenFile(name, os.O_WRONLY, 0666)
	bytestring := make([]byte, 20)
	for {
		_, err := readfilee.Read(bytestring)
		//fmt.Println(bytes)
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			break
		}
		fmt.Print(string(bytestring))
	}

}

func main() {
	filename := "testfile"
	createfile(filename)
	writeinfile(filename)
	readfile(filename)

}
