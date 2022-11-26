// Package numconv implements logic for manipulating
// and transforming slices.
package numconv

// IntToStrings accepts an integer n, and returns a slice of strings or an error.
// (e.g. 123456 -> ["abc","efg"])
func IntToStrings(n int, m func(xs []int) (string, error)) ([]string, error) {

	// declared before initializing so it can be called recursively
	var itsFunc func(int, []int, []string, int) ([]string, error)

	itsFunc = func(n int, ns []int, cs []string, iter int) ([]string, error) {
		// there definitely are still values to process
		if n != 0 {
			
			i := n % 10
			ns = append([]int{i}, ns...)
			if iter%3 == 0 {
				c, err := m(ns)
				if err != nil {
					return nil, err
				}
				cs = append([]string{c}, cs...)
				ns = []int{}
			}
			iter++
			return itsFunc(n/10, ns, cs, iter)
		}

		// we've reached the end, but there may be a couple more values to process
		if len(ns) != 0 {
			c, err := m(ns)
			if err != nil {
				return nil, err
			}
			cs = append([]string{c}, cs...)
			ns = []int{}
		}
		return cs, nil
	}

	return itsFunc(n, []int{}, []string{}, 1)
}
