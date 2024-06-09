go run . --output=banner.txt "hello" standard | cat -e

go run . "hello" standard | cat -e

go run . --output=banner1.txt  standard | cat -e

go run . "hello"  | cat -e

go run .  standard shadow | cat -e

go run . --output=banner2.txt "hello" | cat -e

go run . | cat -e


go run . --output=banner3.txt 'Hello There!' shadow | cat -e

go run . --output test00.txt banana standard | cat -e

go run . --output=test00.txt "First\nTest" shadow | cat -e

go run . --output=test01.txt "hello" standard | cat -e

go run . --output=test02.txt "123 -> #$%" standard | cat -e

go run . --output=test03.txt "432 -> #$%&@" shadow | cat -e

go run . --output=test04.txt "There" shadow | cat -e

go run . --output=test05.txt "123 -> \"#$%@" thinkertoy | cat -e

go run . --output=test06.txt "2 you" thinkertoy | cat -e

go run . --output=test07.txt 'Testing long output!' standard | cat -e