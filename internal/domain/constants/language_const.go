package constants

type Language string

type LanguageStruct struct {
	Turkish Language
	English Language
}

var Languages = LanguageStruct{
	Turkish: "Turkish",
	English: "English",
}
