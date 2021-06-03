module monkey

go 1.16

replace monkey/lexer => ./lexer

replace monkey/repl => ./repl

replace monkey/token => ./token

require monkey/repl v0.0.0-00010101000000-000000000000 // indirect
