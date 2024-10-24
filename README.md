# NovaDriftCustomSkinRepository
Based off of [Voices of the Printer](https://github.com/madrod228/voicesoftheprinter) in spirit. 

Go to [Page1](pages/Page1.md) to see the start of this project. 

[main.go](goPageMaker/main.go) is where the main program lies, it goes through every directory in assets and uses them as the basis to create
[assets](goPageMaker/assets.csv) is a csv of all assets in the assets folder
[](goPageMaker/readDirectory.go) is where the functions that read the assets folder lie


Misc Links:
- [reading files in a directory](https://pkg.go.dev/os#ReadDir)
    - [continued](https://stackoverflow.com/questions/14668850/list-directory-in-go)
- [fStrings](https://stackoverflow.com/questions/11123865/format-a-go-string-without-printing#11124241)
- [slices](https://go.dev/ref/spec#Slice_types)
- [fileIO](https://pkg.go.dev/os)
- [file writing](https://gosamples.dev/write-file/)
- [making a go program](https://go.dev/doc/tutorial/getting-started)
