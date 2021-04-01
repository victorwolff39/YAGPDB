{{$code := randInt 6}}

{{if eq $code 2}}
	{{if $db := dbGet .Channel.ID "stickymessage"}}
		{{deleteMessage nil (toInt $db.Value) 0}}
	{{end}}

	{{$string := "Iremos resetar a Mudae no dia **01/04 às 00:00**.\n **O valor dos personagens e badges serão reembolsados.**\n\n```$fullresetalias\n$fullresetalias2\n$fullresetimg\n$fullresetnoteimg\n$fullresetwishes\n$fullresetdisable\n$bitesthedust requiem\n$kakerarefundall```\n\nA partir do próximo reset, será obrigatório ter o cargo <@&826089209948274692> para jogar a Mudae. O cargo pode ser pego no <#818619192541577246>.\n\nQualquer dúvida entre em contato com a Administração."}}
	{{$message := cembed 
	"description" $string 
	"author" (sdict "name" "Atenção jogadores!" "icon_url" "https://cdn.discordapp.com/avatars/432610292342587392/29cb28fbf65a3958105026ab03abd306.png?size=256")
	"color" 9184224
	"footer" (sdict "text" "Discord Oficial - South America Memes" "icon_url" "https://i.imgur.com/vcEoJma.png")}}

	{{$id := sendMessageRetID nil $message}}
	{{dbSet .Channel.ID "stickymessage" (str $id)}}
{{end}}