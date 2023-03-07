package irdata

type IRDataType int

const (
	IRDataTypeRaw IRDataType = iota
)

var IRDataTypeMap = map[IRDataType]string{
	IRDataTypeRaw: "Raw",
}

type RawIRData []byte
