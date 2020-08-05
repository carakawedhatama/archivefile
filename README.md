# ArchiveFile

Archive a file to zip extension file, with additional password to secure your zip file. This feature is using [alexmullins'](https://github.com/alexmullins/zip) library of archiving file based on Go Programming language.

Instead of using `ioutil.ReadAll` due to memory overload (when you need to archive such a large file), I use `io.Copy` to process the file. Thanks to [Haisum's Blog](https://haisum.github.io/2017/09/11/golang-ioutil-readall/) for such a good advice.