package strand
import "strings"
func ToRNA(dna string) string {
	var buf strings.Builder
	an := map[rune]rune{
		'G': 'C',
		'T': 'A',
		'C': 'G',
		'A': 'U',
	}

	for _, v := range dna {
		buf.WriteRune(an[v])
	}

	return buf.String()
}
