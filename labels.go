package dns

// Holds a bunch of helper functions for dealing with labels.

// SplitLabels splits a domainname string into its labels.
func SplitLabels(s string) []string {
	last := byte('.')
	k := 0
        labels := make([]string, 0)
        s = Fqdn(s)     // Make fully qualified
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			if last == '\\' {
				// do nothing
				break
			}
			labels = append(labels,s[k:i])
			k = i + 1 // + dot
		}
		last = s[i]
	}
	return labels
}

// CompareLabels compares the strings s1 and s2 and
// returns how many labels they have in common.
// The compare start with the right most label and stops at
// the label that is different.
// www.miek.nl and miek.nl have two labels in common: miek and nl
// www.miek.nl and www.bla.nl have one label in common: nl
func CompareLabels(s1, s2 string) (n int) {
	l1 := SplitLabels(s1)
	l2 := SplitLabels(s2)

	x1 := len(l1)-1
	x2 := len(l2)-1
	for {
                if x1 < 0 || x2 < 0 {
                        break
                }
                if l1[x1] == l2[x2] {
                        n++
                } else {
                        break
                }
                x1--
                x2--
	}
        return
}
