package main

import (
	"fmt"
	"strings"
	"sync"
)

//дозаполняем пробелами
func buildSpaceStr(val int) string {
	var b strings.Builder
	for i := 0; i < val; i++ {
		b.WriteRune(' ')
	}
	//b.WriteRune('\n')
	return b.String()
}

/*
// с учетом управлющей последовательности в конце

func justifyText(str string, maxWidth int) (string, error) {
	var strBuild strings.Builder

	// пул ошибок
	var errs []error

	words := strings.Split(str, " ")

	letterCount := 0

	for i := 0; i < len(words); i++ {
		// если пул ошибок больше 10 тогда возврат ошибки и пустой строки
		if len(errs)+1 > 10 {
			for _, err := range errs {
				fmt.Println(err)
			}
			return "", errors.New("To much errors in input string.")
		}
		// пул ошибок
		if len(words[i]) >= maxWidth {
			errs = append(errs, fmt.Errorf("Word %s is higher that maxWidth %d. Loose it", words[i], maxWidth))
			continue
		}
		// проходим по каждому слову считаем коливество символов
		letterCount = letterCount + len(words[i]) + 1 // space тоже символ
		// с учетом управлющей последовательности
		if letterCount < maxWidth {
			strBuild.WriteString(words[i])
			strBuild.WriteRune(' ')
		} else {
			i--
			letterCount = 0
			strBuild.WriteString(buildSpaceStr(maxWidth - letterCount))
		}
	}

	return strBuild.String(), nil
}

// без учета
func justifyTextWithout(str string, maxWidth int) (string, error) {
	var strBuild strings.Builder

	// пул ошибок
	var errs []error

	words := strings.Split(str, " ")

	letterCount := 0

	for i := 0; i < len(words); i++ {
		// если пул ошибок больше 10 тогда возврат ошибки и пустой строки
		if len(errs)+1 > 10 {
			for _, err := range errs {
				fmt.Println(err)
			}
			return "", errors.New("To much errors in input string.")
		}
		// пул ошибок
		if len(words[i]) >= maxWidth {
			errs = append(errs, fmt.Errorf("Word %s is higher that maxWidth %d. Loose it", words[i], maxWidth))
			continue
		}
		// проходим по каждому слову считаем коливество символов
		letterCount = letterCount + len(words[i])
		if letterCount < maxWidth {
			strBuild.WriteString(words[i])
			strBuild.WriteRune(' ')
		} else {
			i--
			letterCount = 0
			strBuild.WriteString(buildSpaceStr(maxWidth - letterCount))
		}
	}

	return strBuild.String(), nil
}

// с переносом строки (-)
func justifyTextСarry(str string, maxWidth int) (string, error) {
	var strBuild strings.Builder

	words := strings.Split(str, " ")

	letterCount := 0

	for i := 0; i < len(words); i++ {
		// проходим по каждому слову считаем количество символов

		// если слово с пробелом помещается
		if letterCount+len(words[i]) <= maxWidth {
			if letterCount+len(words[i]) == maxWidth {
				strBuild.WriteString(words[i])
				strBuild.WriteRune('\n')
			}
			strBuild.WriteString(words[i])
			strBuild.WriteRune(' ')
			letterCount = letterCount + len(words[i]) + 1
		} else if letterCount+len(words[i]) > maxWidth {
			// если слово длинее оставшегося capacity тогда берем slice[]
			// берем сколько еще влезает
			lenSz := letterCount + len(words[i]) - maxWidth
			lenSz--
			strBuild.WriteString(words[i][:lenSz])
			strBuild.WriteString("-\n")
			strBuild.WriteString(words[i][lenSz:])
			strBuild.WriteString(" ")

			letterCount = len(words[i]) - lenSz
		} else if letterCount+len(words[i]) == maxWidth {
			strBuild.WriteString("\n")
			letterCount = 0

		}
	}
	return strBuild.String(), nil
}
*/
// еще одна вариация justifyTheText

func joinLineWithSpace(words []string, start int, end int, num_spaces int) string {
	// Number of words in current line
	numWordsCurLine := end - start + 1
	var strb strings.Builder

	for i := start; i < end; i++ {
		strb.WriteString(words[i])
		numWordsCurLine--

		//numCurSpace := int(math.Ceil(float64(num_spaces / numWordsCurLine)))
		numCurSpace := num_spaces / numWordsCurLine
		strb.WriteString(buildSpaceStr(numCurSpace))
		num_spaces -= numCurSpace
	}
	strb.WriteString(words[end])
	strb.WriteString(buildSpaceStr(num_spaces))

	return strb.String()
}

/*Попробую параллельно*/
// создаем задания
func CreateTask(words []string, start int, end int, num_spaces int) func(ch chan<- string) {
	task := func(ch chan<- string) {
		ch <- joinLineWithSpace(words, start, end, num_spaces)
	}

	return task
}

func justifyParallelTheText(str string, maxWidth int) string {

	var tasks []func(chan<- string)

	var builder strings.Builder
	words := strings.Split(str, " ")

	// слайс заданий
	ch := make(chan string, len(words))

	curLineStart := 0
	numWordsCurLine := 0
	curLineLength := 0

	for i := 0; i < len(words); i++ {
		numWordsCurLine++

		lookahead_line_length := curLineLength + len(words[i]) + (numWordsCurLine - 1)

		if lookahead_line_length == maxWidth {
			tasks = append(tasks, CreateTask(words, curLineStart, i, i-curLineStart))
			//builder.WriteString(fmt.Sprintf("%s\n", jo16
			curLineStart = i + 1

			numWordsCurLine = 0
			curLineLength = 0
		} else if lookahead_line_length > maxWidth {
			tasks = append(tasks, CreateTask(words, curLineStart, i-1, maxWidth-curLineLength))
			//builder.WriteString(fmt.Sprintf("%s\n", joinLineWithSpace(words, curLineStart, i-1, maxWidth-curLineLength)))
			//builder.WriteRune('\n')

			curLineStart = i

			numWordsCurLine = 1

			curLineLength = len(words[i])
		} else {
			curLineLength += len(words[i])
		}
	}
	if numWordsCurLine > 0 {
		tasks = append(tasks, CreateTask(words, curLineStart, len(words)-1, numWordsCurLine-1))
		builder.WriteString(buildSpaceStr(maxWidth - curLineLength - (numWordsCurLine - 1)))
	}

	//горутины
	var wg sync.WaitGroup
	for _, val := range tasks {
		wg.Add(1)

		vv := val

		go func() {
			defer wg.Done()
			vv(ch)
		}()
	}
	wg.Wait()
	close(ch)
	// читаем в билдер
	for str := range ch {
		builder.WriteString(str)
		builder.WriteRune('\n')
	}
	//builder.WriteString(buildSpaceStr(maxWidth - curLineLength - (numWordsCurLine - 1)))

	return builder.String()
}

func justifyTheText(str string, maxWidth int) string {

	var builder strings.Builder
	words := strings.Split(str, " ")

	curLineStart := 0
	numWordsCurLine := 0
	curLineLength := 0

	for i := 0; i < len(words); i++ {
		numWordsCurLine++

		lookahead_line_length := curLineLength + len(words[i]) + (numWordsCurLine - 1)

		if lookahead_line_length == maxWidth {
			builder.WriteString(fmt.Sprintf("%s\n", joinLineWithSpace(words, curLineStart, i, i-curLineStart)))
			//builder.WriteRune('\n')
			curLineStart = i + 1

			numWordsCurLine = 0
			curLineLength = 0
		} else if lookahead_line_length > maxWidth {
			builder.WriteString(fmt.Sprintf("%s\n", joinLineWithSpace(words, curLineStart, i-1, maxWidth-curLineLength)))
			//builder.WriteRune('\n')

			curLineStart = i

			numWordsCurLine = 1

			curLineLength = len(words[i])
		} else {
			curLineLength += len(words[i])
		}
	}
	// Last line is to be left-aligned
	if numWordsCurLine > 0 {
		builder.WriteString(joinLineWithSpace(words, curLineStart, len(words)-1, numWordsCurLine-1))
		builder.WriteString(buildSpaceStr(maxWidth - curLineLength - (numWordsCurLine - 1)))
	}

	return builder.String()
}

// из LeetCode

func middleJustify(words []string, diff int, i int, j int) string {
	spacedNeeded := j - i - 1

	spaces := diff / spacedNeeded
	extraSpaces := diff % spacedNeeded

	var strb strings.Builder
	strb.WriteString(words[i])

	for k := i + 1; k < j; k++ { // O(n)
		var tmp int

		if extraSpaces > 0 {
			tmp = 1
		} else {
			tmp = 0
		}
		extraSpaces--

		spacesToApply := spaces + tmp
		strb.WriteString(strings.Repeat(" ", spacesToApply) + words[k])
	}

	return strb.String()
}

func leftJustify(words []string, diff int, i int, j int) string {
	spacesOnRight := diff - (j - i - 1) // сколько пробелов нужно проставить справа
	var strb strings.Builder

	strb.WriteString(words[i])

	for k := i + 1; k < j; k++ { //O(n)
		strb.WriteString(" " + words[k])
	}

	strb.WriteString(strings.Repeat(" ", spacesOnRight))

	return strb.String()
}

func justifyTheTextWords(words []string, maxWidth int) []string {
	var result []string

	//проход по всему слайсу
	i := 0
	for i < len(words) { // O(pow(N,2))
		j := i + 1
		lineLength := len(words[i])
		for j < len(words) && (lineLength+len(words[j])+(j-i-1) < maxWidth) {
			lineLength += len(words[j])
			j++
		}
		diff := maxWidth - lineLength
		numOfWords := j - i
		if numOfWords == 1 || j >= len(words) {
			result = append(result, leftJustify(words, diff, i, j))
		} else {
			result = append(result, middleJustify(words, diff, i, j))
		}
		i = j
	}
	for _, val := range result {
		fmt.Println(val)
	}

	return result
}

func main() {
	//text := "The self-study lessons in this section are written and organised according to the levels of the Common European Framework of Reference for languages (CEFR). There are different types of texts and interactive exercises that practise the reading skills you need to do well in your studies, to get ahead at work and to communicate in English in your free time."
	//text := "Работа не волк, в лес не убежит. Иногда хочу сок. Может и чаю тоже хочу."
	text := "Geeks for Geek is the best computer science portal for geeks."
	maxWidth := 16

	//fmt.Println(justifyText(text, maxWidth))
	//fmt.Println(justifyTextWithout(text, maxWidth))
	//fmt.Println(justifyTextСarry(text, maxWidth))
	fmt.Println(justifyTheText(text, maxWidth))

	var words = []string{"This", "is", "an", "example", "of", "text", "justification."}
	//text := "This is an example of text justification."
	fmt.Println(justifyTheTextWords(strings.Split(text, " "), maxWidth))
	fmt.Println()
	fmt.Println(justifyTheTextWords(words, maxWidth))
}
