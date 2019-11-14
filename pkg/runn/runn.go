package runn

// Execute statements as stop when error
func Execute(stmts ...interface{}) error {
	return Executor{StopWhenError: true}.Execute(stmts...)
}
