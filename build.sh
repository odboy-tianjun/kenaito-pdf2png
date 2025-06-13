rm -f kenaito-pdf2png.exe
rm -f kenaito-pdf2png-release.exe
# build
go build -o kenaito-pdf2png-release.exe main.go
# upx compress
./upx -o kenaito-pdf2png.exe kenaito-pdf2png-release.exe
rm -f kenaito-pdf2png-release.exe
