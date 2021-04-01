{{$embed := cembed
    "thumbnail" (sdict "url" "https://i.imgur.com/jUn8cNO.png")
    "description" "Yeeeeyy! \nVou te liberar por agora senpai...\nN-não esqueça senpai, eu estou sempre de olho."
    "color" 4437377
	"footer" (sdict "text" (joinStr " " "Enviado para:" .Member.User.ID)) 
    "timestamp" currentTime
	}}
    
{{sendMessage nil (complexMessage
 "content" (print .User.Mention " foi liberado(a)!")
 "embed" $embed
)}}