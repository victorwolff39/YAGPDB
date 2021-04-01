{{$string := ""}}
{{$url := ""}}
{{$choice := (randInt 2)}}

{{if eq $choice 0}}
    {{$string = "Ha! \nAcabou suas tentativas senpai! \nVocê será levado sob custódia por falhar a verificação!"}}
    {{$url = "https://i.imgur.com/uG7Njat.png"}}
{{else if eq $choice 1}}
    {{$string = "Owwn! Kawaii... \nEle não conseguiu completar a verificação! \nVou me lembrar disso senpai. uwu"}}
    {{$url = "https://i.imgur.com/Di8bYSM.png"}}
{{end}}

{{$embed := cembed
    "thumbnail" (sdict "url" $url)
    "description" $string
    "color" 13632027
	"footer" (sdict "text" (joinStr " " "Enviado para:" .Member.User.ID)) 
    "timestamp" currentTime
	}}

{{sendMessage nil (complexMessage
 "content" (print "Você falhou na verificação " .User.Mention)
 "embed" $embed
)}}