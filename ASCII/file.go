package ascii_art

import (
	"os"
)

// Declare a file object:
type File struct {
	Name    string
	Data    []byte
	Symbols []string
	Parsed  *Sentence
	Result  string
}

// Instantiate a new file object:
func NewFile(name string) *File {
	new_file := new(File)
	new_file.Name = name
	data, err := os.ReadFile(new_file.Name)
	if err != nil {
		panic(err)
	}
	new_file.Data = data
	new_file.Symbols = new_file.Extract()
	new_file.Parsed = nil
	new_file.Result = ""
	return new_file
}

// Create a method to extract the leters from my file:
func (file *File) Extract() []string {
	fragment := ""
	symbols := []string{}
	for _, bit := range file.Data {

		fragment += string(bit)
		if bit == 10 {
			if fragment != "" {
				symbols = append(symbols, fragment)
				fragment = ""
			}
		}
	}
	return symbols
}

func (file *File) GetLetter(key uint) []string {
	hash := 855 - ((127 - key) * 9)
	return file.Symbols[hash : hash+9]
}

// split the user input based on the occurrence of the <\n>
func (file *File) SplitAndExec(in string) string {
	str := ""
	result := ""
	happend := false
	for i := 0; i < len(in); i++ {
		if i < len(in)-1 && rune(in[i]) == 92 && in[i+1] == 110 {
			happend = true
		}
		if happend {
			file.NewSentence(str)
			if i != len(in)-1 {
				result += file.Parsed.Output  + "\n"

			} else {
				result += file.Parsed.Output
			}
			str = ""
			happend = false
			i+=1
			continue
		}
		str += string(rune(in[i]))
	}
	if !happend {
		file.NewSentence(str)
		if str != "" {
			if file.Parsed.Output != "" {
				result += file.Parsed.Output
			}
			str = ""
		}
	}
	file.Result = result
	return file.Result
}
