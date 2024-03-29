// Package numconv implements logic to convert numbers to other types.
package numconv

// IntToStrings accepts an integer n, and returns a slice of strings or an error.
// (e.g. 123456 -> ["abc","efg"])
func IntToStrings(n int, m func(xs []int) (string, error)) ([]string, error) {
	// Helper function so that we don't have to pass in empty initial values
	// to IntToStrings. Along with the value to be processed, the helper function
	// accepts an empty slice of ints (used to temporarily hold chunks of the number),
	// an empty slice of strings (used to hold chunks that have been mapped to strings).
	var itsFunc func(int, []int, []string) ([]string, error)
	itsFunc = func(n int, ns []int, cs []string) ([]string, error) {
		// there are still values to process
		if n != 0 {
			// n mod 10 to grab the smallest part of number
			i := n % 10
			ns = append([]int{i}, ns...)

			// if length of ns is 3, then we are ready to process a 3-value collection of numbers.
			// (e.g. the "345" or "678" in "12,345,678"
			if len(ns) == 3 {
				c, err := m(ns)
				if err != nil {
					return nil, err
				}
				cs = append([]string{c}, cs...)
				ns = []int{}
			}
			// recursively call itsFunc, dividing n by 10 (and implicitly taking the floor of the result),
			// so we can handle the next smallest value in it
			return itsFunc(n/10, ns, cs)
		}

		// We've reached the end, but there may be a couple more values to process
		// (e.g. the "12" in "12,345,678")
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

	// call helper function with n + initial values and return
	return itsFunc(n, []int{}, []string{})
}
