package substringsOfSizeThreeWithDistinctCharacters

func countGoodSubstrings(s string) int {
    var ctr int
    for i := 0; i < len(s)-2; i++ {
        if s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2] {
            ctr++
        }
    }
    return ctr
}