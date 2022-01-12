package Builder

type NotificationBuilder struct{
	title string
	subtitle string
	userId string
	time int64
}

func newNotificationBuilder() *NotificationBuilder{
	return &NotificationBuilder{}
}


func(nb *NotificationBuilder) setTitle(title string){
	nb.title = title
}


func(nb *NotificationBuilder) setSubTitle(subtitle string){
	nb.subtitle = subtitle
}

func(nb *NotificationBuilder) setuserId(userId string){
	nb.userId = userId
}

func(nb *NotificationBuilder) setTime(time int64){
	nb.time = time
}

func(nb *NotificationBuilder) newNotification() (*notification,error){
	//add logic for builder to pass errors
	return &notification{
		time: nb.time,
		title: nb.title,
		subtitle: nb.subtitle,
		userId: nb.userId,
	},nil
}
