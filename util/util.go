package util

func Printerr(_err error) {
	if _err != nil {
		panic(_err)
	}
}
