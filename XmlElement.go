package xmlCreation

type XmlElement struct {
	Name           string
	Attributes     []Attribute
	ElementContent []XElement
	FinalElement   BaseElement
}

func (element *XmlElement) GetName() string {
	return element.Name
}

func (element *XmlElement) GetAttributes() string {
	var stringAttributes string
	if len(element.Attributes) < 1 {
		return ""
	}
	for _, attr := range element.Attributes {
		stringAttributes += attr.Name + "=" + attr.Content
	}
	return stringAttributes
}

func (element *XmlElement) GetContent() []XElement {
	return element.ElementContent
}

func (element *XmlElement) ToString(tabCount int) string {

	var baseTabs string
	for i:=1; i<=tabCount; i++ {
		baseTabs += "\t"
	}

	name :=  generateElementNameWithAttrs(element)
	var finalContent string
	for i, content := range element.ElementContent {

		finalContent += "\n" + baseTabs + content.ToString(i + 1)
	}
	var closingTag string
	if len(element.ElementContent) > 0 {
		closingTag += "\n"
	}
	closingTag += baseTabs + "<" + element.Name + ">"

	return baseTabs + name + finalContent + closingTag
}

func NewXmlElement(name string, attributes []Attribute, elementContent []XElement) *XmlElement {
	element := XmlElement{Name: name, Attributes: attributes, ElementContent: elementContent}
	return &element
}
