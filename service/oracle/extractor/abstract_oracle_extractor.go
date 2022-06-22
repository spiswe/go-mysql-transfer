package extractor

type AbstractOracleRecordExtractor struct {
	extractor AbstractRecordExtractor
}

func (a *AbstractOracleRecordExtractor) Extractor() AbstractRecordExtractor {
	return a.extractor
}

func (a *AbstractOracleRecordExtractor) SetExtractor(extractor AbstractRecordExtractor) {
	a.extractor = extractor
}

//func (a *AbstractRecordExtractor) getColumnValue()
