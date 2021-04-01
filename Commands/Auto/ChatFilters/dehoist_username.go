{{/*
    Removes "!" from the start of a username.

    Trigger: Regex
    Regex: \A

    Usage: Auto

    Made by: Victor Wolff#147
*/}}

{{$name := .Member.User.Username}}
{{if .Member.Nick}}
	{{$name = .Member.Nick}}
{{end}}
{{$test := reFind `(^!)` $name}}
{{if $test}}
	{{$newName := reReplace `(^!)` $name ""}}
	{{editNickname $newName}}
	{{$embed := cembed 
	"author" (sdict "name" .Member.User.String "icon_url" (.User.AvatarURL "256"))
    "description" (joinStr " " .Member.User.Mention "**nickname de-hoisted**")
    "color" 3375061
	"fields" (cslice 
        (sdict "name" "Antes" "value" $name "inline" false) 
        (sdict "name" "Depois" "value" $newName "inline" false)
    )
	"footer" (sdict "text" (joinStr " " "ID:" .Member.User.ID)) 
    "timestamp" currentTime
	}}
	{{sendMessage 821018660976984084 $embed}} {{/*Channel to send logs*/}}
{{end}}