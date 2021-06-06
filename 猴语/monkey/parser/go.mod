module parser

go 1.16

replace monkey/lexer => ../lexer

replace monkey/token => ../token

replace monkey/ast => ../ast

require (
	monkey/ast v0.0.0-00010101000000-000000000000 // indirect
	monkey/lexer v0.0.0-00010101000000-000000000000 // indirect
	monkey/token v0.0.0-00010101000000-000000000000 // indirect
)
