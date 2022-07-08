package positioner

type RecordPositioner interface {
	GetLatest() OraclePosition
	Persist(position OraclePosition)
}
