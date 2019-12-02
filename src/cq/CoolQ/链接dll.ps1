$dll="E:\Git\Code\golang\src\cq\core\dev\me.cqp.Air.demo\app.dll"
$json="E:\Git\Code\golang\src\cq\core\dev\me.cqp.Air.demo\app.json"

# New-Item -ItemType SymbolicLink -Path "E:\Git\Code\golang\src\cq\core\dev\me.cqp.Air.demo\app.dll" -Target app.dll
New-Item -ItemType HardLink -Path $dll -Target app.dll
New-Item -ItemType HardLink -Path $json -Target app.json

