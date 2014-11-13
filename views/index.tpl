<html>
<head>
<title>Go Pastebin</title>
</head>
<body>

<header>
<h2>Create a Pastebin</h2>
</header>

<main>
<form action="/new" method="post">
<p><b>Title</b>
<input type="text" name="title" required /></p>
<p><b>Content</b><br />
<textarea name="content" required></textarea></p>
<b>Languages</b><br />
<select name="language" required>
	{{range $k, $v := .Languages}}
	<option value={{$k}}>{{$k}}</option>
	{{end}}
</select>
<p><input type="submit" value="Create Pastebin" /></p>
</form>
</main>

<aside>
<ul>
{{range .Pastes}}
<li><a href="/paste/{{.Id}}">{{.Title}}</a></li>
{{end}}
</ul>
</aside>

</body>
</html>