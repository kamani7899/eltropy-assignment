## curl samples


curl -vk -X POST -d'{"username":"chandra","password":"Vemulapadu"}'  'http://localhost:8080/SignUp'

curl -vk -X POST -d'{"username":"chandra","password":"Vemulapadu"}'  'http://localhost:8080/signin'

curl -vk -X POST -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MTE3NjY0MTgsInVzZXJuYW1lIjoiY2hhbmRyYSJ9.ZIiQN8gv02UV7e_9e9KnFOvXlTjjlH1VIuQhhNduO6k" -d'{"name":"Lakshmi","Address":"Vemulapadu","contact":"sdfds","Kyc":"knasd8usa","Active":true,"Age":56}'  'http://localhost:8080/customers'


curl -vk -X POST -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MTE3NjY0MTgsInVzZXJuYW1lIjoiY2hhbmRyYSJ9.ZIiQN8gv02UV7e_9e9KnFOvXlTjjlH1VIuQhhNduO6k" -d'{"accounttype":"Savings","currentbalance":5000,"status":false,"details":"sdasdas","customerid":1}'  'http://localhost:8080/accounts'


curl -vk -X POST -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MTE3NjY0MTgsInVzZXJuYW1lIjoiY2hhbmRyYSJ9.ZIiQN8gv02UV7e_9e9KnFOvXlTjjlH1VIuQhhNduO6k" -d'{"accountid":1,"amount":5000,"employeeid":1,"transactiontype":"debit","customerid":1}'  'http://localhost:8080/transactions'


curl -vk -X POST -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MTE3NjY0MTgsInVzZXJuYW1lIjoiY2hhbmRyYSJ9.ZIiQN8gv02UV7e_9e9KnFOvXlTjjlH1VIuQhhNduO6k" -d'{"accountid":1, "name":"chandra","age":23,"active":false,"details":"sdasdas","address":"wewe","contact":"32wff","salary":1234}'  'http://localhost:8080/employees'





curl -vk -X POST -H "Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE23NjY0MTgsInVzZXJuYW1lIjoiY2hhbmRyYSJ9.ZIiQN8gv02UV7e_9e9KnFOvXlTjjlH1VIuQhhNduO6k"   'http://localhost:8080/signout'



# Useful commands to use the redis server
## install redis 
 brew install redis
## start redis
brew services start redis
## stop redis
brew services stop redis
## install redis cli
npm install -g redis-cli                  
## connect to redis using redis cli
rdcli



## Project Structure
	
	models corresponding to databse table are in models folder

	Business logic for handling each table is written in controllers folder

	func main in main.go is the starting point for execution

	Relied on https://github.com/go-gorm/gorm for orm purposes

	Refer to tables.sql for databse script

	 Relied on  Redis for storing tokens

	Used [gin routing server](https://github.com/gin-gonic/gin) for routing purpose
 
