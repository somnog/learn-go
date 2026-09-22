package main

import "fmt"

type Speaker struct {
	Name  string
	Title string
}

type Session struct {
	ID       int
	Title    string
	Speaker  Speaker
	TimeSlot TimeSlot
	Erolled  int
	isEnded  bool
}

type TimeSlot struct {
	Start string
	End   string
}

// func funcName(args int)string {
// }
func (s Session) Display() string {
	return fmt.Sprintf("ID: %d, Title: %s, Speaker: %s, TimeSlot: %s-%s, Erolled: %d, isEnded: %t",
		s.ID, s.Title, s.Speaker.Name, s.TimeSlot.Start, s.TimeSlot.End, s.Erolled, s.isEnded)
}

func main() {
	session := Session{
		ID:    1,
		Title: "Integrated Services",
		Speaker: Speaker{
			Name:  "Ali Ahmed",
			Title: "Senior Developer",
		},

		TimeSlot: TimeSlot{
			Start: "10:00 AM",
			End:   "12:00 PM",
		},
		Erolled: 10,
		isEnded: false,
	}

	sessionInfo := session.Display()

	fmt.Println(sessionInfo)
}

// fmt.Println("Hello, GO!")
// PrintWelcome()
// // How to define variables in go

// var conferenceName string = "Go Conference"

// fmt.Println(conferenceName)
// conference := Conference()
// fmt.Println(conference)
// year := Year()
// fmt.Println(year)

// welcome := WelcomeToSomNOG("Ahmed")
// fmt.Println(welcome)

// Variables
// Strings
// Integers
// Bool
//
// isOpen := true
// var isSomNOGisOpen bool = false
// var attendees int = 10
// var name string = ""
// capability := 140
// maxAttendees := 100
// fmt.Println("isOpen:", isOpen, "isSomNOGisOpen:", isSomNOGisOpen, "attendees:", attendees, "name:", name, "capability:", capability, "maxAttendees:", maxAttendees, "Attendees:")

// // Control Flow
// capacity := 200
// enrolled := 250
// if enrolled < capacity {
// 	fmt.Println("You can enroll!")
// } else {
// 	fmt.Println("Sorry, the conference is full.")
// }

// // Switch
// track := "track3"

// switch track {
// case "track1":
// 	fmt.Println("You are in Network track.")
// case "track2":
// 	fmt.Println("You are in Systems track.")
// case "track3":
// 	fmt.Println("You are in Software track.")
// case "track4":
// 	fmt.Println("You are in Security track.")
// default:
// 	fmt.Println("You are not in any track.")
// }

// for i := 1; i <= 100; i++ {
// 	fmt.Println("Welcome to SomNOG!", i)
// }

// tracks := []string{"Network", "Systems", "Software", "Security"}

// for _, track := range tracks {

// 	fmt.Println("All Tracks: ", track)

// 	// fmt.Println("Track, ", i+1, ":", track)
// }

// fmt.Println("Out side the loop")

// fmt.Println(len(tracks), "tracks available.")

// extended := append(tracks, "AI & Machine Learning")

// fmt.Println(len(extended), "tracks available.")

// for _, track := range extended {
// 	fmt.Println("All Tracks: ", track)
// }
// subtrack := extended[:2]
// fmt.Println("Subtrack: ", subtrack)

// // Maps
// applicants := map[string]any{
// 	"Hassan":       "Network",
// 	"Abdi":         "Systems",
// 	"isRegistered": true,
// 	"isAllowed":    false,
// 	"age":          20,
// }

// for name, track := range applicants {
// 	fmt.Println(name, "is in", track, "track.")
// }

// type Applicant struct {
// 	Name         string `json:"name"`
// 	Track        string `json:"track"`
// 	Age          int    `json:"age"`
// 	IsRegistered bool   `json:"isRegistered"`
// 	IsAllowed    bool   `json:"isAllowed"`
// }

// applicant := Applicant{
// 	Name:         "Hassan",
// 	Track:        "Network",
// 	Age:          20,
// 	IsRegistered: true,
// 	IsAllowed:    false,
// }
// fmt.Println(applicant)
