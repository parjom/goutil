커버리지 프로파일 뽑는 명령
<pre>
go test -v -coverprofile cover.out
</pre>
커버리지 프로파일로 html 파일 생성
<pre>
go tool cover -html=cover.out -o cover.html
</pre>