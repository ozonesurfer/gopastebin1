<!DOCTYPE html>
<html>
<head>
<title>Go Pastebin</title>
<link rel="stylesheet" href="/static/themes/prism-funky.css" />
<script src="/static/prism.js"></script>
</head>
<body>
<p>{{.Paste.Title}}</p>
<pre><code class="language-{{.Prism}}">{{.Paste.Content}}</code></pre>
<p>{{.Paste.Language}}</p>
<a href="/">New Paste</a>
</body>
</html>
