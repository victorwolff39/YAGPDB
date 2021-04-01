{{/* ---- DEBUG CODE ---- */}} {{$attemptCount := 3}} {{/* ---- DEBUG CODE ---- */}}
{{/* ---- DEBUG CODE ---- */}} {{$successLog := 821018660976984084}} {{/* ---- DEBUG CODE ---- */}}
{{/* ---- DEBUG CODE ---- */}} {{$messageID := 826929121295859733}} {{/* ---- DEBUG CODE ---- */}}

{{$messageLink := (print "https://discord.com/channels/" .Guild.ID "/" .Channel.ID "/" $messageID)}}

{{$embed := cembed 
"author" (sdict "name" .Member.User.String "icon_url" (.User.AvatarURL "256"))
"description" (print "Verificado com sucesso. \nTentativas: `" $attemptCount "`\n[Link da mensagem](" $messageLink ")")
"color" 4437377
"footer" (sdict "text" (joinStr " " "ID:" .Member.User.ID)) 
"timestamp" currentTime
}}

{{sendMessage $successLog $embed}}