package darts

func Score(x, y float64) int {
    calc := (x*x)+(y*y)
		switch{
            case calc > 100:
            	return 0
            case calc > 25:
            	return 1 
            case calc > 1:
            	return 5 
            case calc > -1:
            	return 10 
            default:
            	return 0
        }
}
