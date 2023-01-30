package xmlCreation

type Attribute struct {
	Name    string
	Content string
}

func generateElementNameWithAttrs(element XElement) string {
	var finalName string
	unclosedName := "<" + element.GetName()
	attrs := element.GetAttributes()
	if attrs == "" {
		finalName = unclosedName + ">"
		return finalName
	}
	finalName = unclosedName + " " + attrs + ">"
	return finalName
}

type XElement interface {
	GetName() string
	GetAttributes() string
	ToString(tabCount int) string
}
