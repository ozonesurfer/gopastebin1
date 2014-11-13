<!DOCTYPE html>

<html>
<head>
<title>Go Pastebin</title>
</head>
<body>
<header>
<h3>Create a Pastebin</h3>
</header>

<main>

	<form action="/new" method="post">
		<p><b>Title</b><br />
			<input type="text" name="title" required /></p>
		<p><label>Content</label><br />
			<textarea name="content" required></textarea></p>
		<label>Languages</label>
		<select name="language" required>
			{{range $k, $v := .Languages}}
				<option value="{{$k}}">{{$k}}</option>
			{{end}}
		</select>
		<p><input type="submit" value="Create Pastebin" /></p>
	</form>

</main>

<aside>
<h3>Existing Pastes</h3>
<ul>
	{{range .Pastes}}
		<li><a href="/paste/{{.Id}}">{{.Title}}</a></li>
	{{end}}
</ul>
</aside>
</body>
</html>