// In package mathutil
package somesubpackage

// Exported function
func Add(a, b int) int {
	return a + b
}

// Unexported function
func subtract(a, b int) int {
	return a - b
}

// Exported struct
type Calculator struct {
	// Exported field
	Result int
}

// Exported method
func (c *Calculator) Add(values ...int) {
	for _, value := range values {
		c.Result += value
	}
}
