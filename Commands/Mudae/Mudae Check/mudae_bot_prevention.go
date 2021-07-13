{{/*
    A command to check a player's activity to help detect selfBots.
	I originally made for the bot Mudae, but it is possible to change it for your own use.

    Trigger: Regex
    Regex: \A

    Usage: Auto

    Made by: Victor Wolff#147
*/}}

{{/*---- Configuration ---- */}}
{{$timer := 120}}
{{$triggerChance := 8}}
{{$successLog := 821018660976984084}}
{{$failedLog := 821018660976984084}}
{{$startedLog := 821018660976984084}}

{{/* Initiating variables */}}
{{$trigger := randInt $triggerChance}}
{{$attemptCount := 0}}
{{$securityCode := 0}}
{{$strikes := 0}}
{{$messageID := ""}}

{{/* Getting database values */}}
{{$isWaitingResponse := (dbGet .User.ID "mudaecIsWaitingResponse").Value}}
{{$strikes = (toInt (dbGet .User.ID "mudaecStrikes").Value)}}
{{if $isWaitingResponse}} {{/* No need to get runCount and securityCode if the bot is not waiting for a response */}}
	{{$attemptCount = (toInt (dbGet .User.ID "mudaecAttemptCount").Value)}}
	{{$securityCode = (toInt (dbGet .User.ID "mudaecSecurityCode").Value)}}
{{end}}

{{/* If command was NOT executed by scheculeUniqueCC (triggered by a user) */}}
{{if not .ExecData}}
	{{/* Execute the command randomly - First time the command is executed */}}
	{{if or (and (eq $trigger 0) (not $isWaitingResponse)) (and (eq .User.ID 121709907681476610) (eq .Message.Content "--trigger"))}}
		{{/* Generates the 4-pin code */}}
		{{$generatedCode := randInt 1000 10000}}

		{{/* Create the embed with the generated security code, sends it and return the message ID */}}
		{{$embed := cembed
			"image" (sdict "url" "https://i.imgur.com/U8UpPNI.jpg")
			"description" (print "Co-com licença senpai... \n\nSou a FBI-chan... Preciso que você me diga se está jogando mesmo... \n Tem gente que usa bot... Isso me deixa triste. \n\n Digite o código no chat pra provar que você não é um robô. \n\n**Código:** ||" $generatedCode "||")
			"color" 3375061
			"footer" (sdict "text" (joinStr " " "Enviado para:" .Member.User.ID)) 
			"timestamp" currentTime
		}}
		{{$messageID = sendMessageRetID nil (complexMessage
		"content" (joinStr " " .User.Mention " foi enquadrado(a)!")
		"embed" $embed
		)}}

		{{/* Save values to the database and schedules the timeout command */}}
		{{dbSetExpire .User.ID "mudaecSecurityCode" $generatedCode $timer}}
		{{dbSetExpire .User.ID "mudaecIsWaitingResponse" true $timer}}
		{{scheduleUniqueCC .CCID nil $timer (print "mudaecTimeoutResponse" .User.ID) 1}}

		{{/* Create the log embed and sends it to the specified channel */}}
		{{$messageLink := (print "https://discord.com/channels/" .Guild.ID "/" .Channel.ID "/" $messageID)}}
		{{$embed := cembed 
			"author" (sdict "name" .Member.User.String "icon_url" (.User.AvatarURL "256"))
			"description" (print "Verificação de atividade iniciada para " .User.Username ". [Link](" $messageLink ")")
			"color" 16312092
			"footer" (sdict "text" (joinStr " " "ID:" .Member.User.ID)) 
			"timestamp" currentTime
		}}
		{{sendMessage $startedLog $embed}}
	{{/* Execute when the command is waiting response */}}
	{{else if $isWaitingResponse}}
		{{/* Add +1 attempt */}}
		{{$attemptCount = (add $attemptCount 1)}}

		{{if eq (toInt .Message.Content) $securityCode}}
			{{/* Create success message embed, sends it and return the message's ID */}}
			{{$embed := cembed
				"thumbnail" (sdict "url" "https://i.imgur.com/jUn8cNO.png")
				"description" "Yeeeeyy! \nVou te liberar por agora senpai...\nN-não esqueça senpai, eu estou sempre de olho."
				"color" 4437377
				"footer" (sdict "text" (joinStr " " "Enviado para:" .Member.User.ID)) 
				"timestamp" currentTime
			}}
			{{$messageID := sendMessageRetID nil (complexMessage
			"content" (print .User.Mention " foi liberado(a)!")
			"embed" $embed
			)}}
			
			{{/* Cancel scheduled timeout CC */}}
			{{cancelScheduledUniqueCC .CCID (print "mudaecTimeoutResponse" .User.ID)}}

			{{/* Clean database entries */}}
			{{dbDel .User.ID "mudaecSecurityCode"}}
			{{dbDel .User.ID "mudaecIsWaitingResponse"}}
			{{dbDel .User.ID "mudaecAttemptCount"}}
			
			{{/* Create the log embed and sends it to the specified channel */}}
			{{$messageLink := (print "https://discord.com/channels/" .Guild.ID "/" .Channel.ID "/" $messageID)}}
			{{$embed := cembed 
				"author" (sdict "name" .Member.User.String "icon_url" (.User.AvatarURL "256"))
				"description" (print "Verificado com sucesso. [Link](" $messageLink ").\nTentativas: `" $attemptCount "`")
				"color" 4437377
				"footer" (sdict "text" (joinStr " " "ID:" .Member.User.ID)) 
				"timestamp" currentTime
			}}
			{{sendMessage $successLog $embed}}
		{{else if lt $attemptCount 5}}
			{{/* Create the embed with a random error message/image and sends it */}}
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
				"title" (print "Tentativa " (add $attemptCount 1) "/5")
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

			{{/* Record attempt in the database with expiration time */}}
			{{dbSetExpire .User.ID "mudaecAttemptCount" $attemptCount $timer}}
		{{else}}
			{{/* Create the embed with the failed message, sends it and return the message ID */}}
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
			{{$messageID = sendMessageRetID nil (complexMessage
				"content" (print "Você falhou na verificação " .User.Mention)
				"embed" $embed
			)}}
			
			{{/* Cancel execution of the timeout CC */}}
			{{cancelScheduledUniqueCC .CCID (print "mudaecTimeoutResponse" .User.ID)}}

			{{/* Removes all verification database entries */}}
			{{dbDel .User.ID "mudaecSecurityCode"}}
			{{dbDel .User.ID "mudaecIsWaitingResponse"}}
			{{dbDel .User.ID "mudaecAttemptCount"}}
			{{dbSet .User.ID "mudaecStrikes" (add $strikes 1)}}

			{{/* Create the log embed and sends it to the specified channel */}}
			{{$messageLink := (print "https://discord.com/channels/" .Guild.ID "/" .Channel.ID "/" $messageID)}}
			{{$embed := cembed 
				"author" (sdict "name" .Member.User.String "icon_url" (.User.AvatarURL "256"))
				"description" (print "Falhou a verificação. [Link](" $messageLink ") \nTentativas: `" $attemptCount "`")
				"color" 14763791
				"footer" (sdict "text" (joinStr " " "ID:" .Member.User.ID)) 
				"timestamp" currentTime
			}}
			{{sendMessage $failedLog $embed}}
		{{end}}
	{{end}}

{{/* If command was executed by scheculeUniqueCC (timeout) */}}
{{else}}
	{{/* Create the timeout embed and sends it */}}
	{{$embed := cembed
    "thumbnail" (sdict "url" "https://i.imgur.com/zZniOmM.png")
    "description" "Ãn? \n\nSe-sennpai me deixou no vácuo... \nVou lembrar disso eim!"
    "color" 13632027
	"footer" (sdict "text" (joinStr " " "Enviado para:" .Member.User.ID)) 
    "timestamp" currentTime
	}}

	{{sendMessage nil (complexMessage
	"content" .User.Mention
	"embed" $embed
	)}}
	{{/* Delete all the user's attempt */}}
	{{dbDel .User.ID "mudaecAttemptCount"}}
	{{dbSet .User.ID "mudaecStrikes" (add $strikes 1)}}

	{{/* Create the log embed and sends it to the specified channel */}}
	{{$messageLink := (print "https://discord.com/channels/" .Guild.ID "/" .Channel.ID "/" $messageID)}}
	{{$embed := cembed 
		"author" (sdict "name" .Member.User.String "icon_url" (.User.AvatarURL "256"))
		"description" (print "Falhou a verificação. [Link](" $messageLink ") \n:alarm_clock: **TIMEOUT**")
		"color" 14763791
		"footer" (sdict "text" (joinStr " " "ID:" .Member.User.ID)) 
		"timestamp" currentTime
	}}
	{{sendMessage $failedLog $embed}}
{{end}}
