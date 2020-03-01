package main
import "fmt"
func longestCommonPrefix(strs []string) string {
	result := ""
     if len(strs) == 1 {
        result = strs[0]
        return result
	}
    minlen, maxlen := maxLen(strs)
    if minlen == 0 {
        return result
    }
	for i:= 0; i < maxlen; i++{
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

func maxLen(strs []string) (int, int) {
    var minlen int
    var maxlen int
    for n, str := range strs {
        if n == 1{
            minlen=len(str)
            maxlen=len(str)
            continue
        }
        if len(str) < minlen  {
            minlen = len(str)            
        }
        if len(str) >maxlen {
            maxlen = len(str)
        }
}
    return minlen, maxlen
}

func main(){
	input := []string{"", "b"}
	hoge := ""
	var hogehoge int 
	hogehoge = len(hoge)
	fmt.Printf(string(hogehoge))
	fmt.Printf(longestCommonPrefix(input))
}
