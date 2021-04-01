{{$embed := cembed
    "image" (sdict "url" "https://i.imgur.com/U8UpPNI.jpg")
    "description" "Co-com licença senpai... \n\nSou a FBI-chan... Preciso que você me diga se está jogando mesmo... \n Tem gente que usa bot... Isso me deixa triste. \n\n Digite o código no chat pra provar que você não é um robô. \n\n**Código:** ||0147|| \n\nBom jogo!"
    "color" 3375061
	"footer" (sdict "text" (joinStr " " "Enviado para:" .Member.User.ID)) 
    "timestamp" currentTime
	}}

{{sendMessage nil (complexMessage
 "content" (joinStr " " .User.Mention " foi enquadrado(a)!")
 "embed" $embed
)}}