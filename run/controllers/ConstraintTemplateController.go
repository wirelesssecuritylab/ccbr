package controllers

import (
	"reflect"
)

type result struct {
	KeyName         string
	PropertiesValue string
} /*
func ListConstraintTemplate(ctx *gin.Context) {
	dynamicClient, err :=utils.GetDynamicClient("/root/.kube/config")
	if err!=nil {
		log.Fatal(err)
	}
	constrainttemplates,err := utils.ListConstraintTemplate(dynamicClient)

	var results []result

	for i:=0;i<len(constrainttemplates);i++ {
		//.Object.Spec.Crd.Spec.Names.Kind
		var temp result
		temp.KeyName = constrainttemplates[i].Object.Spec.Crd.Spec.Names.Kind
		if !utils.IsNil( constrainttemplates[i].Object.Spec.Crd.Spec.Validation.OpenAPIV3Schema.Properties) {
			bs_, _ := json.Marshal(constrainttemplates[i].Object.Spec.Crd.Spec.Validation.OpenAPIV3Schema.Properties)
			map2 := make(map[string]interface{})
			err = json.Unmarshal(bs_, &map2)
			if err != nil {
				fmt.Println(err)
			}
			output := searchKeyValue(map2,"")
			temp.PropertiesValue = "<ul style=\"padding-left:2em\">" + output + "</ul>"
		}else {
			temp.PropertiesValue = ""
		}
		results = append(results,temp)
	}

	//response.Success(ctx, gin.H{"constrainttemplates": constrainttemplates}, "obtain_successful")
	ctx.HTML(http.StatusOK,"constrainttemplates.html",gin.H{"constrainttemplates":constrainttemplates,"results":results})
}


*/

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
