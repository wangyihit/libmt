package servers


type BytesProcessor interface {
	Run(buffer []byte, length int) int
}
