package extractor

import "go-mysql-transfer/service/oracle/models"

const (
	FORMAT        string = "select * from (select {0} from {1}.{2} t where {3} > ? order by {3} asc) where rownum <= ?"
	MIN_PK_FORMAT string = "select min({0}) from {1}.{2}"
)

type OracleFullRecordExtractor struct {
	absOraRecExt  AbstractOracleRecordExtractor
	oracleContext models.OracleContext
}
