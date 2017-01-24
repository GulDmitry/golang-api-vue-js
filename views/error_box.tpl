{{if .flash.notice}}
<div class='alert alert-success alert-dismissable fade in'>
    <a href='#' onClick="" class='close' data-dismiss='alert' aria-label='close'>&times;</a>
    {{str2html .flash.notice}}
</div>
{{end}}
{{if .flash.warning}}
<div class='alert alert-warning alert-dismissable fade in'>
    <a href='#' onClick="" class='close' data-dismiss='alert' aria-label='close'>&times;</a>
    {{str2html .flash.warning}}
</div>
{{end}}
{{if .flash.error}}
<div class='alert alert-danger alert-dismissable fade in'>
    <a href='#' onClick="" class='close' data-dismiss='alert' aria-label='close'>&times;</a>
    {{str2html .flash.error}}
</div>
{{end}}