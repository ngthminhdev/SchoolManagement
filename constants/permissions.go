package constants

const OS_READ_WRITE_FILE_PERMISSION = 0666

type Gender int

const (
  MALE Gender = iota + 1
  FEMALE
  OTHER
)

