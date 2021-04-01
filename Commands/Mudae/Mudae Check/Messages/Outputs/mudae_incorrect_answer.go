{{/* ---- DEBUG CODE ---- */}} {{$securityCode := 1234}} {{/* ---- DEBUG CODE ---- */}}
{{/* ---- DEBUG CODE ---- */}} {{$runCount := 2}} {{/* ---- DEBUG CODE ---- */}}

{{$string := ""}}
{{$url := ""}}
{{$choice := (randInt 3)}}

{{if eq $choice 0}}
    {{$string = "Hã? \nPelos meus cálculos, sua resposta está errada!"}}
    {{$url = "https://i.imgur.com/v6feiM8.png"}}
{{else if eq $choice 1}}
    {{$string = "BAKA! \nComo assim senpai? Esse código não faz sentido!"}}
    {{$url = "https://i.imgur.com/z8C1VPt.png"}}
{{else if eq $choice 2}}
    {{$string = "Ara ara... \nSenpai não sabe escrever! hahaha"}}
    {{$url = "https://i.imgur.com/vMEBN7d.png"}}
{{end}}

{{$embed := cembed
    "title" (print "Tentativa " $runCount "/5")
    "thumbnail" (sdict "url" $url)
    "description" (print $string "\n\n**Código:** ||" $securityCode "||")
    "color" 16312092
	"footer" (sdict "text" (joinStr " " "Enviado para:" .Member.User.ID)) 
    "timestamp" currentTime
	}}

{{sendMessage nil (complexMessage
 "content" .User.Mention
 "embed" $embed
)}}