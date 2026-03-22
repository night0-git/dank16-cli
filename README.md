## This is a standalone binary of the dank16 color generator from DankMaterialShell.

### Usage:
- Build with:
```
cd dank16-cli
go build -o dank16
```
- Run:
```
./dank16 --color <base hex color>
Options:
  --color string
    	Primary hex color (Required)
  --light
    	Generate a light mode palette instead
  --no-dps
    	Disable Delta Phi Star contrast
  --bg string
    	Custom background hex
  --string
    	Output as string list instead of JSON
```
