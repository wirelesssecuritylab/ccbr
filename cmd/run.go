package cmd

import (
	"ccbr/log"
	"ccbr/run/routers"
	"ccbr/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"net/http"
	"reflect"
)

var server string
var port string
var config string

func init() {
	rootCmd.AddCommand(policymanageServerCmd)
	policymanageServerCmd.Flags().StringVarP(&server, "server", "s", "127.0.0.1", "policy manager server ip address")
	policymanageServerCmd.Flags().StringVarP(&port, "port", "p", "9876", "policy manager server port")

}

type result struct {
	KeyName         string
	PropertiesValue string
}

var policymanageServerCmd = &cobra.Command{
	Use:   "run",
	Short: "run policy manage server",
	Long:  "run policy manage server",
	Run: func(cmd *cobra.Command, args []string) {
		config = utils.ParserArgs(config)
		router := gin.Default()
		router.Use(log.LoggerToFile())
		router = routers.CollectRoute(router)
		//router.LoadHTMLFiles()

		router.StaticFS("/static", http.Dir("./web/static"))
		router.LoadHTMLGlob("./web/html/*")

		addr := server + ":" + port
		router.Run(addr)

	},
}

func searchKeyValue(input map[string]interface{}, html string) string {
	for k, v := range input {
		if reflect.ValueOf(v).Kind() == reflect.Map {
			html += "<li>" + k + ": " + searchKeyValue(v.(map[string]interface{}), html) + "</li>"
		} else if reflect.ValueOf(v).Kind() == reflect.Array {

		} else if reflect.ValueOf(v).Kind() == reflect.String {
			html += "<li>" + k + ": " + v.(string) + "</li>"

		} else if reflect.ValueOf(v).Kind() == reflect.Slice {
			var s = "["
			for _, val := range v.([]interface{}) {
				s += val.(string) + ","
			}
			temp := s[0:len(s)-1] + "]"
			html += "<li>" + k + ": " + temp + "</li>"
		}
	}
	return html
}
