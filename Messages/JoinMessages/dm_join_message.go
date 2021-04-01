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
    "title" "Bem vindo(a) ao Discord oficial da South America Memes!"
    "description" `Ficamos felizes por vocês fazerem parte da tropa!
    
    Que tal dar uma olhada em alguns canais úteis no servidor primeiro?

    <#775417758435246140>: onde ficam as regras do servidor;
    <#758831162083901451>: aqui temos muitas informações úteis que podem te ajudar;
    <#758837183066275863>: aqui onde ficam os memes postados pela comunidade;
    <#791104364370460692>: selecione sobre o que deseja ser notificado no servidor, basta reagir nas mensagens;
    
    Se quiser fazer um registro para contar aos outros um pouco mais sobre você vá no <#763894998588194817>.
    
    Agora que você já leu tudo... Fique a vontade para conversar com todos no servidor no <#761034261733179412> ou entrar em qualquer chamada de voz!
    Se tiver qualquer dúvida, fique a vontade para perguntar no <#768030777627377685>, estaremos sempre felizes em ajudar! :smile:
    `
    "color" 66666
    "thumbnail" (sdict "url" "https://i.imgur.com/q63tkeg.png")
    "footer" (sdict "text" "Discord Oficial - South America Memes" "icon_url" "https://i.imgur.com/vcEoJma.png")
}}

{{sendMessage nil $embed}}