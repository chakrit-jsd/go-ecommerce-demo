package utils

func CustomQuery(name string, categoryId int) (qName, qCat []interface{}) {
	qName = make([]interface{}, 2)
	qCat = make([]interface{}, 2)

	if name != "" {
		qName[0] = "name LIKE ?"
		qName[1] = "%" + name + "%"
	}

	if categoryId != 0 {
		qCat[0] = "category_id = ?"
		qCat[1] = categoryId
	}

	return qName, qCat

}
