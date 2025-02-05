package fuzzPgxmock

import "strconv"
import "github.com/pashagolub/pgxmock/v2"


func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {
    
        case 0:
            content := string(bytes)
            pgxmock.QueryMatcherRegexp.Match("mayhem", content)
            return 0

        default:
            content := string(bytes)
            pgxmock.NewResult(content, 15)
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}