package resource

/*
This file is autogenerated, do not edit;
changes will be undone by the next 'generate' command.

Updates to this type are made my editing the schema file
and executing the 'generate' command
*/

// Model is autogenerated from the json schema
type Model struct {
	TPSCode            *string  `json:",omitempty"`
	Title              *string  `json:",omitempty"`
	CoverSheetIncluded *bool    `json:",omitempty"`
	DueDate            *string  `json:",omitempty"`
	ApprovalDate       *string  `json:",omitempty"`
	Memo               *Memo    `json:",omitempty"`
	SecondCopyOfMemo   *Memo    `json:",omitempty"`
	TestCode           *string  `json:",omitempty"`
	Authors            []string `json:",omitempty"`
}

// Memo is autogenerated from the json schema
type Memo struct {
	Heading *string `json:",omitempty"`
	Body    *string `json:",omitempty"`
}