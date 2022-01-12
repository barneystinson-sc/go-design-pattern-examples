package Singleton

func SingletonExample(){
	Logger := getInstance()
	Logger.Log("Hello World")

	Logger = getInstance()
	Logger.Log("Hello World 2")

	//try concurrency safety : uncomment above lines to see full effect

	for i:=0;i<10;i++{
		go getInstance()
	}
	//fmt.Scanln()
}
