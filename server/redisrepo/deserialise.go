package redisrepo

import (
	"chatapp-websoket-redis/model"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

type Document struct {
	ID      string `json:"_id"`
	Payload []byte `json:"payload"`
	Total   int64  `json:"total"`
}


// Deserialise: This is a special function, you can even use it in other applications also where you’re using Redisearch. 
// Redisearch returns an interface that is of an array of interface `[]interface{}` types. 
// We are converting this response to a custom `Document` struct.


// DeserialiseChat: This function will convert the Document to model.Chat 
// schema and return an array of the model.Chat.


// DeserialiseContactList: SortedSet returns the data in member and score format. 
// This function will convert it to the model.ContactList.

func Deserialise(res interface{}) []Document {
	switch v := res.(type) {
	case []interface{}:
		if len(v) > 1 {
			total := len(v) - 1
			var docs = make([]Document, 0, total/2)

			for i := 1; i <= total; i = i + 2 {
				arrOfValues := v[i+1].([]interface{})
				value := arrOfValues[len(arrOfValues)-1].(string)

				// add _id in the response
				doc := Document{
					ID:      v[i].(string),
					Payload: []byte(value),
					Total:   v[0].(int64),
				}

				docs = append(docs, doc)
			}
			return docs
		}
	default:
		log.Printf("different response type otherthan []interface{}. type: %T", res)
		return nil
	}

	return nil
}

func DeserialiseChat(docs []Document) []model.Chat {
	chats := []model.Chat{}
	for _, doc := range docs {
		var c model.Chat
		json.Unmarshal(doc.Payload, &c)

		c.ID = doc.ID
		chats = append(chats, c)
	}

	return chats
}

func DeserialiseContactList(contacts []redis.Z) []model.ContactList {
	contactList := make([]model.ContactList, 0, len(contacts))

	// improvement tip: use switch to get type of contact.Member
	// handle unknown type accordingly
	for _, contact := range contacts {
		contactList = append(contactList, model.ContactList{
			Username:     contact.Member.(string),
			LastActivity: int64(contact.Score),
		})
	}

	return contactList
}