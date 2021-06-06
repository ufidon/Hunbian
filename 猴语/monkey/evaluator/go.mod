module evaluator

go 1.16

replace monkey/object => ../object

replace monkey/ast => ../ast

replace monkey/token => ../token

replace monkey/lexer => ../lexer

replace monkey/parser => ../parser

require (
	monkey/ast v0.0.0-00010101000000-000000000000 // indirect
	monkey/lexer v0.0.0-00010101000000-000000000000 // indirect
	monkey/object v0.0.0-00010101000000-000000000000 // indirect
	monkey/parser v0.0.0-00010101000000-000000000000 // indirect
)
