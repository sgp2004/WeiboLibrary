package dao

import "strconv"
import "../wb_util"

type BorrowInfo struct {
	Id           int
	BookId       int
	BookInfo     *Book
	OwnerId      int
	Owner        *User
	BorrowerId   int
	Borrower     *User
	Note         string
	StartTime    string
	EndTime      string
	Status       int
    returned     bool
}

func CreateInfo(info *BorrowInfo) {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "insert into borrow_info(book_id,owner_id,borrower_id) values(?,?,?)"
	db.Prepare(sql)
    print("create info")
    print(info)
	exec(db, sql, info.BookId, info.OwnerId, info.BorrowerId)
}

func BooksBorrowInfoByUserName(userName string) ([]*BorrowInfo, []*BorrowInfo ) {
    user := QueryUserByUserName(userName)

	db := NewLibraryDB()
	defer closeDB(db)
	sql := "select id,book_id,owner_id,borrower_id,note,start_time,end_time,is_back from borrow_info where (borrower_id=? or Owner_Id=?) and is_back =0 order by start_time desc"
	db.Prepare(sql)
	rows, _ := db.Query(sql, user.Id, user.Id)
	borrowInfos := make([]*BorrowInfo, 0, 10)
	lendInfos := make([]*BorrowInfo, 0, 10)
    bookIds := make([]string, 0, 10)
    userIds := make([]string, 0, 10)
    print(sql)
    print(rows)
	for rows.Next() {
		info := new(BorrowInfo)
		row_err := rows.Scan(&info.Id, &info.BookId, &info.OwnerId, &info.BorrowerId, &info.Note, &info.StartTime, &info.EndTime, &info.Status)
		if row_err != nil {
			print("Row error!!")
			return borrowInfos,lendInfos
		}
        print(info)       
        if(info.BorrowerId == user.Id){	
		    borrowInfos = append(borrowInfos, info)
        }else{
             lendInfos = append(lendInfos, info)
        }

		owner := strconv.Itoa(info.OwnerId)
		if wb_util.FindStrPos(owner, userIds) == -1 {
			userIds = append(userIds, owner)
		}

		borrower := strconv.Itoa(info.BorrowerId)
		if wb_util.FindStrPos(borrower, userIds) == -1 {
			userIds = append(userIds, borrower)
		}

		bookId := strconv.Itoa(info.BookId)
		if wb_util.FindStrPos(bookId, bookIds) == -1 {
			bookIds = append(bookIds, bookId)
		}
	}
    print(borrowInfos)
    print(lendInfos)
	/*get books name, owners name and borrowers name */
    usersMap ,booksMap := make(map[string]*User),make(map[string]*Book)
    size := len(borrowInfos) + len(lendInfos)
    if size>0{
	    usersMap = QueryUserMapsByIds(userIds)
	    booksMap = QueryBookMapsByIds(bookIds)
    }
    if len(borrowInfos) > 0 {
    	for _, v := range borrowInfos {
	    	v.BookInfo = booksMap[strconv.Itoa(v.BookId)]
		    v.Owner = usersMap[strconv.Itoa(v.OwnerId)]
    	}
    }

    if len(lendInfos) > 0 {
    	for _, v := range lendInfos {
	    	v.BookInfo = booksMap[strconv.Itoa(v.BookId)]
		    v.Borrower = usersMap[strconv.Itoa(v.BorrowerId)]
	    }
	}

	return borrowInfos,lendInfos


    
}

func TimeLine(start, limit int) []*BorrowInfo {
	db := NewLibraryDB()
	defer closeDB(db)
	sql := "select id,book_id,owner_id,borrower_id,note,start_time,end_time,is_back from borrow_info order by start_time desc limit  ?,?"
	db.Prepare(sql)
	rows, _ := db.Query(sql, start, limit)
	infos := make([]*BorrowInfo, 0, limit)
	userIds := make([]string, 0, limit)
	bookIds := make([]string, 0, limit)

	for rows.Next() {
		info := new(BorrowInfo)
		row_err := rows.Scan(&info.Id, &info.BookId, &info.OwnerId, &info.BorrowerId, &info.Note, &info.StartTime, &info.EndTime, &info.Status)
		if row_err != nil {
			print("Row error!!")
			return infos
		}
		infos = append(infos, info)
		owner := strconv.Itoa(info.OwnerId)
		if wb_util.FindStrPos(owner, userIds) == -1 {
			userIds = append(userIds, owner)
		}

		borrower := strconv.Itoa(info.BorrowerId)
		if wb_util.FindStrPos(borrower, userIds) == -1 {
			userIds = append(userIds, borrower)
		}
		bookId := strconv.Itoa(info.BookId)
		if wb_util.FindStrPos(bookId, bookIds) == -1 {
			bookIds = append(bookIds, bookId)
		}

	}
	/*get books name, owners name and borrowers name */
	if len(infos) > 0 {
		usersMap := QueryUserMapsByIds(userIds)
		booksMap := QueryBookMapsByIds(bookIds)
        print("begin")
        print(usersMap)
        print(booksMap)
		for _, v := range infos {
			v.BookInfo = booksMap[strconv.Itoa(v.BookId)]
			v.Owner = usersMap[strconv.Itoa(v.OwnerId)]
			v.Borrower = usersMap[strconv.Itoa(v.BorrowerId)]
		}
	}
	return infos
}
func SetBorrowInfoNote(id int,note string){
    db := NewLibraryDB()
    defer closeDB(db)
    sql := "update `borrow_info` set note=?  where id =?" 
    print(sql)
    db.Prepare(sql)
    exec(db, sql, note, id)
}

func SetBorrowInfoStatus(id, status int) {
    db := NewLibraryDB()
    defer closeDB(db)
    sql := "update `borrow_info` set is_back=?  where id =?"
    print(sql)
    db.Prepare(sql)
    exec(db, sql, status, id)
}
