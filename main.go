package main

import (

)

func main() {
	urls := fmt.Sprintf("https://restapi.amap.com/v5/ip?key=e93f35e859dfc4efc1f2bae4ac222cf2&type=4&ip=%s", "39.144.104.249")
	response, _ := http.Get(urls)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	fmt.Println(string(body))
	hello()
}

