��#   l a b 
 
go mod init
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get github.com/asaskevich/govalidator
go get github.com/onsi/gomega@v1.19

git checkout -b issue-9
git add .
git commit -m “ทำ Entity Playlist - close #9”
git push origin main


ทำการ merge = main
git merge issue-9 --no-ff
