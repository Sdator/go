# 设置go的工作目录
[environment]::SetEnvironmentvariable("GOPATH", $(get-location), "User")
# 设置编译目录bin
[environment]::SetEnvironmentvariable("GOBIN", "$(get-location)\bin", "User")