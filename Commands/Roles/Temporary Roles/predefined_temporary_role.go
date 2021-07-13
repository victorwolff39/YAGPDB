{{/*
    Add two predefined roles to a user.
	The time is optional, a default and max value can be configured in the "Configuration" section.

    Trigger: Command (mention/cmd prefix)

    Usage: !{commandName} userId/userMention delay

    Made by: Victor Wolff#147
*/}}
 
{{/*---- Configuration ---- */}}
{{/* Roles to give (role ids): */}}{{$roleList := cslice 788927143023214612 782767457156071495}}
{{/* Default time in days: */}}{{$defaultTime := 14}}
{{/* Max time in days: */}}{{$maxTime := 14}}
{{/* Log channel ID: */}}{{$logChannel := 818619162300645416}}


{{/*---- Parse Arguments ---- */}}
{{$args := parseArgs 1 "Please inform a user."
	(carg "member" "User to add roles.")
	(carg "duration" "Optional duration.")}}

{{$defaultTime = (mult $defaultTime 86400)}}
{{$maxTime = (mult $maxTime 86400)}}

{{$target := $args.Get 0}}
{{$targetRoles := $target.Roles}}
{{$roleListLenght := len $roleList}}
{{$duration := (toDuration (print (toString $maxTime) "s"))}}

{{if $args.IsSet 1}}
	{{$duration = $args.Get 1}}	
{{end}}

{{if not (gt (toInt $duration.Seconds) (toInt $maxTime))}}
	{{$qtdRolesIgnored := 0}}
	{{$addedRolesString := ""}}
	{{range $roleList}}
		{{if not (in $targetRoles .)}}
			{{giveRoleID $target.User.ID .}}
			{{$addedRolesString = (print $addedRolesString " <@&" . ">")}}
		{{else}}
			{{$qtdRolesIgnored = (add $qtdRolesIgnored 1)}}
		{{end}}
		{{takeRoleID $target.User.ID . $duration.Seconds}}
	{{end}}
	{{if not (gt $qtdRolesIgnored 0)}}
		{{- print "Roles added to **" $target.User.Username "** for " $duration "." -}}
	{{else}}
		{{- print "Roles added to **" $target.User.Username "** for " $duration ". Ignored: " $qtdRolesIgnored -}}
	{{end}}

	{{if not (eq $qtdRolesIgnored $roleListLenght)}}
		{{$embed := cembed 
			"author" (sdict "name" .Member.User.String "icon_url" (.User.AvatarURL "256"))
			"description" (print "Added roles: " $addedRolesString " to " $target.User.Mention " for **" $duration "**.")
			"color" 3375061
			"footer" (sdict "text" (joinStr " " "ID:" .Member.User.ID)) 
			"timestamp" currentTime
			}}
		{{sendMessage $logChannel $embed}} {{/*Channel to send logs*/}}
	{{end}}
{{else}}
	{{- print " Duration can't be greater than " $duration -}}
{{end}}
