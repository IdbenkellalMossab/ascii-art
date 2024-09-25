package ascii_art


// Declare Sentence object as:
type Sentence struct {
	Input  [][]string
	Output string
}

// Instantiate a new Sentence:
func (file *File) NewSentence(input string) {
	sentence := new(Sentence)
	sentence.Input = make([][]string, len(input))
	for i, char := range input {
		sentence.Input[i] = file.GetLetter(uint(char))
	}
	sentence.Merge()
	file.Parsed = sentence
}

// Form a string from slices:
func (sentence *Sentence) Merge() {
	output := ""
	if len(sentence.Input) == 0 {
		sentence.Output = ""
		return
	}
	merged := make([]string, 9)
	for i := 0; i < len(sentence.Input); i++ {
		for j := 0; j < len(sentence.Input[i]); j++ {
			if merged[j] != "" {
				merged[j] = merged[j][:len(merged[j])-1] + sentence.Input[i][j]
			} else {
				merged[j] = sentence.Input[i][j]
			}
		}
	}

	for _, char := range merged {
		output += string(char)
	}
	sentence.Output = output[1:len(output)-1]
}
