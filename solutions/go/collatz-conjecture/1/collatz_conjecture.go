package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
        return 0, fmt.Errorf("must be a positive integer")
    }
    steps := 0
    for n != 1 {
        if n % 2 == 0 {
            n /= 2
        } else {
            n = n*3 + 1
        }
        steps++
    }
    return steps, nil
}
