Setup folders
---
- `shared`
- `server`
- `client`

Open folders in vscode as root:

1. open command palette (`F1` or `shift + ctrl + P`)
2. ```
   Add Folder to Workspace
   ```
3. Select all top folders

Compile proto (just use a .bat from an earlier project or this one)
---
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <myfile.proto>
```
```
go mod tidy 
```

Creating a module
---
```
cd <module folder>
```

```
go mod init <module name (must match folder path)>
```

After including new packages/modules
---
```
go mod tidy
```

Add in `go.mod` for non-shared packages
---
```GO
(...)
replace example.com/grpcdemo/shared => ../shared
(...)
require (
	(...)
	example.com/grpcdemo/shared v0.0.0
	(...)
)
```