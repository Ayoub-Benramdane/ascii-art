go run . "banana" standard abc | cat -e

go run . "hello" standard | cat -e

go run . "hello world" shadow | cat -e

go run . "nice 2 meet you" thinkertoy | cat -e

go run . "you & me" | cat -e

go run . "123" shadow | cat -e

go run . "/(\")" thinkertoy | cat -e

go run . "\"#$%&/()*+,-./" thinkertoy | cat -e

go run . "It's Working" thinkertoy | cat -e

go run . --output=banner.txt 'Hello There!' shadow | cat -e

go run . --output test00.txt banana standard | cat -e

go run . --output=test00.txt "First\nTest" shadow | cat -e

go run . --output=test01.txt "hello" standard | cat -e

go run . --output=test02.txt "123 -> #$%" standard | cat -e

go run . --output=test03.txt "432 -> #$%&@" shadow | cat -e

go run . --output=test04.txt "There" shadow | cat -e

go run . --output=test05.txt "123 -> \"#$%@" thinkertoy | cat -e

go run . --output=test06.txt "2 you" thinkertoy | cat -e

go run . --output=test07.txt 'Testing long output!' standard | cat -e

go run . --color red "banana"

go run . --color=red "hello world"

go run . --color=green "1 + 1 = 2"

go run . --color=yellow "(%&) ??"

go run . --color=orange GuYs "HeY GuYs"

go run . --color=blue B 'RGB()'

