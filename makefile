lexer-machine-grpah:
	ragel -Vp ../lexer/lexer-machine.rl -o lexer.dot
	dot lexer.dot -Tpng -o lexer.png
	rm lexer.dot
