package main
import "fmt"
func longestCommonPrefix(strs []string) string {
	result := ""
     if len(strs) == 1 {
        result = strs[0]
        return result
	}
	var minlen int
    minlen = maxLen(strs)
    if minlen == 0 {
        return result
    }
	for i:= 0; i < minlen; i++{
        before := ""
		for _, str := range strs{
			if before == "" {
				before = getRuneAt(str, i)
				continue
			}
			if getRuneAt(str, i) != before{
				return result
			}
	    
		}
        result += before
	}
    return result
}

func getRuneAt(s string, i int) string {
    rs := []rune(s)
    return string(rs[i])
}

func maxLen(strs []string) int {
    var minlen int
    for n, str := range strs {
        if n == 0{
            minlen=len(str)
            continue
        }
        if len(str) < minlen {
            minlen = len(str)
        }
}
    return minlen
}

func main(){
	input := []string{"a", "aa"}
	hoge := ""
	var hogehoge int 
	hogehoge = len(hoge)
	fmt.Printf(string(hogehoge))
	fmt.Printf(longestCommonPrefix(input))
}
