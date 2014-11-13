gopastebin1
===========

A golang version of Pastebin with syntax highlighting, driven by MongoDB

# Dependencies

<code>go get labix.org/v2/mgo</code>

You might need Mercurial or Bazaar

<code>go get github.com/astaxie/beego</code>

<code>go get github.com/beego/bee</code>

Copy the bee (or bee.exe) program to your PATH.

# Credits
This repository includes the Prism syntax highlighting system... from <b>github.com/LeaVerou/prism</b> and
<b>prismjs.com</b>.

#Configuration and Running
You will need to install MongoDB. You might want to change the database name in <b>src/routers/router.go</b> 
and the MongoDB server URL in <b>src/model/repository.go</b>. Add the project to GOPATH. The port is in <b>app.conf</b>.  

Move to the root of this repository and enter
<code>bee run</code>
