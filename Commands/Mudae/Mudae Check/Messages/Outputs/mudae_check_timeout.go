{{$embed := cembed
    "thumbnail" (sdict "url" "https://i.imgur.com/zZniOmM.png")
    "description" "Ãn? \n\nSe-sennpai me deixou no vácuo... \nVou lembrar disso eim!"
    "color" 3375061
	"footer" (sdict "text" (joinStr " " "Enviado para:" .Member.User.ID)) 
    "timestamp" currentTime
	}}

{{sendMessage nil (complexMessage
 "content" .User.Mention
 "embed" $embed
)}}