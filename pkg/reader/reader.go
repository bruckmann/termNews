package newsParser

import (
	"encoding/xml"
	"fmt"

)



func parseXmlFileNewsToXml(file *[]byte) xml.CharData {
 
  parsedFile, err := xml.Marshal(file)
  
  if err != nil {
    fmt.Printf("Error to created new xml based on the file: %s", err)
  }

  return parsedFile
}


func ParseNewsToString(newsFile *xml.CharData) string {
  return ""
}

