##curl samples


curl -vk -X POST -d'{"name":"Lakshmi","Address":"Vemulapadu","contact":"sdfds","Kyc":"knasd8usa","Active":true,"Age":56}'  'http://localhost:8080/customers'


curl -vk -X POST -d'{"accounttype":"Savings","currentbalance":5000,"status":false,"details":"sdasdas","customerid":1}'  'http://localhost:8080/accounts'


curl -vk -X POST -d'{"accountid":1,"amount":5000,"employeeid":1,"transactiontype":"debit","customerid":1}'  'http://localhost:8080/transactions'


curl -vk -X POST -d'{"accountid":1, "name":"chandra","age":23,"active":false,"details":"sdasdas","address":"wewe","contact":"32wff","salary":1234}'  'http://localhost:8080/employees'


