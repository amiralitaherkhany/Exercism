
package proverb

import "fmt"


func Proverb(rhyme []string) []string {                                   
	verb := make([]string, 0, 7)
	if len(rhyme) == 0{
        return verb
    }
	for i := 0; i < len(rhyme)-1; i++ {
		verb = append(verb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}
	verb = append(verb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return verb
}