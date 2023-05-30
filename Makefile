dago-mac:
	echo "module ${module_name}\n\ngo ${go_version}\n\nrequire (\n\tgithub.com/gin-gonic/gin v1.8.1\n\tgithub.com/go-ozzo/ozzo-validation v3.6.0+incompatible\n\tgithub.com/joho/godotenv v1.4.0\n\tgorm.io/driver/mysql v1.3.4\n\tgorm.io/gorm v1.23.7\n)" >> go.mod
	go get
