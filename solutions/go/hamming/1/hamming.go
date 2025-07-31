package hamming

import (
    "errors"
)

func Distance(a, b string) (int, error) {
    hamm := 0
	if len(a) != len(b){
        return 0, errors.New("")
    }
    for i := 0 ; i < len(a) ; i ++{
        if a[i] != b[i]{
            hamm++
        }
    }
    return hamm, nil
}
