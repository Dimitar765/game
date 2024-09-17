package main

import (
	"fmt"
	"github.com/Dimitar765/game/questions"
	"github.com/charmbracelet/lipgloss"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Player structure to store player information
type Player struct {
	Name   string
	Score  int
	Epoch  string
	Spells []string
}

// Mythological Creature structure
type Creature struct {
	Name   string
	Health int
	Damage int
}

// Question structure
type Question struct {
	Epoch    string
	Question string
	Choices  []string
	Correct  int
}

// High Scores
var highScores = map[string][]Player{}

// Function to display ASCII art or game logo
func displayArt() {
	fmt.Println(`
		=============================
		   MYTHOLOGICAL QUIZ GAME
		=============================
	`)
}

// Function to display high scores for a selected epoch
func displayHighScores(epoch string) {
	clearScreen()
	fmt.Println(color.YellowString("High Scores for %s", epoch))
	if len(highScores[epoch]) == 0 {
		fmt.Println("No high scores yet.")
		return
	}
	for i, player := range highScores[epoch] {
		fmt.Printf("%d. %s - %d points\n", i+1, player.Name, player.Score)
	}
}

// Function to ask a multiple-choice question
func askMultipleChoiceQuestion(q Question) bool {
	fmt.Println(lipgloss.NewStyle().Bold(true).Render(q.Question))

	// Display choices
	for i, choice := range q.Choices {
		fmt.Printf("%d. %s\n", i+1, choice)
	}

	// Get player input
	fmt.Print("Enter the number of your answer: ")
	input := getUserInput()

	// Convert input to integer
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(q.Choices) {
		color.Red("Invalid input. Please try again.")
		return askMultipleChoiceQuestion(q)
	}

	// Check if the answer is correct
	if choice-1 == q.Correct {
		color.Green("Correct!")
		return true
	}

	color.Red("Incorrect.")
	return false
}

// Function to start the battle phase
func startBattlePhase(player *Player) {
	fmt.Println(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("5")).Render("Entering the battle phase..."))

	creatures := []Creature{
		{Name: "Hydra", Health: 100, Damage: 10},
		{Name: "Minotaur", Health: 120, Damage: 15},
		{Name: "Cyclops", Health: 150, Damage: 20},
	}

	for stage := 1; stage <= 3; stage++ {
		fmt.Printf("\nStage %d: \n", stage)
		creature := creatures[stage-1]
		battleWithCreature(player, &creature)
		if len(player.Spells) == 0 {
			fmt.Println("You have run out of spells!")
			break
		}
	}
}

// Function to handle battle with a specific creature
func battleWithCreature(player *Player, creature *Creature) {
	fmt.Printf("You are battling %s with %d health and %d damage.\n", creature.Name, creature.Health, creature.Damage)

	for creature.Health > 0 {
		fmt.Printf("Creature: %s | Health: %d\n", creature.Name, creature.Health)
		fmt.Printf("1. Fireball (damage 20)\n2. IceBlast (damage 15)\n3. Thunderstrike (damage 25)\n")
		fmt.Print("Choose a spell to use (enter number): ")
		spellIndex := getUserInput()

		// Convert input to integer
		index, err := strconv.Atoi(spellIndex)
		if err != nil || index < 1 || index > len(player.Spells) {
			color.Red("Invalid choice. Please try again.")
			continue
		}

		// Use selected spell
		spell := player.Spells[index-1]
		performSpellEffect(player, creature, spell)

		// Display updated information
		fmt.Printf("You used %s on %s.\n", spell, creature.Name)
		fmt.Printf("The %s's health is now %d.\n", creature.Name, creature.Health)
		if creature.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", creature.Name)
			player.Spells = append(player.Spells, fmt.Sprintf("NewSpell%d", len(player.Spells)+1)) // Adding a new spell
			fmt.Printf("You received a new spell: NewSpell%d\n", len(player.Spells))
			return
		}

		// Creature attacks
		damageReceived := creature.Damage
		fmt.Printf("The %s attacks you for %d damage!\n", creature.Name, damageReceived)
		fmt.Println("You received", damageReceived, "damage.")

		// For simplicity, assume player health is not tracked in this example
	}
}

// Function to check if a slice contains a specific element
func contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// Function to perform spell effects
func performSpellEffect(player *Player, creature *Creature, spell string) {
	fmt.Printf("You use the spell: %s\n", spell)
	var spellDamage int
	switch spell {
	case "Fireball":
		spellDamage = 20
	case "IceBlast":
		spellDamage = 15
	case "Thunderstrike":
		spellDamage = 25
	default:
		spellDamage = 0
	}

	creature.Health -= spellDamage
	fmt.Printf("The %s's health is now %d.\n", creature.Name, creature.Health)
}

// Style text using lipgloss
func textStyle(text string) string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFA500")).Bold(true)
	return style.Render(text)
}

// Start the quiz stage
func startQuizStage(epoch string) {
	clearScreen()
	displayArt()

	var player Player

	// Get player name
	fmt.Print("Enter your name: ")
	player.Name = getUserInput()
	player.Epoch = epoch
	totalQuestions := 10
	correctAnswers := 0

	// Questions for each epoch
	questions := map[string][]Question{
		"Ancient History": {
			{Question: "Who was the first emperor of Rome?", Choices: []string{"Julius Caesar", "Augustus", "Nero"}, Correct: 1},
			{Question: "What is the largest pyramid in Egypt?", Choices: []string{"Pyramid of Khufu", "Pyramid of Khafre", "Pyramid of Menkaure"}, Correct: 0},
			{Epoch: "Ancient History", Question: "Who was the king of the Greek gods?", Choices: []string{"Hades", "Zeus", "Apollo", "Ares"}, Correct: 1},
			{Epoch: "Ancient History", Question: "What is the name of Thor's hammer?", Choices: []string{"Stormbreaker", "Mjolnir", "Gungnir", "Excalibur"}, Correct: 1},
			{Epoch: "Ancient History", Question: "What was the primary language of ancient Rome?", Choices: []string{"Latin", "Greek", "Hebrew", "Aramaic"}, Correct: 0},
			{Epoch: "Ancient History", Question: "Who was the first emperor of China?", Choices: []string{"Qin Shi Huang", "Han Wudi", "Kublai Khan", "Li Shimin"}, Correct: 0},
			{Epoch: "Ancient History", Question: "Which ancient civilization built the pyramids of Giza?", Choices: []string{"Egyptians", "Mayans", "Romans", "Greeks"}, Correct: 0},
			{Epoch: "Ancient History", Question: "Who was the famous Carthaginian general who fought against Rome?", Choices: []string{"Hannibal", "Scipio", "Julius Caesar", "Alexander the Great"}, Correct: 0},
			{Epoch: "Ancient History", Question: "What is the name of the ancient Greek epic poem attributed to Homer?", Choices: []string{"The Iliad", "The Odyssey", "The Aeneid", "The Argonautica"}, Correct: 1},
			{Epoch: "Ancient History", Question: "Which empire was known for its road network that spanned from Britain to the Middle East?", Choices: []string{"Roman Empire", "Persian Empire", "Mongol Empire", "Ottoman Empire"}, Correct: 0},
			{Epoch: "Ancient History", Question: "Who was the famous philosopher and teacher of Plato?", Choices: []string{"Socrates", "Aristotle", "Pythagoras", "Epicurus"}, Correct: 0},
			{Epoch: "Ancient History", Question: "What was the name of the battle in which Alexander the Great defeated the Persian Empire?", Choices: []string{"Battle of Gaugamela", "Battle of Thermopylae", "Battle of Marathon", "Battle of Salamis"}, Correct: 0},
			{Epoch: "Ancient History", Question: "Which ancient civilization is credited with the invention of writing?", Choices: []string{"Sumerians", "Egyptians", "Phoenicians", "Chinese"}, Correct: 0},
			{Epoch: "Ancient History", Question: "What was the primary purpose of the Colosseum in ancient Rome?", Choices: []string{"Entertainment", "Religious ceremonies", "Political meetings", "Educational purposes"}, Correct: 0},
			{Epoch: "Ancient History", Question: "Who was the Egyptian queen known for her alliances with Julius Caesar and Mark Antony?", Choices: []string{"Cleopatra", "Nefertiti", "Hatshepsut", "Aset"}, Correct: 0},
			{Epoch: "Ancient History", Question: "Which Greek city-state was known for its military-oriented society?", Choices: []string{"Sparta", "Athens", "Corinth", "Thebes"}, Correct: 0},
			{Epoch: "Ancient History", Question: "What was the name of the famous ancient library located in Alexandria?", Choices: []string{"Library of Alexandria", "Library of Pergamum", "Library of Ephesus", "Library of Rome"}, Correct: 0},
			{Epoch: "Ancient History", Question: "Who was the Greek god of the underworld?", Choices: []string{"Hades", "Poseidon", "Zeus", "Apollo"}, Correct: 0},
			{Epoch: "Ancient History", Question: "What is the name of the ancient city located in modern-day Iraq that was an important center of learning?", Choices: []string{"Babylon", "Ur", "Nineveh", "Sumer"}, Correct: 0},
			{Epoch: "Ancient History", Question: "Which Roman general was famously assassinated on the Ides of March?", Choices: []string{"Julius Caesar", "Augustus", "Nero", "Tiberius"}, Correct: 0},
			{Epoch: "Ancient History", Question: "What was the primary language of the Byzantine Empire?", Choices: []string{"Greek", "Latin", "Arabic", "Turkish"}, Correct: 0},
			{Epoch: "Ancient History", Question: "Which ancient civilization is known for constructing the massive stone structures called Moai on Easter Island?", Choices: []string{"Rapa Nui", "Maya", "Inca", "Aztec"}, Correct: 0},
			{Epoch: "Ancient History", Question: "Who was the Roman god of war?", Choices: []string{"Mars", "Jupiter", "Neptune", "Mercury"}, Correct: 0},
			{Epoch: "Ancient History", Question: "What was the name of the Greek hero known for his strength and his twelve labors?", Choices: []string{"Heracles", "Achilles", "Perseus", "Theseus"}, Correct: 0},
			{Epoch: "Ancient History", Question: "What was the name of the famous ancient Greek playwright known for his tragedies?", Choices: []string{"Sophocles", "Euripides", "Aeschylus", "Aristophanes"}, Correct: 0},
			// Add 18 more questions here
		},
		"Medieval History": {
			{Question: "Who was the first king of England?", Choices: []string{"William the Conqueror", "Henry VIII", "Richard the Lionheart"}, Correct: 0},
			{Question: "What was the main language of medieval scholarship?", Choices: []string{"Latin", "French", "German"}, Correct: 0},
			// Add 18 more questions here
		},
		"Modern History": {
			{Question: "Who was the President of the United States during the Civil War?", Choices: []string{"Abraham Lincoln", "Ulysses S. Grant", "Andrew Johnson"}, Correct: 0},
			{Question: "What year did World War I begin?", Choices: []string{"1912", "1914", "1916"}, Correct: 1},
			// Add 18 more questions here
		},
	}

	epochQuestions := questions[epoch]
	if totalQuestions > len(epochQuestions) {
		totalQuestions = len(epochQuestions)
	}

	// Randomly select 10 questions from the epoch
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(epochQuestions), func(i, j int) {
		epochQuestions[i], epochQuestions[j] = epochQuestions[j], epochQuestions[i]
	})
	epochQuestions = epochQuestions[:totalQuestions]

	// Ask the selected questions
	for i := 0; i < totalQuestions; i++ {
		success := askMultipleChoiceQuestion(epochQuestions[i])
		if success {
			correctAnswers++
		}
	}

	player.Score = correctAnswers
	fmt.Printf("You got %d out of %d correct!\n", correctAnswers, totalQuestions)

	// Store high score only if the player scores 7 or more correct answers
	if correctAnswers >= 7 {
		highScores[epoch] = append(highScores[epoch], player)
		color.Green("You qualify for the battle phase!")
		player.Spells = []string{"Fireball", "IceBlast", "Thunderstrike"} // Initial spells
		startBattlePhase(&player)
	} else {
		color.Red("You did not score enough correct answers to face the mythological creature.")
		fmt.Println("1. Retry")
		fmt.Println("2. Exit")
		fmt.Print("> ")

		input := getUserInput()
		switch input {
		case "1":
			epochMenu()
		case "2":
			color.Yellow("Exiting the game. Goodbye!")
			os.Exit(0)
		default:
			color.Red("Invalid choice. Exiting the game.")
			os.Exit(1)
		}
	}
}

// Main menu
func mainMenu() {
	clearScreen()
	displayArt()
	fmt.Println("1. Start Game")
	fmt.Println("2. View High Scores")
	fmt.Println("3. Exit")
	fmt.Print("> ")

	input := getUserInput()
	switch input {
	case "1":
		epochMenu()
	case "2":
		highScoreMenu()
	case "3":
		color.Yellow("Exiting the game. Goodbye!")
		os.Exit(0)
	default:
		color.Red("Invalid choice. Please try again.")
		mainMenu()
	}
}

// Epoch selection menu
func epochMenu() {
	clearScreen()
	fmt.Println("Choose a Historical Epoch:")
	fmt.Println("1. Ancient History")
	fmt.Println("2. Medieval History")
	fmt.Println("3. Modern History")
	fmt.Print("> ")

	input := getUserInput()
	switch input {
	case "1":
		startQuizStage("Ancient History")
	case "2":
		startQuizStage("Medieval History")
	case "3":
		startQuizStage("Modern History")
	default:
		color.Red("Invalid epoch choice.")
		epochMenu()
	}
}

// High score menu to view high scores
func highScoreMenu() {
	clearScreen()
	fmt.Println("Choose a Historical Epoch to view High Scores:")
	fmt.Println("1. Ancient History")
	fmt.Println("2. Medieval History")
	fmt.Println("3. Modern History")
	fmt.Print("> ")

	input := getUserInput()
	switch input {
	case "1":
		displayHighScores("Ancient History")
	case "2":
		displayHighScores("Medieval History")
	case "3":
		displayHighScores("Modern History")
	default:
		color.Red("Invalid epoch choice.")
		highScoreMenu()
	}
	waitForUser()
}

// Utility functions for input and screen handling
func getUserInput() string {
	var input string
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func waitForUser() {
	fmt.Println("\nPress Enter to continue...")
	fmt.Scanln()
}

// Main game entry point
func main() {
	mainMenu()
}
