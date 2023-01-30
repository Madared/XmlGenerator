package xmlCreation

type BaseElement struct {
	Name       string
	Attributes []Attribute
	Content    string
}

func (element *BaseElement) GetName() string {
	return element.Name
}
func (element *BaseElement) GetAttributes() string {
	var attrString string
	for _, attr := range element.Attributes {
		attrString += attr.Name + "=" + attr.Content
	}
	return attrString
}

func (element *BaseElement) GetContent() string {
	return element.Content
}

func (element *BaseElement) ToString(tabCount int) string {
	var totalTabs string
	for i:=1; i==tabCount; i++ {
		totalTabs += "\t"
	}
	finalName := "<" + element.Name
	nameWithAttrs := generateElementNameWithAttrs(element)
	if len(nameWithAttrs) > 0 {
		finalName += " " 
		for _, attr := range(element.Attributes) {
			finalName += attr.Name + "=" + attr.Content
		}
	}
	return totalTabs + finalName + element.Content + "<" + element.Name + ">"
}
