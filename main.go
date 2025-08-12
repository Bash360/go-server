package main

import (
"rest-api/backend"
)




func main(){
backend.Run("localhost:"+backend.Server.Port)
 

}
