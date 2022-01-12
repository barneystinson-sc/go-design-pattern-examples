package Builder

import (
	"fmt"
	"time"
)

func BuilderExample(){
	notificationBuilder := newNotificationBuilder()
	notificationBuilder.setTitle("new title")
	notificationBuilder.setSubTitle("new subtitle")
	notificationBuilder.setTime(time.Now().UnixMilli())
	notificationBuilder.setuserId("user_1")

	newNotification,_ := notificationBuilder.newNotification()

	fmt.Println("Notification obj %+v",newNotification)
}
