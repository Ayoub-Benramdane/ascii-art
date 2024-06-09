go run . "aanana" standard aac

go run . "hello" standard | cat -e

go run . "hello world" shadow | cat -e

go run . "nice 2 meet you" thinkertoy | cat -e

go run . "you & me" standard | cat -e

go run . "123" shadow | cat -e

go run . "/(\")" thinkertoy | cat -e

go run . "aaCDEFGHIJKLMNOPQRSTUVWXYZ" shadow | cat -e

go run . "\"#$%&/()*+,-./" thinkertoy | cat -e

go run . "It's Working" thinkertoy | cat -e