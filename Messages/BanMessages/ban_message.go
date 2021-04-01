{{- $silent := ($desc :=  (joinStr " " "**ID:**" (.User.ID))) -}}
{{- $silent = ($desc =  (joinStr " " $desc "\n")) -}}
{{ if .Reason }}
    {{- $silent = ($desc =  (joinStr " " (joinStr " " $desc "**Motivo:**") (.Reason))) -}}
    {{- $silent = ($desc =  (joinStr " " $desc "\n")) -}}
{{ end }}
{{ if .Duration }}
    {{- $silent = ($desc =  (joinStr " " (joinStr " " $desc "**Duração:**") (.HumanDuration))) -}}
    {{- $silent = ($desc =  (joinStr " " $desc "\n")) -}}
{{ else }}
    {{- $silent = ($desc =  (joinStr " " (joinStr " " $desc "**Duração:**" "Permanente"))) -}}
    {{- $silent = ($desc =  (joinStr " " $desc "\n")) -}}
{{ end }}
{{ $desc = (joinStr "" $desc "\n\nSINTA O PODER DO MARTELO!! <a:VM_ban:763967230428839949>") }}

{{ $moderator := (getMember .Author.ID) }}

{{if $moderator.Nick}}
    {{$modname := $moderator.Nick}}
{{else}}
    {{$modname := $moderator.User.Username}}
{{end}}

{{ $title := (joinStr " " (joinStr " " (joinStr "" ($modname)) "livrou o servidor de") (.User.Username)) }}

{{$embed := cembed
    "title" $title
    "description" $desc
    "color" 13632027
    "thumbnail" (sdict "url" ($moderator.User.AvatarURL "256"))  
    "image" (sdict "url" "https://i.imgur.com/j2pwQop.gif")
    "footer" (sdict "text" "Att South America Memes (✿◠‿◠)" "icon_url" "https://i.imgur.com/vcEoJma.png")
}}

{{ sendMessage  761034261733179412 $embed}}
