package canal

var (
	systemSchemas = []string{"mysql", "information_schema"}
)

func isSystemSchema(schema string) bool {
	for _, item := range systemSchemas {
		if item == schema {
			return true
		}
	}
	return false
}
