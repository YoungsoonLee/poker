package cmd

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/YoungsoonLee/poker/poker"
	"github.com/YoungsoonLee/poker/types"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "poker-cli {command} ",
	Short: "This is cli tool for pocker game.",
	Long:  `This is cli tool for pocker game.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(randomRsCmd)

	randomMultiHandsCmd := randomMtCmd()
	rootCmd.AddCommand(randomMultiHandsCmd)

	rootCmd.AddCommand(promptCmd)
}

// randomRsCmd represents the command for generating a random single hand and evaluating it.
var randomRsCmd = &cobra.Command{
	Use:   "rs",
	Short: "Random-Single-Hand: Generate random a hand and evaluate",

	Run: func(cmd *cobra.Command, args []string) {
		// get a random card
		hand := poker.RandomCards(1)
		log.Printf("Cards: %s\n", hand.Cards)

		// check valid ranks
		if !hand.HasValidRanks() {
			log.Printf("Invalid ranks: %+v\n", hand)
			return
		}

		// check valid suits
		if !hand.HasValidSuits() {
			log.Printf("Invalid suits: %+v\n", hand)
			return
		}

		sRank, iRank := hand.Evaluate()
		log.Printf("Rank Order: %d, Rank Title: %s\n", iRank, sRank)
	},
}

// randomMtCmd returns a Cobra command for generating random multi hands and evaluating them.
// It takes an integer input as a flag and generates the specified number of random hands.
// It then evaluates the hands and logs the results.
func randomMtCmd() *cobra.Command {
	var input int

	c := &cobra.Command{
		Use:   "rm",
		Short: "Random-Multi-Hands : Generate random multi hands and evaluate",

		Run: func(cmd *cobra.Command, args []string) {

			if input < 1 {
				log.Printf("Invalid input: %d\n", input)
				return
			}

			hands := poker.RandomCardsToHands(input)

			results := poker.EvaluateHands(hands)

			log.Printf("Congrats! Win Hand ID:%d, Rank: %s, RankOrder: %d, Cards: %+v\n", results[0].HandID, results[0].Rank, results[0].RankOrder, results[0].Card)

			for i, result := range results {
				log.Printf("Result Rank [%d]. ID:%d, Rank: %s, RankOrder: %d, Cards: %+v\n", i+1, result.HandID, result.Rank, result.RankOrder, result.Card)
			}
		},
	}

	c.Flags().IntVar(&input, "input", 0, "Number of input")
	return c
}

// promptCmd represents the command for creating new hands and evaluating them.
// It takes a string input as a flag and generates the specified number of hands.
// It then evaluates the hands and logs the results.
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Prompt: Create new hands and create new cards by each hand through prompt(your input)",

	Run: func(cmd *cobra.Command, args []string) {
		hands := createNewHands()

		results := poker.EvaluateHands(hands)

		log.Printf("Congrats! Win Hand ID:%d, Rank: %s, RankOrder: %d, Cards: %+v\n", results[0].HandID, results[0].Rank, results[0].RankOrder, results[0].Card)

		for i, result := range results {
			log.Printf("Result Rank [%d]. ID:%d, Rank: %s, RankOrder: %d, Cards: %+v\n", i+1, result.HandID, result.Rank, result.RankOrder, result.Card)
		}
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func promptGetInput(pc promptContent) string {
	// validate
	validate := func(input string) error {
		if input == "" {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	// templates
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	// prompt
	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

func createNewHands() []poker.Hand {
	handsPromptContent := promptContent{
		"Please provide how many hands you want to create ex) 5 ",
		"Input Number of hands >> ",
	}

	inputHands := promptGetInput(handsPromptContent)
	log.Printf("Input hands: %s\n", inputHands)

	if inputHands == "0" {
		log.Printf("Please provide hands more than 0\n")
		os.Exit(1)
	}

	// ignore error
	input, _ := strconv.Atoi(inputHands)
	hands := make([]poker.Hand, input)

	definitionPromptContent := promptContent{
		"Please provide cards you want to create by each hand with comma seperate. It's okay not to sensitive case.But The entered hands and ,(comma) must have the same number of hands,and must be a Rank and Suit pair. ex) if you input hands 3, ex) 3s4h5d6c7s,9H3CTSQSAS,4DASAC7H9C ",
		"Input Cards: ex) 3s4h5d6c7s,9H3CTSQSAS,4DASAC7H9C) >> ",
	}

	cards := promptGetInput(definitionPromptContent)
	log.Printf("Input cards: %s\n", cards)

	// remove space
	cards = strings.ReplaceAll(cards, " ", "")

	arrCards := strings.Split(cards, ",")
	if len(arrCards) != input {
		log.Printf("Please provide the same number as the number of hands\n")
		os.Exit(1)
	}

	for i, card := range arrCards {
		c, err := types.NewCard(card)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		hands[i] = poker.Hand{HandID: i + 1, Cards: c}
	}

	return hands
}
