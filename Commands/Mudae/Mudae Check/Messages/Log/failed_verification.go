{{/* ---- DEBUG CODE ---- */}} {{$attemptCount := 3}} {{/* ---- DEBUG CODE ---- */}}
{{/* ---- DEBUG CODE ---- */}} {{$failedLog := 821018660976984084}} {{/* ---- DEBUG CODE ---- */}}
{{/* ---- DEBUG CODE ---- */}} {{$messageID := 826929121295859733}} {{/* ---- DEBUG CODE ---- */}}

{{$messageLink := (print "https://discord.com/channels/" .Guild.ID "/" .Channel.ID "/" $messageID)}}

{{$embed := cembed 
"author" (sdict "name" .Member.User.String "icon_url" (.User.AvatarURL "256"))
"description" (print "Falhou a verificação. \nTentativas: `" $attemptCount "`\n[Link da mensagem](" $messageLink ")")
"color" 14763791
"footer" (sdict "text" (joinStr " " "ID:" .Member.User.ID)) 
"timestamp" currentTime
}}

{{sendMessage $failedLog $embed}}