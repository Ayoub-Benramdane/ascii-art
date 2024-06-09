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

go run . --color=red hello --output=t.txt hello shadow

go run . --output=t1.txt hello --color=red hello shadow

go run . --output=t2.txt --color=red hello shadow

go run . --output=t3.txt --color=red

go run . --output=t4.txt helle --color=red 

go run . --color=red --output=t5.txt hello

go run . --color=red --output=t6.txt

go run . --color=red

go run . --color=red shadow

go run . --color=red hii

go run . --color=red --color=blue hello --output=t7.txt hii

go run . --color=red h  hello --output=t8.txt hello

go run . --output=t9.txt shadow

go run . --output=t10.txt --color=red --output=t11.txt lol

go run .

go run . lool

go run . standard

go run . hello standard

go run . --output=t12.txt hello

go run . --output=t13.txt hello shadow

go run . --output=t14.txt shadow

go run . --output=t15.txt --output=t16.txt --output=t17.txt hello

go run . --output=t18.txt --output=t19.txt --output=t20.txt hello shadow

  go run . --color=red h --color=green o --color=blue l --color=orange e hello

  go run . --color=red h --color=green o --color=blue l --color=orange e hello shadow 

  go run . --color=red h --color=green o --color=blue l --color=orange e hello shadow shadow
  
  go run . --output=rrr.txt --color=red h --color=green o --color=blue l --color=orange e hello shadow

  go run . --output=rrr.txt --color=red h --color=green o --color=blue l --output=rr.txt --color=orange e hello shadow

  go run . --output=rrr.txt --color=red h --color=green o --color=blue l --output=rr.txt --color=orange e hello --color=red h shadow

  go run . --color=red "a\n\ny" --color=blue oub "a\n\nyoub"
  
go run . --color=red "a\n\n\ny" --color=blue oub "a\n\n\nyoub"

go run . --color=red "a\n\n\ny" --color=blue oub "a\n\nyoub"

go run . --color=red "ay" --color=blue oub "a\n\nyoub"

go run . --color=red "a\n\n\ny" --color=blue oub "a\n\nyoub"

go run . --color=red "a\n\n\ny" --color=blue oub "a\n\n\nyoub"

go run . --color=red b --color=blue b "a\n\n\nyoub"

go run . --color=red b --color=green ba --color=blue b "a\n\n\nyouba"

go run . --color=red b --color=green ba --color=blue y "a\n\n\nyouba"

go run . --color=red b --color=green ba --color=blue y "a\n\n\nyouba" f

go run . --color=red b --color=green ba --color=blue y "a\n\n\nyouba" shadow.txt 

go run . --color=red b --color=green ba --color=blue y "a\n\n\nyouba" shadow

go run . --output=ttt.txt "heélo\ntà\ngggg" 

go run . --color=red --output=tx.txt hello shadow

go run . --color=red h helo  --output=tx.txt hello shadow

go run . --output=tx.txt  "helléo éh ehé" standard 

go run . --output=tx.txt  "helléo éh ehé" graffiti

go run . --output=tx.txt  "helléo éh ehé" varsity

go run . --output=tx.txt  "helléo éh\nehé" varsity lol lol lol lol