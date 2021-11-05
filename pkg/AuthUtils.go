package pkg

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func GenerateUUID() string {
	var OriginalUUID, err = uuid.NewUUID()
	if err != nil {
		savedUUId := strings.Replace(OriginalUUID.String(),"-","",-1);
		return savedUUId
	}
	fmt.Println("Error thrown when creating UUID")
	return ""
}
