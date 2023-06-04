package backend

import (
	"database/sql"
	"log"
	"real-time-forum/db"
	"strconv"

	"github.com/gorilla/websocket"
)

type TData struct {
	Cats    []db.Cat
	Threads []db.Thread
}

type TCData struct {
	Body  db.Thread
	Comms []db.Comm
}

func ShowThreads(uname string, conn *websocket.Conn, allConnections []USocket) {
	response := EData{}
	allCategories := db.GetCats()
	allPosts, err := db.GetThreads(uname)
	if err != nil {
		response.Errrr = "Threads ERR"
		conn.WriteJSON(response)
	}
	allData := &TData{
		Cats:    allCategories,
		Threads: allPosts,
	}
	reply := Reply{"threads", allData}
	conn.WriteJSON(reply)
}

func ShowThread(topic_id string, allConnections []USocket) {
	topicContent := db.GetThread(topic_id)
	topicComments := db.GetThreadComms(topic_id)
	allData := &TCData{
		Body:  topicContent,
		Comms: topicComments,
	}
	reply := Reply{"thread_" + topic_id, allData}
	for _, oneConnection := range allConnections {
		log.Println("KUI SIIA J6UAB SIIS KORRAS")
		oneConnection.Conn.WriteJSON(reply)
	}
}

func ShowCat(uname string, name string, conn *websocket.Conn, allConnections []USocket) {
	response := EData{}
	categoryId, err := db.GetCatId(name)
	if err != nil {
		response.Errrr = "Cat ERR"
		conn.WriteJSON(response)
	}
	allCategories := db.GetCats()
	allPosts, err := db.GetThreadsByCat(categoryId)
	if err != nil {
		response.Errrr = "Cat ERR"
		conn.WriteJSON(response)
	}
	if name == "All" {
		allPosts, _ = db.GetThreads(uname)
	}
	allData := &TData{
		Cats:    allCategories,
		Threads: allPosts,
	}
	reply := Reply{"cat_" + name, allData}
	conn.WriteJSON(reply)
}

func PostThread(catName []string, uid int, content string, privacy string, whosees []string, images []string, time string, gid int, conn *websocket.Conn, allConnections []USocket) {
	response := EData{}
	u, err := db.GetUser(uid)
	if err != nil {
		log.Println(1)
		response.Errrr = "Thread ERR"
		conn.WriteJSON(response)
	}
	topic := db.Thread{
		TId:    0,
		TUId:   uid,
		TUName: u.UName,
		TBody:  content,
		TTime:  time,
		TGid:   gid,
		TWho:   whosees,
		TPriv:  privacy,
	}
	err = db.CreateThread(topic)
	if err != nil {
		log.Println(2)
		response.Errrr = "Thread ERR"
		conn.WriteJSON(response)
	}
	topicId := 0
	if gid > 0 {
		topicId, _ = db.GetLastGroupThread(gid)
	} else {
		topicId, _ = db.GetLastThread()
	}
	for _, i := range catName {
		cid, err := db.GetCatId(i)
		if err == sql.ErrNoRows {
			log.Println("heehee")
			db.CreateCat(i)
			cid, _ := db.GetCatId(i)
			err = db.LinkThreadToCats(topicId, cid)
			if err != nil {
				log.Println(4)
				response.Errrr = "Thread ERR"
				conn.WriteJSON(response)
			}
		} else if err != nil {
			response.Errrr = "Thread ERR"
			conn.WriteJSON(response)
		} else {
			err = db.LinkThreadToCats(topicId, cid)
			if err != nil {
				log.Println(5)
				response.Errrr = "Thread ERR"
				conn.WriteJSON(response)
			}
		}
	}
	log.Println("passed")
	for c, i := range images {
		ProfPic(i, topicId, "thread_", strconv.Itoa(c+1), conn)
		db.CreateThreadPicConnection(topicId, i)
	}
	log.Println("passed")
	for _, i := range whosees {
		db.CreateThreadUserConnection(topicId, i)
	}
	log.Println("passed")
	eCat := db.GetCats()
	eThread, err2 := db.GetThreads(u.UName)
	if err2 != nil {
		log.Println("frickmaster")
		log.Println(err2)
		response.Errrr = "Thread ERR"
		conn.WriteJSON(response)
	}
	log.Println("passed")
	allData := &TData{
		Cats:    eCat,
		Threads: eThread,
	}
	reply := Reply{"threads", allData}
	for _, oneConnection := range allConnections {
		oneConnection.Conn.WriteJSON(reply)
	}
}

func PostComm(topic_id int, creator_id int, content string, creation_date string, post_id string, images []string, conn *websocket.Conn, allConnections []USocket) {
	response := EData{}
	cid, err := db.CreateComm(topic_id, creator_id, content, creation_date)
	if err != nil {
		response.Errrr = "Comm ERR"
		conn.WriteJSON(response)
	}

	for c, i := range images {
		ProfPic(i, cid, "comment_", strconv.Itoa(c+1), conn)
		db.CreateCommPicConnection(cid, i)
	}
	topicContent := db.GetThread(post_id)
	topicComments := db.GetThreadComms(post_id)
	allData := &TCData{
		Body:  topicContent,
		Comms: topicComments,
	}
	reply := Reply{"thread_" + strconv.Itoa(topic_id), allData}
	for _, oneConnection := range allConnections {
		oneConnection.Conn.WriteJSON(reply)
	}
}
