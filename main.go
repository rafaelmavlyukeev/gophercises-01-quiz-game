package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fileLine struct {
	question string
	answer   int
}

func parseFileLine(line string) (fileLine, error) {
	question := ""
	answer := 0

	answerIndex := strings.LastIndex(line, ",")
	if answerIndex == -1 {
		err := errors.New("Wrong question format: missing answer. " + line)
		return fileLine{}, err
	}

	answerStr := line[answerIndex+1:]
	answer, err := strconv.Atoi(answerStr)
	if err != nil {
		err := errors.New("Wrong question format")
		return fileLine{}, err
	}

	if line[answerIndex-1] == '?' {
		question = line[:answerIndex]
	} else {
		question = line[:answerIndex] + "?"
	}

	return fileLine{
		question: question,
		answer:   answer,
	}, nil
}

func main() {

	corr := 0
	incorr := 0

	// read file
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	reader := bufio.NewReader(os.Stdin)

	for scanner.Scan() {
		fileLineParsed, err := parseFileLine(scanner.Text())

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		fmt.Println(fileLineParsed.question)
		userAnswerInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
		}

		userAnswer, err := strconv.Atoi(strings.TrimSuffix(userAnswerInput, "\n"))
		if err != nil {
			fmt.Println(err.Error())
		}

		if userAnswer == fileLineParsed.answer {
			corr++
			fmt.Println("Correct answer")
		} else {
			incorr++
			fmt.Println("Wrong answer")
		}
	}

	fmt.Println("End of quiz. Your Score:", corr, "out of", corr+incorr)

}
