package main

import (
	"encoding/json"
	"fmt"
)

func injectMultiTitles(meta map[string]interface{}) error {
	meta["codeinjection_head"] = ""

	_, ok := meta["title"].(string)
	if !ok {
		titleDataMap, ok := meta["title"].(map[string]interface{})
		if !ok {
			log.Error("missing default title")
			return nil
		}

		defaultTitle, ok := titleDataMap["default"].(string)
		if !ok {
			log.Error("missing default title in multi-title")
			return nil
		}

		meta["title"] = defaultTitle

		titleDataBytes, err := json.Marshal(titleDataMap)
		if err != nil {
			log.Error("error marshaling title data: %v", err)
			return err
		}

		titleDataStr := string(titleDataBytes)
		log.Debug("multi title: ", titleDataStr)

		meta["codeinjection_head"] = fmt.Sprintf(`<script>
			changetitle(%s);
			</script>`, titleDataStr)
	}

	return nil
}
