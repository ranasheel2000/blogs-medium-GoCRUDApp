APIs

ToDo Items | 
--- |

UseCase | HttpMethod | route | RouteDetails | 
--- | --- | --- | --- | 
Create| POST | '/v1/todo/' | To create new ToDo item|
List all| GET | '/v1/todo/' | To list all ToDo items| 
List specific| GET | '/v1/todo/<id>' | To list all details of specific ToDo item| 
Delete| DELETE | '/v1/todo/<id>' | To delete specific ToDo item| 
Update| PUT | '/v1/todo/<id>' | To update specific ToDo item|
Update status of ToDo item| PUT | '/v1/todo/status/<id>' | To update status of specific ToDo item|

MoM Items | 
--- |

UseCase | HttpMethod | route | RouteDetails | 
--- | --- | --- | --- | 
Create| POST | '/v1/mom/' | To create new MoM item|
List all| GET | '/v1/mom/' | To list all MoM items| 
List specific| GET | '/v1/mom/<id>' | To list all details of specific MoM item| 
Delete| DELETE | '/v1/mom/<id>' | To delete specific MoM item| 
Update| PUT | '/v1/mom/<id>' | To update specific MoM item details|



##INIT
git clone git@github.com:ranasheel2000/blogs-medium-GoCRUDApp.git
go mod init github.com/ranasheel2000/blogs-medium-GoCRUDApp

go get -u github.com/spf13/viper
go get -u github.com/gorilla/mux
go get -u gorm.io/gorm
gorm.io/driver/postgres

mkdir configs #For App configurations.
mkdir deployments #For DevOps stuff.
mkdir docs #For project documents.
mkdir test #For unit tests.
mkdir cmd #Main directory for application code.
mkdir -p web/app #For frontend-javascript/react etc.
mkdir -p web/static #For static file like images etc.
mkdir -p web/template #For html, yaml files etc.

go get -u github.com/spf13/cobra@latest
go install github.com/spf13/cobra-cli@latest
cobra-cli init #It will create root.go inside cmd directory and main.go in current directory.






