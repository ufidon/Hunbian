module monkey

go 1.16

replace monkey/repl => ./repl

replace monkey/lexer => ./lexer

replace monkey/token => ./token

replace monkey/ast => ./ast

replace monkey/parser => ./parser

replace monkey/object => ./object

replace monkey/evaluator => ./evaluator

require (
	monkey/ast v0.0.0-00010101000000-000000000000 // indirect
	monkey/evaluator v0.0.0-00010101000000-000000000000 // indirect
	monkey/object v0.0.0-00010101000000-000000000000 // indirect
	monkey/parser v0.0.0-00010101000000-000000000000 // indirect
	monkey/repl v0.0.0-00010101000000-000000000000 // indirect
)
