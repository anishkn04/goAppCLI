package methods

import (
	"fmt"
	"os"
)

func CheckEnv() {
	_, err := os.Stat(".env");
	if err.Error() == "CreateFile .env: The system cannot find the file specified."{
		var botToken string;
		fmt.Println("Enter Discord Bot Token: ");
		fmt.Scanln(&botToken)
		var channelId string;
		fmt.Println("Enter the Channel Id where you want to post changes: ");
		fmt.Scanln(&channelId)
		var waitTime string;
		fmt.Println("Enter the interval (in seconds) in which you want to check for updates: ");
		fmt.Scanln(&waitTime)
		var pageID string;
		fmt.Println("Enter Facebook Page ID: ");
		fmt.Scanln(&pageID)
		var pageAccessToken string;
		fmt.Println("Enter Facebook Page Access Token: ");
		fmt.Scanln(&pageAccessToken)
		var eventSite string;
		fmt.Println("Enter the site or file name (relative) with the events data: ");
		fmt.Scanln(&eventSite)

		finalString := fmt.Sprintf("BOTTOKEN='%s'\nCHANNELID='%s'\nWAITTIME='%s'\nPAGE_ID='%s'\nPAGE_ACCESS_TOKEN='%s'\nSITE='%s'", botToken, channelId, waitTime, pageID, pageAccessToken, eventSite)
		err = os.WriteFile(".env", []byte(finalString), 0777)
		Check(&err);
	}
	Check(&err);
}