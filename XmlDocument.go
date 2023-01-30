package xmlCreation

type XmlDocument struct {
	Elements []XElement
}

func (doc *XmlDocument) ToString() string {
	var finalDoc string
	for _, element := range doc.Elements {
		finalDoc += element.ToString(0)
	}
	return finalDoc
}

func CreateXml(elements []XElement) XmlDocument {
	var doc XmlDocument = XmlDocument{}
	doc.Elements = append(doc.Elements, elements...)
	return doc
}
