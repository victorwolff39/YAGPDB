{{ $colorCode := randInt 3 }}
{{ $string := "" }}
{{ if eq $colorCode 0 }}
    {{ $string =  "```fix\n"}}
{{ else if eq $colorCode 1 }}
    {{ $string =  "```yaml\n"}}
{{ else if eq $colorCode 2 }}
    {{ $string =  "```diff\n+ "}}
{{ end }}

{{$embed := cembed 
    "title" "Um novo membro entrou no servidor!"
    "description" ( joinStr "" $string "Um " ( joinStr " " .User.Username "selvagem apareceu.```" ))
    "color" 65453
    "thumbnail" (sdict "url" (.User.AvatarURL "256"))
}}

{{sendMessage nil (complexMessage
 "content" .User.Mention
 "embed" $embed
)}}