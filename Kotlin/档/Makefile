all:*.tex
	xelatex --synctex=-1 -shell-escape 主
	bibtex 章1
	bibtex 章2
	#bibtex 章3
	#bibtex 章4
	#bibtex 章5
	xelatex --synctex=-1 -shell-escape 主
	xelatex --synctex=-1 -shell-escape 主
