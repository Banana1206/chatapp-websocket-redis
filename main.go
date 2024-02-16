package main

import (
	// "chatapp-websoket-redis/model"
	// "chatapp-websoket-redis/model"
	"chatapp-websoket-redis/server/redisrepo"
	"fmt"
	// "log"
	// "time"
)

func main() {
	redisrepo.InitialiseRedis()

	// fmt.Println("============1. UpdateCOntactList Test=========")
	// fmt.Println()
	// fmt.Println()

	// for i := 14; i < 18; i++ {
	// 	redisrepo.UpdateContactList("user2", fmt.Sprintf("user%d", i))
	// 	time.Sleep(time.Second * 2)
	// }

	// fmt.Println("============2. FetchContactList Test=========")
	// fmt.Println()
	// fmt.Println()

	// res, err := redisrepo.FetchContactList("user2")
	// if err != nil {
	// 	fmt.Println("error in fetch", err)
	// 	return
	// }
	// fmt.Println(res)

// fmt.Println("============3. CreateChat Test=========")
	// fmt.Println()
	// fmt.Println()

	// // Số lượng cuộc trò chuyện cần tạo
	// numberOfChats := 5

	// // Sử dụng vòng lặp để tạo số lượng cuộc trò chuyện mong muốn
	// for i := 0; i < numberOfChats; i++ {
	// 	// Tạo một đối tượng Chat
	// 	chat := &model.Chat{
	// 		ID:        fmt.Sprintf("%d", i+1),
	// 		From:      fmt.Sprintf("user%d", 2), // user1 hoặc user2 lặp lại
	// 		To:        fmt.Sprintf("user%d", (i+1)%2+15), // user2 hoặc user1 lặp lại
	// 		Msg:       fmt.Sprintf("Hello from user%d to user%d", 2, (i+1)%2+15),
	// 		Timestamp: 1645225200 + int64(i), // Thời gian Unix tăng dần cho mỗi cuộc trò chuyện
	// 	}

	// 	// Thêm cuộc trò chuyện vào slice
	// 	// chats = append(chats, chat)

	// 	// Gọi hàm CreateChat
	// 	chatKey, err := redisrepo.CreateChat(chat)
	// 	if err != nil {
	// 		log.Fatalf("Lỗi khi tạo chat: %v", err)
	// 	}

	// 	fmt.Printf("Chat đã được tạo thành công. Key của chat là: %s\n", chatKey)
	// }


	// fmt.Println("============4. CreateFetchChatBetweenIndex Test=========")
	// fmt.Println()
	// fmt.Println()

	// redisrepo.CreateFetchChatBetweenIndex() 


	fmt.Println("============5. FetchChatBetween Test=========")
	fmt.Println()
	fmt.Println()
	res, err := redisrepo.FetchChatBetween("user15", "user2", "0", "+inf")

	if err != nil {
		fmt.Println("error in fetch", err)
		return
	}

	fmt.Println("success", res)


}
