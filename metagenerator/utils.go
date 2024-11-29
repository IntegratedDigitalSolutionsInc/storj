package metagenerator

import (
	"bufio"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/Netflix/go-expect"
	"github.com/google/goterm/term"
)

var (
	genres = []string{
		"Action",
		"Adventure",
		"Comedy",
		"Drama",
		"Fantasy",
		"Horror",
		"Mystery",
		"Romance",
		"Sci-Fi",
		"Thriller",
	}

	actors = []string{
		"Catherine Missal",
		"Monica Bellucci",
		"Michael Doven",
		"Jasmine Reate",
		"Tom Cruise",
		"Scarlett Johansson",
		"Anna Raadsveld",
		"Jason Statham",
		"Emilia Clarke",
		"Wentworth Miller",
		"Dwayne Johnson",
		"Rose Byrne",
		"Rachel McAdams",
		"Michelle Rodriguez",
		"Tom Hanks",
		"Jake Gyllenhaal",
		"Tom Hardy",
		"Chloë Grace Moretz",
		"Johnny Depp",
		"Arnold Schwarzenegger",
		"Sarah Wayne Callies",
		"Vincent Cassel",
		"Lisa Ulliel",
		"Rachel Weisz",
		"Robin Williams",
		"Chris Hemsworth",
		"Edwige Fenech",
		"Linda Fiorentino",
		"Robert Downey Jr.",
		"Alexandra Daddario",
		"Anthony Hopkins",
		"Claudia Koll",
		"Lucy Liu",
		"Samuel L. Jackson",
		"Kate Beckinsale",
		"Jordana Brewster",
		"Ashley Benson",
		"Charlize Theron",
		"Rosamund Pike",
		"Christian Bale",
		"Jennifer Connelly",
		"Adam Sandler",
		"Paul Walker",
		"Brad Pitt",
		"Amy Adams",
		"Milla Jovovich",
		"Chris Evans",
		"Amber Heard",
		"Edward Norton",
		"Julianne Moore",
		"Carice van Houten",
		"Evangeline Lilly",
		"Michelle Monaghan",
		"Forest Whitaker",
		"Ben Stiller",
		"Carla Gugino",
		"Liam Neeson",
		"Eric Roberts",
		"Sammo Hung",
		"Sigourney Weaver",
		"Sylvester Stallone",
		"Helen Mirren",
		"Chris Pratt",
		"Jude Law",
		"Katherine Heigl",
		"Matthew McConaughey",
		"Jodi Lyn O'Keefe",
		"Richard Gere",
		"Jeremy Renner",
		"Kirsten Dunst",
		"Mark Hamill",
		"Rami Malek",
		"Donnie Yen",
		"Shia LaBeouf",
		"Megan Fox",
		"Maria Bello",
		"Zoe Saldana",
		"Kristen Stewart",
		"Vin Diesel",
		"Tommy Lee Jones",
		"Justin Timberlake",
		"Ariadne Shaffer",
		"Keanu Reeves",
		"Michael Fassbender",
		"Leonard Nimoy",
		"Carrie-Anne Moss",
		"Michael Caine",
		"Natalie Dormer",
		"Jack Black",
		"Jennifer Aniston",
		"Steven Spielberg",
		"Ashley Greene",
		"Colin Firth",
		"Selma Blair",
		"Nicolas Cage",
		"Lacey Chabert",
		"Mark Ruffalo",
		"Clint Eastwood",
		"Sharon Stone",
		"Penélope Cruz",
		"Winona Ryder",
		"Pierce Brosnan",
		"Morgan Freeman",
		"Bruce Willis",
		"Katy Mixon",
		"Sean Connery",
		"Donald Sutherland",
		"Hugh Jackman",
		"Daniel Radcliffe",
		"Danny Trejo",
		"Marion Cotillard",
		"Rebecca Ferguson",
		"Lee Majors",
		"Philip Seymour Hoffman",
		"Julia Stiles",
		"Paul Giamatti",
		"Salma Hayek",
		"Anna Faris",
		"Jon Hamm",
		"Sandra Bullock",
		"Cate Blanchett",
		"John Hurt",
		"Nick Nolte",
		"Christopher Nolan",
		"Alonna Shaw",
		"Dabney Coleman",
		"Dominic Cooper",
		"Anne Hathaway",
		"Sienna Guillory",
		"Denise Richards",
		"George Clooney",
		"Elizabeth Banks",
		"John Malkovich",
		"Diane Lane",
		"Renee Rea",
		"Angelina Jolie",
		"Rachelle Lefevre",
		"Hayden Christensen",
		"Nicole Kidman",
		"Colin Farrell",
		"Kate Winslet",
		"Carmen Electra",
		"Olga Kurylenko",
		"Natalie Portman",
		"Emma Stone",
		"Natalie Martinez",
		"Sean Bean",
		"Ryan Reynolds",
		"Ryan Gosling",
		"Fajah Lourens",
		"Orla Brady",
		"Nina Dobrev",
		"Harrison Ford",
		"Olivia Wilde",
		"Ben Affleck",
		"Noomi Rapace",
		"Fan Bingbing",
		"Jamie Lee Curtis",
		"Tara Elders",
		"Al Pacino",
		"Mila Kunis",
		"Eddie Redmayne",
		"Gerard Butler",
		"Henry Winkler",
		"Amanda Seyfried",
		"Lena Headey",
		"Kristen Wiig",
		"Léa Seydoux",
		"Louise Fletcher",
		"Channing Tatum",
		"Anton Yelchin",
		"Alec Baldwin",
		"Tyler Perry",
		"Peter Dinklage",
		"Sam Neill",
		"Orlando Bloom",
		"Linda Hamilton",
		"Dennis Hopper",
		"Danny Glover",
		"María Valverde",
		"Alexa PenaVega",
		"Daniel Craig",
		"Jean-Claude Van Damme",
		"Mark Boone Junior",
		"Bill Murray",
		"Mark Strong",
		"Kelly Hu",
		"Natalya Vdovina",
		"Julia Roberts",
		"Karl Urban",
		"Edward Furlong",
		"Leonardo DiCaprio",
		"Kate Hudson",
		"Mickey Rourke",
		"Jessica Alba",
		"Gina Gershon",
		"Chris Pine",
		"Emma Watson",
		"Lucy Hale",
		"Claire Danes",
		"Gaspard Ulliel",
		"Rosanna Arquette",
		"Amanda Page",
		"John Goodman",
		"Jon Voight",
		"Cary Guffey",
		"Shane West",
		"Ron Perlman",
		"Cameron Diaz",
		"Brigitte Nielsen",
		"Kim Basinger",
		"Eva Mendes",
		"Demi Moore",
		"Kaley Cuoco",
		"Emily Blunt",
		"Tuba Büyüküstün",
		"Ralph Fiennes",
		"Kevin Spacey",
		"Joaquin Phoenix",
		"Liana Liberato",
		"Harvey Keitel",
		"Kristanna Loken",
		"Eddie Murphy",
		"Shin Eun-Kyung",
		"Robin Tunney",
		"Jennifer Lawrence",
		"Oliver Platt",
		"Kate Mara",
		"50 Cent",
		"Richard Madden",
		"Norman Reedus",
		"Nicholas Hoult",
		"Christopher Lee",
		"Nathan Fillion",
		"Jackie Chan",
		"Jessica Chastain",
		"Brendan Gleeson",
		"Ethan Hawke",
		"Clive Owen",
		"Laurence Fishburne",
		"Shailene Woodley",
		"Shu Qi",
		"Elisha Cuthbert",
		"Vincent D'Onofrio",
		"Dominic Purcell",
		"Bryan Cranston",
		"Peyton List",
		"Robert Swenson",
		"George Miller",
		"Ernest Borgnine",
		"Owen Wilson",
		"Woody Harrelson",
		"Jeremy Irons",
		"Kristen Bell",
		"Claudia Cardinale",
		"Quentin Tarantino",
		"Mia Kirshner",
		"Alice Eve",
		"Dougray Scott",
		"Taron Egerton",
		"Tyrese Gibson",
		"Solène Rigot",
		"Luke Evans",
		"Rebecca Hall",
		"Viggo Mortensen",
		"Izabella Miko",
		"J. Pat O'Malley",
		"Emily Watson",
		"Sam Worthington",
		"Naomi Watts",
		"Rene Russo",
		"Ice Cube",
		"Uma Thurman",
		"Kaya Scodelario",
		"David O'Hara",
		"Jan Sterling",
		"Glenn Thomas Jacobs",
		"Dakota Blue Richards",
		"William Shatner",
		"Will Ferrell",
		"Nora Miao",
		"Emmanuelle Chriqui",
		"Seth Rogen",
		"John Leguizamo",
		"Ioan Gruffudd",
		"Jeremy Sumpter",
		"Charlotte Gainsbourg",
		"Viola Davis",
		"Dan Duryea",
		"Jamie Foxx",
		"Cliff Curtis",
		"Cara Delevingne",
		"James Gandolfini",
		"Ida Lupino",
		"Patrick Wilson",
		"Mel Gibson",
		"James Purefoy",
		"Rachael Leigh Cook",
		"James Cameron",
		"Jonah Hill",
		"Kellan Lutz",
		"Joseph Gordon-Levitt",
		"Ian McKellen",
		"Garrett Hedlund",
		"Amy Poehler",
		"Elijah Wood",
		"Adam Baldwin",
		"Eva Green",
		"Michelle Williams",
		"Lea Thompson",
		"Ivy Chen",
		"Kris Kristofferson",
		"ناهد جبر",
		"Shannon Tweed",
		"Hayley Atwell",
		"Kelly Overton",
		"Emma Roberts",
		"Ken Duken",
		"Ray Stevenson",
		"Tony Leung Ka-Fai",
		"Gwyneth Paltrow",
		"Ray Liotta",
		"Christopher Lloyd",
		"Raquel Welch",
		"Rebecca Harrell Tickell",
		"Jodie Foster",
		"Stellan Skarsgård",
		"Pam Grier",
		"Kim Cattrall",
		"Victor Mature",
		"Alona Tal",
		"Matt Damon",
		"Aaron Paul",
		"Stella Stevens",
		"Charlton Heston",
		"Tommy Flanagan",
		"Jack Nicholson",
		"Nicola Peltz",
		"Katey Sagal",
		"Julie Andrews",
		"Luis Guzmán",
		"Laura Harring",
		"Miles Teller",
		"Beau Bridges",
		"Ali Larter",
		"James Spader",
		"Jon Bernthal",
		"Logan Lerman",
		"Sacha Baron Cohen",
		"Jason Clarke",
		"Tom Wilkinson",
		"Robert De Niro",
		"Lee Pace",
		"Laura Antonelli",
		"Geoffrey Rush",
		"Aishwarya Rai Bachchan",
		"Connie Nielsen",
		"Jim Broadbent",
		"Catherine Zeta-Jones",
		"Paul McGann",
		"Louis C.K.",
		"James McAvoy",
		"Christina Hendricks",
		"John C. Reilly",
		"Seth MacFarlane",
		"Dennis Chan",
		"Ted de Corsia",
		"Robin Wright",
		"Kim Dickens",
		"Yuen Biao",
		"Anne Bancroft",
		"Peter Stormare",
		"Hugh Keays-Byrne",
		"Drew Barrymore",
		"Sophie Marceau",
		"Alain Delon",
		"Susan Sarandon",
		"Danielle Panabaker",
		"Ellen Page",
		"Dolph Lundgren",
		"Hugh Grant",
		"Meg Ryan",
		"Thomas Kretschmann",
		"Terrence Howard",
		"David Hemmings",
		"Angie Harmon",
		"Eric Bana",
		"Saori Hara",
		"Christopher Walken",
		"Selma Ergeç",
		"Malcolm McDowell",
		"Keira Knightley",
		"Jaime Pressly",
		"Bradley Cooper",
		"Alexander Ludwig",
		"Mads Mikkelsen",
		"Lisa Kudrow",
		"Dana Ashbrook",
		"Maureen O'Hara",
		"John Russell",
		"Will Smith",
		"Sebastian Stan",
		"Kat Dennings",
		"Jake Johnson",
		"Kevin Hart",
		"Kabby Hui",
		"Bill Paxton",
		"Robert Duvall",
		"Tim Robbins",
		"James Marsden",
		"Ornella Muti",
		"Felicity Jones",
		"Ron Howard",
		"Zooey Deschanel",
		"Bridget Fonda",
		"George Takei",
		"Erin Cummings",
		"Katie Holmes",
		"Nikolaj Lie Kaas",
		"Emmy Rossum",
		"Amy Smart",
		"Helen Hunt",
		"Simon Pegg",
		"Ed Harris",
		"Valeria Golino",
		"Richard Widmark",
		"Mandy Moore",
		"Tom Hiddleston",
		"Dakota Fanning",
		"Josh Duhamel",
		"Leslie Bibb",
		"Natalie Mendoza",
		"Shelley Winters",
		"Armand Assante",
		"Vanessa Hudgens",
		"Mädchen Amick",
		"Roselyn Sánchez",
		"Jessica Lange",
		"Lana Parrilla",
		"Margot Robbie",
		"Anjelica Huston",
		"Hikari Mitsushima",
		"Leng Hussein",
		"Ninet Tayeb",
		"Ann-Margret",
		"Dave Bautista",
		"Rutger Hauer",
		"Jason Segel",
		"Max Riemelt",
		"Elaine Collins",
		"Melissa McCarthy",
		"Ruzaidi Abdul Rahman",
		"Rowan Atkinson",
		"Heather Graham",
		"Lauren Cohan",
		"Ewan McGregor",
		"Bianca Haase",
		"Denzel Washington",
		"Karoline Herfurth",
		"Christopher Lambert",
		"Alycia Debnam-Carey",
		"Liu Shishi",
		"Paul Bettany",
		"Famke Janssen",
		"Adam G. Sevani",
		"Mae Whitman",
		"Ed Asner",
		"Chris Rock",
		"Jet Li",
		"Emile Hirsch",
		"Ashton Kutcher",
		"Jake McDorman",
		"Julie Dreyfus",
		"Ian Somerhalder",
		"Matthew Lillard",
		"Steve Martin",
		"Renée Zellweger",
		"Selma Stern",
		"Greg Kinnear",
		"James Coburn",
		"Colm Meaney",
		"Elisha Cook Jr.",
		"Richard Harris",
		"Elisabeth Shue",
		"Ti Lung",
		"Naomie Harris",
		"Whoopi Goldberg",
		"Tobey Maguire",
		"Carly Chaikin",
		"Tim Roth",
		"Charlotte Rampling",
		"Pamela Anderson",
		"Toni Collette",
		"Noah Taylor",
		"Tony Curtis",
		"Ben Chaplin",
		"Ray Winstone",
		"C. Thomas Howell",
		"John Cho",
		"Arielle Kebbel",
	}

	languages = []string{
		"en-US", // American English
		"es-ES", // Spanish (Spain)
		"fr-FR", // French (France)
		"de-DE", // German (Germany)
		"ja-JP", // Japanese (Japan)
		"zh-CN", // Chinese (Simplified, China)
		"ar-SA", // Arabic (Saudi Arabia)
		"hi-IN", // Hindi (India)
		"pt-BR", // Portuguese (Brazil)
	}

	resolutions = map[string][]int{
		"Quarter Quarter VGA (QQVGA)":   []int{160, 120},
		"Half QVGA (HQVGA)":             []int{240, 160},
		"HVGA":                          []int{320, 480},
		"Video Graphics Array (VGA)":    []int{640, 480},
		"Extended Graphics Array (XGA)": []int{1024, 768},
	}

	adjectives   []string
	nouns        []string
	verbs        []string
	prefixes     = []string{"The", "A", "One", "When", "Where", "Why", "How", "Who", "What", "If", "As", "While", "Before", "After", "During", "Beyond", "Beneath", "Above", "Below", "Inside", "Outside"}
	conjunctions = []string{"and", "or", "but", "yet", "for", "nor", "so"}
	prepositions = []string{"in", "on", "at", "by", "for", "with", "about", "against", "between", "into", "through", "during", "before", "after", "above", "below", "to", "from", "up", "down", "over", "under"}
)

func prettyPrint(data interface{}) {
	b, _ := json.Marshal(data)

	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")
	fmt.Println(out.String())
}

func putFile(record *Record) error {
	localPath := filepath.Join("/tmp", strings.ReplaceAll(record.Path, "/", "_"))
	record.Path = "sj://" + Label + record.Path

	file, err := os.Create(localPath)
	if err != nil {
		return err
	}
	file.Close()

	// Copy file
	// TODO: rerfactor with uplink library
	cmd := exec.Command("uplink", "cp", localPath, record.Path)
	cmd.Dir = clusterPath
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		return err
	}

	return os.Remove(localPath)
}

func deleteFile(record *Record) error {
	cmd := exec.Command("uplink", "rm", record.Path)
	cmd.Dir = clusterPath
	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))

	return err
}

func UplinkSetup(satelliteAddress, apiKey string) {
	c, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	cmd := exec.Command("uplink", "setup", "--force")
	cmd.Dir = clusterPath
	cmd.Stdin = c.Tty()
	cmd.Stdout = c.Tty()
	cmd.Stderr = c.Tty()

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	c.ExpectString("Enter name to import as [default: main]:")
	c.Send(Label + "\n")
	c.ExpectString("Enter API key or Access grant:")
	c.Send(apiKey + "\n")
	c.ExpectString("Satellite address:")
	c.Send(satelliteAddress + "\n")
	c.ExpectString("Passphrase:")
	c.Send(Label + "\n")
	c.ExpectString("Again:")
	c.Send(Label + "\n")
	c.ExpectString("Would you like to disable encryption for object keys (allows lexicographical sorting of objects in listings)? (y/N):")
	c.Send("y\n")
	c.ExpectString("Would you like S3 backwards-compatible Gateway credentials? (y/N):")
	c.Send("y\n")
	fmt.Println(term.Greenf("Uplink setup done"))
}

func GeneratorSetup(bS, wN, tR int, apiKey, projectId, metaSearchEndpoint string, db *sql.DB, ctx context.Context) {
	// Initialize batch generator
	batchGen := NewBatchGenerator(
		db,
		bS,
		wN,
		tR,
		GetPathCount(ctx, db),
		projectId,
		apiKey,
		DbMode,
		metaSearchEndpoint,
	)

	// Generate and insert/debug records
	//startTime := time.Now()

	if err := batchGen.GenerateAndInsert(ctx); err != nil {
		panic(fmt.Sprintf("failed to generate records: %v", err))
	}

	//fmt.Printf("Generated %v records in %v\n", tR, time.Since(startTime))
}

func Clean() {
	//Remove bucket
	cmd := exec.Command("uplink", "rb", "sj://"+Label, "--force")
	cmd.Dir = clusterPath

	out, err := cmd.CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		panic(err.Error())
	}
}

func randomDuration() time.Duration {
	r := rand.Intn(1800)
	return time.Duration(r) * time.Second
}

func randomGenres() []string {
	gN := len(genres)
	sP := rand.Intn(gN)
	eP := sP + rand.Intn(gN-sP) + 1

	return genres[sP:eP]
}

func randomCast() []string {
	gN := len(actors)
	sP := rand.Intn(gN)
	eP := sP + rand.Intn(gN-sP) + 1

	return actors[sP:eP]
}

func randomLanguage() string {
	return languages[rand.Intn(len(languages))]
}

func randomResolution() (res []int) {
	keys := reflect.ValueOf(resolutions).MapKeys()
	return resolutions[keys[rand.Intn(len(keys))].Interface().(string)]
}

func randomYear() int {
	min := 1888
	max := time.Now().Year()
	return rand.Intn(max-min) + min
}

func fetchWordsFromURL(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var words []string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func initializeWordLists() error {
	var err error

	adjectives, err = fetchWordsFromURL("https://gist.github.com/hugsy/8910dc78d208e40de42deb29e62df913/raw/eec99c5597a73f6a9240cab26965a8609fa0f6ea/english-adjectives.txt")
	if err != nil {
		return fmt.Errorf("failed to fetch adjectives: %v", err)
	}

	nouns, err = fetchWordsFromURL("https://gist.github.com/hugsy/8910dc78d208e40de42deb29e62df913/raw/eec99c5597a73f6a9240cab26965a8609fa0f6ea/english-nouns.txt")
	if err != nil {
		return fmt.Errorf("failed to fetch nouns: %v", err)
	}

	verbs, err = fetchWordsFromURL("https://github.com/aaronbassett/Pass-phrase/raw/master/verbs.txt")
	if err != nil {
		return fmt.Errorf("failed to fetch verbs: %v", err)
	}

	return nil
}

func generateName() string {
	adj := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	prefix := prefixes[rand.Intn(len(prefixes))]
	return fmt.Sprintf("%s %s %s", strings.Title(prefix), strings.Title(adj), strings.Title(noun))
}

func generateCallsign(name string) (callsign string) {
	words := strings.Split(name, " ")
	for _, w := range words {
		callsign = fmt.Sprint(callsign, string(w[0]))
	}
	return
}

func generateSimpleTitle() string {
	adj := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	return fmt.Sprintf("The %s %s", strings.Title(adj), strings.Title(noun))
}

func generateComplexTitle() string {
	titleLength := rand.Intn(3) + 3 // Generate titles with 3-5 words

	var titleParts []string
	titleParts = append(titleParts, prefixes[rand.Intn(len(prefixes))])

	for i := 1; i < titleLength; i++ {
		switch rand.Intn(5) {
		case 0:
			titleParts = append(titleParts, strings.Title(adjectives[rand.Intn(len(adjectives))]))
		case 1:
			titleParts = append(titleParts, strings.Title(nouns[rand.Intn(len(nouns))]))
		case 2:
			titleParts = append(titleParts, strings.Title(verbs[rand.Intn(len(verbs))]))
		case 3:
			titleParts = append(titleParts, conjunctions[rand.Intn(len(conjunctions))])
		case 4:
			titleParts = append(titleParts, prepositions[rand.Intn(len(prepositions))])
		}
	}

	return strings.Join(titleParts, " ")
}

func generateDescription(f, t int) string {
	descriptionLength := rand.Intn(t-f) + 10 // Generate description with 10-30 words

	var descriptionParts []string
	descriptionParts = append(descriptionParts, prefixes[rand.Intn(len(prefixes))])

	for i := 1; i < descriptionLength; i++ {
		switch rand.Intn(5) {
		case 0:
			descriptionParts = append(descriptionParts, strings.Title(adjectives[rand.Intn(len(adjectives))]))
		case 1:
			descriptionParts = append(descriptionParts, strings.Title(nouns[rand.Intn(len(nouns))]))
		case 2:
			descriptionParts = append(descriptionParts, strings.Title(verbs[rand.Intn(len(verbs))]))
		case 3:
			descriptionParts = append(descriptionParts, conjunctions[rand.Intn(len(conjunctions))])
		case 4:
			descriptionParts = append(descriptionParts, prepositions[rand.Intn(len(prepositions))])
		}
	}

	return strings.Join(descriptionParts, " ")
}
